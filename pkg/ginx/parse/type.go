package parse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Wrap 只做参数解析，单纯才好拼接
func Wrap[T any](fn func(ctx *gin.Context, req T)) gin.HandlerFunc {
	return func(context *gin.Context) {
		var t T
		if err := context.Bind(&t); err != nil {
			//参数异常
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		fn(context, t)
	}
}
