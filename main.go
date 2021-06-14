package main

import (
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
	IdCard      TableIdCard
	Teachers    []TableTeacher `gorm:"many2many:table_student_teacher;"`
}

type TableIdCard struct {
	gorm.Model
	Num int
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
	db.AutoMigrate(&TableStudent{}, &TableTeacher{}, &TableIdCard{}, &TableClass{})

	i := TableIdCard{
		Num: 123456,
	}
	s := TableStudent{
		StudentName: "卢强波",
		IdCard:      i,
	}
	t := TableTeacher{
		TeacherName: "老师李",
		Students:    []TableStudent{s},
	}
	c := TableClass{
		ClassName: "三年二班",
		Students:  []TableStudent{s},
	}
	_ = db.Create(&c).Error
	_ = db.Create(&t).Error
	defer db.Close()
}
