package source

import (
	"github.com/gookit/color"
	"go-class/global"
	"gorm.io/gorm"
)

type authoritiesMenus struct{}

var AuthoritiesMenus = new(authoritiesMenus)

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"888", 1},
	{"888", 2},
	{"8881", 1},
	{"8881", 2},
	{"9528", 1},
	{"9528", 2},
}

//sys_authority_m2m_sys_base_menu 表数据初始化

func (a *authoritiesMenus) Init() error {
	return global.GVA_DB.Table("sys_authority_m2m_sys_base_menu").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 6 {
			color.Danger.Println("\n[Mysql] --> sys_authority_m2m_sys_base_menu 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_m2m_sys_base_menu 表初始数据成功!")
		return nil
	})
}
