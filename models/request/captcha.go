package request

type Captcha struct {
	CaptchaId  string `json:"captchaId"`
	CaptchaVal string `json:"captchaVal"`
}
