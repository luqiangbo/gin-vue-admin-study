package source

import (
	"github.com/gookit/color"
	"go-class/global"
	modelSystemDB "go-class/model/system/tables"
	"gorm.io/gorm"
)

type userAuthority struct {
}

var UserAuthority = new(userAuthority)

var userAuthorityModel = []modelSystemDB.SysUseAuthority{
	{1, "888"},
	{1, "8881"},
	{1, "9528"},
	{2, "888"},
}

// 数据库初始化

func (a *userAuthority) Init() error {
	return global.GVA_DB.Model(&modelSystemDB.SysUseAuthority{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_user_id IN (1, 2)").Find(&[]modelSystemDB.SysUseAuthority{}).RowsAffected == 4 {
			color.Danger.Println("\n[Mysql] --> sys_user_authority 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userAuthorityModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_user_authority 表初始数据成功!")
		return nil
	})
}
