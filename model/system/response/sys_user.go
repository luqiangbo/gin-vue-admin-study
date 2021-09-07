package response

import (
	"go-class/model/system/tables"
)

type SysUserResponse struct {
	User tables.SysUser `json:"user"`
}

type LoginResponse struct {
	User      tables.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expires_at"`
}
