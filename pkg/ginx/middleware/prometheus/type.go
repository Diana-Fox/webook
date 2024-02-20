package prometheus

import "github.com/gin-gonic/gin"

type PrometheusMiddleware interface {
	//普罗米修斯的一些构造
	Build() gin.HandlerFunc //构造
}
