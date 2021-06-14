package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TableClass struct {
	gorm.Model
	ClassName string
	Students  []TableStudent
}

type TableStudent struct {
	gorm.Model
	StudentName string
	ClassId     uint
	Card        TableCard
	Teachers    []TableTeacher `gorm:"many2many:table_student_teacher;"`
}

type TableCard struct {
	gorm.Model
	Num       int
	StudentId uint
}

type TableTeacher struct {
	gorm.Model
	TeacherName string
	Students    []TableStudent `gorm:"many2many:table_student_teacher;"`
}

func main() {
	dsn := "root:Aa123456@tcp(104.128.94.5:3306)/go_class?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&TableStudent{}, &TableTeacher{}, &TableCard{}, &TableClass{})

	defer db.Close()
	r := gin.Default()
	// 创建
	r.POST("/student", func(c *gin.Context) {
		var student TableStudent
		_ = c.BindJSON(&student)
		db.Create(&student)
	})
	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		var student TableStudent
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("Card").Where("id = ?", id).First(&student)
		c.JSON(200, gin.H{
			"s": student,
		})
	})
	r.Run(":1001")
}
