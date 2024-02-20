package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

type SummaryEscalation struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
}

func NewSummaryEscalation(Namespace string, Subsystem string,
	Name string, Help string) PrometheusMiddleware {
	return &SummaryEscalation{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      Name,
		Help:      Help,
	}
}
func (s SummaryEscalation) Build() gin.HandlerFunc {
	labels := []string{"method", "pattern", "status"}
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: s.Namespace,
		Subsystem: s.Subsystem,
		Name:      s.Name,
		Help:      s.Help,
	}, labels)
	return func(context *gin.Context) {
		start := time.Now()
		defer func() {
			duration := time.Since(start)
			path := context.FullPath()
			if path == "" {
				path = "unknown"
			}
			summary.WithLabelValues(context.Request.Method,
				path,
				strconv.Itoa(context.Writer.Status())).
				Observe(float64(duration.Milliseconds()))
		}()
		context.Next() //去执行实际方法

	}
}
