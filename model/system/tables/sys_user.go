package tables

import "go-class/global"
import "github.com/satori/go.uuid"

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"user_name" gorm:"comment:用户登录名"`
	Password    string       `json:"-" gorm:"comment:用户登录密码"`
	NickName    string       `json:"nick_name" gorm:"default:系统用户;comment:用户昵称"`
	SideMode    string       `json:"side_mode" gorm:"default:dark;comment:用户侧边主题"`
	HeaderImg   string       `json:"header_img" gorm:"default:123;comment:用户头像"`
	BaseColor   string       `json:"base_color" gorm:"default:#fff";comment:活跃颜色"`
	AuthorityId string       `json:"authority_id" gorm:"default:888;comment:用户角色ID"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}
