package response

import "go-class/model/system/tables"

type SysMenusResponse struct {
	Menus []tables.SysMenu `json:menus`
}
