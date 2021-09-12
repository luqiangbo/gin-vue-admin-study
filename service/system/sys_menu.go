package system

import (
	"go-class/global"
	commonReq "go-class/model/common/request"
	"go-class/model/system/tables"
	"strconv"
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

// 获取全部菜单
func (m *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]tables.SysBaseMenu) {
	var allMenus []tables.SysBaseMenu
	treeMap = make(map[string][]tables.SysBaseMenu)
	err = global.GVA_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
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

//
func (m *MenuService) getBaseChildrenList(menu *tables.SysBaseMenu, treeMap map[string][]tables.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// 获取全部菜单

func (m *MenuService) GetBaseMenuTree() (err error, menus []tables.SysBaseMenu) {
	err, treeMap := m.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = m.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

// 查看角色的菜单树

func (m *MenuService) GetMenuAuthority(param *commonReq.GetAuthorityId) (err error, menus []tables.SysMenu) {
	err = global.GVA_DB.Where("authority_id = ?", param.AuthorityId).Order("sort").Find(&menus).Error
	return err, menus
}

// 增加menu和角色的关系

func (m *MenuService) AddMenuAuthority(menus []tables.SysBaseMenu, authorityId string) (err error) {
	var authNew tables.SysAuthority
	authNew.AuthorityId = authorityId
	authNew.BaseMenuList = menus
	err = AuthorityServiceApp.SetMenuAuthority(&authNew)
	return err
}
