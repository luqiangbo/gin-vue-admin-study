package system

import (
	"errors"
	"gin-vue-admin-study/global"
	commonReq "gin-vue-admin-study/model/common/request"
	"gin-vue-admin-study/model/system/tables"
	"gorm.io/gorm"
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

// 增加 菜单

func (m *MenuService) AddBaseMenu(param tables.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", param.Name).First(&tables.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复的name, 请修改name")
	}
	return global.GVA_DB.Create(&param).Error
}

// 删 菜单

func (d *MenuService) DeleteBaseMenu(id float64) (err error) {
	err = global.GVA_DB.Preload("Parameters").Where("parent_id = ?", id).First(&tables.SysBaseMenu{}).Error
	if err != nil {
		var menu tables.SysBaseMenu
		db := global.GVA_DB.Preload("AuthorityList").Where("id = ?", id).First(&menu).Delete(&menu)
		if len(menu.AuthorityList) > 0 {
			err = global.GVA_DB.Model(&menu).Association("AuthorityList").Delete(&menu.AuthorityList)
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此裁断存在子菜单不可删除")
	}
	return err
}

// 改 菜单

func (d *MenuService) UpdateBaseMenu(param tables.SysBaseMenu) (err error) {
	var oldMenu tables.SysBaseMenu
	updateMap := make(map[string]interface{})
	updateMap["keep_alive"] = param.KeepAlive
	updateMap["close_tab"] = param.CloseTab
	updateMap["default_menu"] = param.DefaultMenu
	updateMap["parent_id"] = param.ParentId
	updateMap["path"] = param.Path
	updateMap["name"] = param.Name
	updateMap["hidden"] = param.Hidden
	updateMap["component"] = param.Component
	updateMap["title"] = param.Title
	updateMap["icon"] = param.Icon
	updateMap["sort"] = param.Sort

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", param.ID).Find(&oldMenu)
		if oldMenu.Name != param.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", param.ID, param.Name).First(&tables.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.GVA_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&tables.SysBaseMenuParameter{}, "sys_base_menu_id = ?", param.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		if len(param.Parameters) > 0 {
			for k := range param.Parameters {
				param.Parameters[k].SysBaseMenuID = param.ID
			}
			txErr = tx.Create(&param.Parameters).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}
		txErr = db.Updates(updateMap).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

// 通过id获取菜单详情

func (d *MenuService) GetBaseMenuById(id float64) (err error, res tables.SysBaseMenu) {
	err = global.GVA_DB.Preload("Parameters").Where("id = ?", id).First(&res).Error
	return
}
