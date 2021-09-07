package request

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录

type Login struct {
	Username  string `json:"username"`   // 用户名
	Password  string `json:"password"`   // 密码
	Captcha   int    `json:"captcha"`    // 验证码
	CaptchaId int    `json:"captcha_id"` // 验证码ID
}

//改密码

type ChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// 设置用户权限

type SetUserAuth struct {
	AuthorityId string `json:"authority_id"`
}

// 设置用户权限

type SetUserAuthorities struct {
	ID           uint     `json:"id"`
	AuthorityIds []string `json:"authority_ids"` // 角色id
}
