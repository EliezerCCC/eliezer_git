package routers

import (
	"blog-go/controller"
	"blog-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//创建默认路由引擎
	r := gin.Default()

	//允许跨域访问
	r.Use(middleware.Cors())

	//用户登录
	r.POST("/login", controller.Login)
	//用户注册
	r.POST("/register", controller.Register)
	//test
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.Test)

	//发表博客
	r.POST("/add", middleware.JWTAuthMiddleware(), controller.Add)
	//获取所有博客
	r.GET("/all", controller.AllRecord)
	//获取某条博客信息
	r.POST("/getone", middleware.JWTAuthMiddleware(), controller.GetOne)
	//获取某个用户所有博客
	r.GET("/oneall", middleware.JWTAuthMiddleware(), controller.OneAllRecord)
	//编辑博客
	r.PUT("/updata", middleware.JWTAuthMiddleware(), controller.UpdataRecord)
	//删除博客
	r.DELETE("/delete", middleware.JWTAuthMiddleware(), controller.DeleteRecord)

	return r
}
