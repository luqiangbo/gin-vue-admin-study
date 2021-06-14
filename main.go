package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100)"`
}

func (u User) TableName() string {
	return "table_users"
}

func main() {
	dsn := "root:Aa123456@tcp(104.128.94.5:3306)/go_class?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	//fmt.Println(hello)

	defer db.Close()
}
