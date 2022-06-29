package controller

import (
	"blog-go/middleware"
	"blog-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(c *gin.Context) {
	userid := c.PostForm("userid")
	userpass := c.PostForm("userpass")

	if !(userid == "admin" && userpass == "1234") {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	user := models.User{
		UserID:       "admin",
		UserName:     "admin1",
		UserPassword: "123456",
	}

	//生成token
	token, err := middleware.GenToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"token": token},
	})
}

func Test(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"user": user},
	})
}
