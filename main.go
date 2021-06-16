package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"net/http"
	"time"
)

type commonModelFields struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	commonModelFields
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

func main() {
	dsn := "root:Aa123456@tcp(104.128.94.5:3306)/go_class?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	r := gin.Default()
	// 创建
	r.POST("/student", func(c *gin.Context) {
		user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		result := db.Create(&user)
		fmt.Println(result)
		fmt.Println(user)
		c.JSON(200, gin.H{
			"code": user,
		})
	})
	r.GET("/student", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"code": 111,
		})
	})
	r.Run(":1001")
}
