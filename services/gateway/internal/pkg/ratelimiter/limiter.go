package ratelimiter

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/arvanStorage/pkg/grpcext"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	quotaService metadataApi.QuotaServiceClient
	visitors     map[string]*Visitor
	mu           *sync.Mutex
}

func NewLimiter(metadataAddress string) *Limiter {
	metadataConn := grpcext.NewConnection(metadataAddress)
	return &Limiter{
		quotaService: metadataApi.NewQuotaServiceClient(metadataConn),
		visitors:     make(map[string]*Visitor, 1000),
		mu:           new(sync.Mutex),
	}
}

func (l *Limiter) QuotaRateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := l.getUserID(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			return
		}

		visitor, err := l.getVisitor(ctx, userID)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		// Allow or block the request based on the rate limit
		if visitor.Limiter.Allow() {
			ctx.Next()
		} else {
			ctx.String(http.StatusTooManyRequests, "request limit exceed")
		}
	}
}

// getUserID extracts the user id from the request JSON payload.
func (l *Limiter) getUserID(ctx *gin.Context) (string, error) {
	model := &struct {
		UserID string `json:"user_id"`
	}{}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return "", err
	}
	_ = ctx.Request.Body.Close()

	err = json.Unmarshal(body, model)
	if err != nil {
		return "", err
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	return model.UserID, nil
}

// getVisitor retrieves the visitor's rate limit information from the cache or creates a new one if it doesn't exist.
func (l *Limiter) getVisitor(ctx *gin.Context, userID string) (*Visitor, error) {
	limit, err := l.quotaService.RequestLimit(ctx, &metadataApi.QuotaRequestLimitReq{UserID: userID})
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	visitor, exists := l.visitors[userID]
	if !exists {
		second := limit.Duration / int64(time.Second)
		requestLimit := limit.Count / second

		limiter := rate.NewLimiter(rate.Limit(requestLimit), 1)
		visitor = &Visitor{
			Limiter: limiter,
			UserID:  userID,
		}
		l.visitors[userID] = visitor
	}

	return visitor, nil
}
