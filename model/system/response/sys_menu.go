package response

import "gin-vue-admin-study/model/system/tables"

type SysMenusResponse struct {
	Menus []tables.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []tables.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu tables.SysBaseMenu `json:"menu"`
}
