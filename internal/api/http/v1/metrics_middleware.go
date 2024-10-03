package v1

import (
	"VacancyService/internal/metrics"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type metricsResponseWriter struct {
	gin.ResponseWriter
	bytesWritten int64
}

func (w *metricsResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.bytesWritten += int64(n)
	return n, err
}

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path

		rec := &metricsResponseWriter{ResponseWriter: c.Writer}
		c.Writer = rec

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start).Seconds()

		metrics.VacancyHTTPRequest.WithLabelValues(method, path).Inc()
		metrics.VacancyRequestDuration.WithLabelValues(method, path, fmt.Sprintf("%d", status)).Observe(duration)

		if status >= 500 && status < 600 {
			metrics.Vacancy5xxRequest.WithLabelValues(method, path).Inc()
		}

		inboundBytes := c.Request.ContentLength
		outboundBytes := rec.bytesWritten

		metrics.VacancyTrafficInbound.Observe(float64(inboundBytes))
		metrics.VacancyTrafficOutbound.Observe(float64(outboundBytes))
	}
}
