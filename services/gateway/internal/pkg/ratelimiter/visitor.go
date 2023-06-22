package ratelimiter

import "golang.org/x/time/rate"

type Visitor struct {
	Limiter *rate.Limiter
	UserID  string
}
