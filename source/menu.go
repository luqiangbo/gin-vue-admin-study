package source

import (
	"github.com/gookit/color"
	"go-class/global"
	"go-class/model/system/tables"
	"gorm.io/gorm"
	"time"
)

type menu struct {
}

var BaseMenu = new(menu)

var menus = []tables.SysBaseMenu{
	{
		GVA_MODEL: global.GVA_MODEL{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		MenuLevel: 0,
		ParentId:  "0",
		Path:      "dashboard",
		Hidden:    false,
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: tables.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	},
	{GVA_MODEL: global.GVA_MODEL{
		ID:        2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
		MenuLevel: 0,
		Hidden:    false,
		ParentId:  "0",
		Path:      "about",
		Name:      "about",
		Component: "view/about/index.vue",
		Sort:      7,
		Meta: tables.Meta{
			Title: "关于我们",
			Icon:  "info",
		},
	},
}

func (m *menu) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]tables.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menu 表的初始数据已存在!")
			return nil
		}

		if err := tx.Create(&menus).Error; err != nil {
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menu 表初始数据成功!")
		return nil
	})
	return nil
}
