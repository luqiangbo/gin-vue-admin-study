package request

// 登录

type Login struct {
	Username  string `json:"username"`   // 用户名
	Password  string `json:"password"`   // 密码
	Captcha   string `json:"captcha"`    // 验证码
	CaptchaId string `json:"captcha_id"` // 验证码ID
}
