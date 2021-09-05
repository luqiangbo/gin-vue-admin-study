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

type ChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
