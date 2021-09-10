package tables

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      `json:"created_at"` // 创建时间
	UpdatedAt       time.Time      `json:"updated_at"` // 更新时间
	DeletedAt       *time.Time     `json:"deleted_at" sql:"index"`
	AuthorityId     string         `json:"authority_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authority_name" gorm:"comment:角色名"`
	ParentId        string         `json:"parent_id" gorm:"comment:父角色ID"`
	AuthorityIdList []SysAuthority `json:"authority_id_list" gorm:"many2many:sys_authority_m2m_sys_authority"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	BaseMenuList    []SysBaseMenu  `json:"base_menu_list" gorm:"many2many:sys_authority_m2m_sys_base_menu;"`
	DefaultRouter   string         `json:"default_router" gorm:"comment:默认菜单;default:dashboard"`
}
