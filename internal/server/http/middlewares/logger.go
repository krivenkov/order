package middlewares

import (
	"net/http"

	"github.com/krivenkov/pkg/mlog"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	next   http.Handler
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Provide(next http.Handler) http.Handler {
	l.next = next
	return l
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := l.logger.With()

	// inject logger to context
	ctx := mlog.CtxWithLogger(r.Context(), logger)

	r = r.WithContext(ctx)
	wp := &LoggerResWriter{prev: w}

	logger = logger.With(
		zap.String("httpMethod", r.Method),
		zap.String("httpURL", r.URL.String()),
	)

	l.next.ServeHTTP(wp, r)

	logger.Debug("http request completed",
		zap.Int("httpStatusCode", wp.StatusCode),
	)
}

type LoggerResWriter struct {
	prev       http.ResponseWriter
	StatusCode int
}

func (l *LoggerResWriter) Header() http.Header {
	return l.prev.Header()
}

func (l *LoggerResWriter) Write(content []byte) (int, error) {
	return l.prev.Write(content)
}

func (l *LoggerResWriter) WriteHeader(statusCode int) {
	l.StatusCode = statusCode
	l.prev.WriteHeader(statusCode)
}
