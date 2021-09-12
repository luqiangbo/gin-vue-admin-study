package source

import (
	"github.com/gookit/color"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/system/tables"
	"gorm.io/gorm"
	"time"
)

type authority struct {
}

var Authority = new(authority)

var authorities = []tables.SysAuthority{
	{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		AuthorityId:   "888",
		AuthorityName: "普通用户",
		ParentId:      "0",
		DefaultRouter: "dashboard",
	},
	{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		AuthorityId:   "8881",
		AuthorityName: "普通用户子角色",
		ParentId:      "888",
		DefaultRouter: "dashboard",
	},
	{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		AuthorityId:   "9528",
		AuthorityName: "测试用户",
		ParentId:      "0",
		DefaultRouter: "dashboard",
	},
}

// 数据初始化

func (a *authority) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ?", []string{"888", "9528"}).Find(&[]tables.SysUseAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil {
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
