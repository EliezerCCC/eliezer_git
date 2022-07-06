package models

import (
	"blog-go/dao/mysql"
	"time"
)

type Record struct {
	RecordId       int       `json:"recordid" gorm:"column:id"`
	RecordTitle    string    `json:"recordtitle" gorm:"column:title"`
	RecordContent  string    `json:"recordcontent" gorm:"column:content"`
	RecordUserId   string    `json:"recorduserid" gorm:"column:userid"`
	RecordUserName string    `json:"recordusername" gorm:"column:username"`
	CreateTime     time.Time `json:"createtime" gorm:"column:createtime"`
}

// AddRecord 发表博客
func AddRecord(record Record) (err error) {
	err = mysql.DB.Create(&record).Error
	return
}

// AllRecord 所有博客信息
func AllRecord() (recordlist []Record, err error) {
	err = mysql.DB.Find(&recordlist).Error
	return
}

// OneAllRecord 某用户博客信息
func OneAllRecord(userid string) (recordlist []Record, err error) {
	err = mysql.DB.Where("userid = ?", userid).Find(&recordlist).Error
	return
}

// DeleteRecord 删除博客
func DeleteRecord(record Record) (err error) {
	err = mysql.DB.Where("Id= ?", record.RecordId).Delete(&Record{}).Error
	return
}

// UpdataRecord 修改博客信息
func UpdataRecord(record Record) (err error) {
	err = mysql.DB.Where("Id = ?", record.RecordId).Save(&record).Error
	return
}
