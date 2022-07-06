package controller

import (
	"blog-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//Add 发表博客
func Add(c *gin.Context) {
	var record models.Record
	c.ShouldBind(&record)

	user := c.MustGet("user").(models.User)
	token := c.MustGet("token")
	_ = models.CUser(&user)
	fmt.Println(user)
	record.RecordUserName = user.UserName
	record.RecordUserId = user.UserID
	record.CreateTime = time.Now()

	fmt.Println("发表博客:", record)

	err := models.AddRecord(record)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("发表博客失败！")
	} else {
		fmt.Println("发表博客成功！")
		c.JSON(200, gin.H{
			"result": "发表博客成功！",
			"token":  token,
		})
	}
}

// AllRecord 获取所有博客
func AllRecord(c *gin.Context) {
	RecordList, err := models.AllRecord()
	fmt.Println("所有博客信息 ", RecordList)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"recordlist": RecordList})
	}
}

// OneAllRecord 获取某人所有博客
func OneAllRecord(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	token := c.MustGet("token")
	RecordList, err := models.OneAllRecord(user.UserID)
	fmt.Println("某用户博客信息 ", RecordList)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"recordlist": RecordList,
			"token":      token,
		})
	}
}

// DeleteRecord 删除某条博客
func DeleteRecord(c *gin.Context) {
	var record models.Record
	c.ShouldBind(&record)
	token := c.MustGet("token")
	fmt.Println("删除博客：", record)

	err := models.DeleteRecord(record)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("删除失败！")
	} else {
		fmt.Println("删除成功！")
		c.JSON(200, gin.H{
			"result": "删除成功！",
			"token":  token,
		})
	}
}

// UpdataRecord 修改文章
func UpdataRecord(c *gin.Context) {
	var record models.Record
	c.ShouldBind(&record)
	token := c.MustGet("token")
	fmt.Println("修改文章:", record)

	err := models.UpdataRecord(record)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("修改失败！")
	} else {
		fmt.Println("修改成功！")
		c.JSON(200, gin.H{
			"result": "修改成功！",
			"token":  token,
		})
	}

}
