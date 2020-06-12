package logging

import (
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
}

type Params struct {
	fx.In

	Logger *zap.Logger
}

func New(p Params) (*Middleware, error) {
	return &Middleware{
		logger: p.Logger,
	}, nil
}

func (m *Middleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := newResponseRecorder(w)
		next.ServeHTTP(recorder, r)
		m.logger.Info(
			"Handled request.",
			zap.Int("status", recorder.status),
			zap.String("method", r.Method),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

type responseRecorder struct {
	http.ResponseWriter

	status      int
	wroteHeader bool
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		status:         0,
		wroteHeader:    false,
	}
}

func (r *responseRecorder) Status() int {
	return r.status
}

func (r *responseRecorder) WriteHeader(code int) {
	if r.wroteHeader {
		return
	}

	r.status = code
	r.wroteHeader = true
}
