package metadata

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/arvanStorage/services/gateway/internal/app/metadata/object"
	"github.com/h-varmazyar/arvanStorage/services/gateway/internal/app/metadata/quota"
	log "github.com/sirupsen/logrus"
)

func RegisterRoutes(log *log.Logger, e *gin.Engine, metadataAddress string) {
	metadataRouter := e.Group("/metadata")
	object.NewController(log, e, metadataAddress).RegisterRoutes(metadataRouter)
	quota.NewController(log, e, metadataAddress).RegisterRoutes(metadataRouter)
}
