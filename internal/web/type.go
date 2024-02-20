package web

import (
	"github.com/gin-gonic/gin"
	"webook/internal/domian/req"
)

type UserHandler interface {
	RegisterRoutes(engine *gin.Engine)
	SingUp(ctx *gin.Context, req req.UserReq)
	//通过验证码登录的
	LoginByCode(ctx *gin.Context)
	//通过密码登录的
	LoginByPassword(ctx *gin.Context)
}
