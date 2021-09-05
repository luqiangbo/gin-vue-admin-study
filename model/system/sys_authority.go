package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authority_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authority_name" gorm:"comment:角色名"`
	ParentId        string         `json:"parent_id" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"data_authority_id" gorm:"many2many:m2m_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"sys_base_menus" gorm:"many2many:m2m_authority_menus;"`
	DefaultRouter   string         `json:"default_router" gorm:"comment:默认菜单;default:dashboard"`
}
