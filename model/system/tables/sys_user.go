package tables

import "go-class/global"
import "github.com/satori/go.uuid"

type SysUser struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username  string    `json:"user_name" gorm:"comment:用户登录名"`
	Password  string    `json:"-" gorm:"comment:用户登录密码"`
	NickName  string    `json:"nick_name" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string    `json:"header_img" gorm:"default:123;comment:用户头像"`
}
