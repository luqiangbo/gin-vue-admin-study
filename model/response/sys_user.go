package response

import "goClass/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expires_at"`
}
