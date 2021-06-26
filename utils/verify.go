package utils

var (
	RegisterVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	LoginVerify    = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
