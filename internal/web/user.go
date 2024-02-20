package web

import (
	"github.com/Diana-Fox/webook/internal/domian"
	"github.com/Diana-Fox/webook/internal/domian/req"
	"github.com/Diana-Fox/webook/internal/service"
	"github.com/Diana-Fox/webook/pkg/ginx/parse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	us service.UserService
}

func (u *userHandler) RegisterRoutes(engine *gin.Engine) {
	group := engine.Group("/users")
	group.POST("/signup", parse.Wrap(u.SingUp)) //注册
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandler{
		us: us,
	}
}
func (u *userHandler) SingUp(ctx *gin.Context, req req.UserReq) {
	//本来打算统一处理异常，但是go都创造者似乎就是希望异常能被返回
	//以及考虑到后续grpc的话，异常还是返回比较好
	du := domian.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	err := u.us.SignUp(ctx, du)
	if err != nil {
		ctx.JSON(http.StatusOK,
			domian.Error[any](500, "注册异常:"+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, domian.Success[any]())
}

// 通过验证码登录
func (u *userHandler) LoginByCode(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// LoginByPassword 通过密码登录
func (u *userHandler) LoginByPassword(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
