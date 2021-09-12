package system

import (
	"go-class/global"
	"go-class/model/system/tables"
)

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

func (m *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]tables.SysMenu) {

	var allMenus []tables.SysMenu
	treeMap = make(map[string][]tables.SysMenu)
	err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// 获取动态菜单树

func (m *MenuService) GetMenuTree(authorityId string) (err error, menus []tables.SysMenu) {
	err, menuTree := m.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = m.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// 获取子菜单

func (m *MenuService) getChildrenList(menu *tables.SysMenu, treeMap map[string][]tables.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
