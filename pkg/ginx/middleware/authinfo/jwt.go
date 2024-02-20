package authinfo

import (
	"github.com/gin-gonic/gin"
	"webook/internal/domian"
	"webook/utils"
)

type JwtAuthorityInfoMiddleware struct {
	utils utils.JWTUtils[domian.AuthorityUserInfo]
}

// 方便扩展
func NewJwtAuthorityInfoMiddleware(utils utils.JWTUtils[domian.AuthorityUserInfo]) AuthorityInfoMiddleware {
	return &JwtAuthorityInfoMiddleware{
		utils: utils,
	}
}

// Build 做一些续约和把token放起来的操作
func (j *JwtAuthorityInfoMiddleware) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		//这里是不是可以考虑，尝试获取jwt，也不判断是否需要jwt，拿不到的话，就是没有嘛，没有就不管，
		//后面权限校验的时候会卡掉的
		token := context.GetHeader("x-jwt-token")
		if token != "" {
			////可以去解析
			info, err := j.utils.AnalysisJWT(context, token)
			if err != nil {
				return
			}
			context.Set("info", info)
		}
	}
}
