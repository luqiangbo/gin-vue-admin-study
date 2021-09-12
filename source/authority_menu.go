package source

import (
	"github.com/gookit/color"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/system/tables"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) Init() error {
	if global.GVA_DB.Migrator().HasTable("authority_menu") && global.GVA_DB.Find(&[]tables.SysMenu{}).RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.GVA_DB.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menu`.`id` AS `id`,`sys_base_menu`.`created_at` AS `created_at`, `sys_base_menu`.`updated_at` AS `updated_at`, `sys_base_menu`.`deleted_at` AS `deleted_at`, `sys_base_menu`.`menu_level` AS `menu_level`,`sys_base_menu`.`parent_id` AS `parent_id`,`sys_base_menu`.`path` AS `path`,`sys_base_menu`.`name` AS `name`,`sys_base_menu`.`hidden` AS `hidden`,`sys_base_menu`.`component` AS `component`, `sys_base_menu`.`title`  AS `title`,`sys_base_menu`.`icon` AS `icon`,`sys_base_menu`.`sort` AS `sort`,`sys_authority_m2m_sys_base_menu`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_m2m_sys_base_menu`.`sys_base_menu_id` AS `menu_id`,`sys_base_menu`.`keep_alive` AS `keep_alive`,`sys_base_menu`.`close_tab` AS `close_tab`,`sys_base_menu`.`default_menu` AS `default_menu` from (`sys_authority_m2m_sys_base_menu` join `sys_base_menu` on ((`sys_authority_m2m_sys_base_menu`.`sys_base_menu_id` = `sys_base_menu`.`id`)))").Error; err != nil {
		return err
	}
	color.Info.Println("\n[Mysql] --> authority_menu 视图创建成功!")
	return nil
}
