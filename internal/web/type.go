package web

import (
	"github.com/Diana-Fox/webook/internal/domian/req"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterRoutes(engine *gin.Engine)
	SingUp(ctx *gin.Context, req req.UserReq)
	//通过验证码登录的
	LoginByCode(ctx *gin.Context)
	//通过密码登录的
	LoginByPassword(ctx *gin.Context)
}
