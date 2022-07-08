package mysql_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

// 申明db
var db *gorm.DB

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GoUserInfo struct {
	//BaseModel
	ID       uint   `gorm:"primary_key"`
	Sid      string `gorm:"not null;size:32"`
	SellerId string `gorm:"size:32"`
	Years    string `gorm:"size:32"`
	Gender   string `gorm:"size:32"`
}

func QueryGoUserInfoBySid(c *gin.Context) {

	sid := 1
	var TargetGoUserInfo []GoUserInfo
	err := db.Model(&GoUserInfo{}).Where("sid = ?", sid).Find(&TargetGoUserInfo).Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败",
			"data":    nil,
		})
	}
	fmt.Println("这是一个查询结果： ?", TargetGoUserInfo)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"data":    &TargetGoUserInfo,
	})
}
