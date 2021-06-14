package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TableHello struct {
	gorm.Model
	Name   string
	Sex    bool
	Age    int
	Friend string
}

func main() {
	dsn := "root:Aa123456@tcp(104.128.94.5:3306)/go_class?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&TableHello{})

	// 创建
	//db.Create(&TableHello{
	//	Name:   "lqb3",
	//	Sex:    false,
	//	Age:    89,
	//	Friend: "xiaoming",
	//})

	// 查找一个
	//var hello TableHello
	//db.First(&hello)

	// 查找多个
	//var hello []TableHello
	//db.Find(&hello)

	// 查询条件
	//var hello []TableHello
	//db.Where("age < ?",21).Find(&hello)

	// 更新1
	//db.Where("id = ?", 1).First(&TableHello{}).Update("name", "小二")

	// 更新 2
	//db.Where("id = ?", 1).First(&TableHello{}).Updates(TableHello{
	//	Name: "更新2",
	//	Age:  2,
	//})

	//更新 3
	//db.Where("id = ?", 1).First(&TableHello{}).Updates(map[string]interface{}{
	//	"Name": "更新31",
	//	"Sex":  false,
	//	"Age":  31,
	//})

	//更新 4
	//db.Where("id in (?)", []int{1, 2}).First(&[]TableHello{}).Updates(map[string]interface{}{
	//	"Name": "批量更新",
	//	"Sex":  false,
	//	"Age":  322,
	//})

	// 删除 软
	//db.Where("id in (?) ", []int{1, 2}).Delete(&TableHello{})

	// 删除 硬
	db.Where("id in (?) ", []int{1, 2}).Unscoped().Delete(&TableHello{})

	//fmt.Println(hello)

	defer db.Close()
}
