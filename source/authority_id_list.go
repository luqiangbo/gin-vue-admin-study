package source

import (
	"github.com/gookit/color"
	"go-class/global"
	"gorm.io/gorm"
)

type authorityIdList struct{}

var AuthorityIdList = new(authorityIdList)

type AuthorityIdTo struct {
	AuthorityId                string `gorm:"column:sys_authority_authority_id"`
	AuthorityIdListAuthorityId string `gorm:"column:authority_id_list_authority_id"`
}

var infos = []AuthorityIdTo{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

func (d *authorityIdList) Init() error {
	return global.GVA_DB.Table("sys_authority_m2m_sys_authority").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '9528') ").Find(&[]AuthorityIdTo{}).RowsAffected == 5 {
			color.Danger.Println("\n[Mysql] --> sys_authority_m2m_sys_authority 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_m2m_sys_authority 表初始数据成功!")
		return nil
	})
}
