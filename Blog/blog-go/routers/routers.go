package routers

import (
	"blog-go/controller"
	"blog-go/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//创建默认路由引擎
	r := gin.Default()

	//允许跨域访问
	r.Use(cors.Default())

	//用户登录
	r.GET("/login", controller.Login)

	//test
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.Test)

	return r
}
