package request

type ResisterStruct struct {
	Username string 	`json:"userName"`
	Password string 	`json:"password"`
	NickName string		`json:"nickName"`
	HeaderImg string	`json:"headerImg"`
	AuthorityId string	`json:"authorityId"`
}

type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

