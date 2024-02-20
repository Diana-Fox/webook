package authinfo

import "github.com/gin-gonic/gin"

type AuthorityInfoMiddleware interface {
	//在这一步解析用户信息,以及做一些续约的事情
	Build() gin.HandlerFunc //构造
}
