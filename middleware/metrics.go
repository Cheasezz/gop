package middleware

import (
	"github.com/Cheasezz/gop/client"
	"github.com/Cheasezz/gop/metrics"
	"strconv"
)

// Metrics sets stats for request and responses
type Metrics struct {
	Metrics *metrics.Metrics
}

func (a *Metrics) ProcessRequest(r *client.Request) {
	a.Metrics.RequestCounter.With("method", r.Method).Add(1)
}

func (a *Metrics) ProcessResponse(r *client.Response) {
	a.Metrics.ResponseCounter.With("status", strconv.Itoa(r.StatusCode)).Add(1)
}
