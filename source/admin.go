package source

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"go-class/global"
	"go-class/model/system/tables"
	"gorm.io/gorm"
	"time"
)

type admin struct {
}

var Admin = new(admin)

// 初始数据
var admins = []tables.SysUser{
	{
		GVA_MODEL: global.GVA_MODEL{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UUID:        uuid.NewV4(),
		Username:    "admin",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		NickName:    "超级管理员",
		HeaderImg:   "http://qmplusimg.henrongyi.top/gva_header.jpg",
		AuthorityId: "888",
	},
	{
		GVA_MODEL: global.GVA_MODEL{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UUID:        uuid.NewV4(),
		Username:    "a303176530",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		NickName:    "QMPlusUser",
		HeaderImg:   "http://qmplusimg.henrongyi.top/1572075907logo.png",
		AuthorityId: "9528",
	},
}

// sys_user 表数据初始化

func (a *admin) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]tables.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_user 表的初始化数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil {
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_user 表初始数据成功")
		return nil
	})
}
