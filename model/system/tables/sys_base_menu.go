package tables

import (
	"go-class/global"
)

type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                              `json:"_"`
	ParentId      string                            `json:"parent_id" gorm:"comment:父菜单ID"`
	Path          string                            `json:"path" gorm:"comment:路由path"`
	Name          string                            `json:"name" gorm:"comment:路由name"`
	Hidden        bool                              `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component     string                            `json:"component" gorm:"comment:对应前端文件路径"`
	Sort          int                               `json:"sort" gorm:"comment:排序标记"`
	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:m2m_authority_menus"`
	Children      []SysBaseMenu                     `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter            `json:"parameters"`
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
}

type Meta struct {
	KeepAlive   bool   `json:"keep_alive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"default_menu" gorm:"comment:是否是基础路由(开发中)"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
	CloseTab    string `json:"close_tab" gorm:"comment:自动关闭tab"`
}

type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}
