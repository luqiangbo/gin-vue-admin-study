package request

import (
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/system/tables"
)

type AddMenuAuthorityInfo struct {
	Menus       []tables.SysBaseMenu `json:"menus"`
	AuthorityId string               `json:"authority_id"`
}

func DefaultMenu() []tables.SysBaseMenu {
	return []tables.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: tables.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
