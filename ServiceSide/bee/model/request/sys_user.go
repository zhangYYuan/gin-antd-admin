package request

import "bee/model"

// 注册
type RegisterResponse struct {
	Username string 	`json:"userName"`
	Password string 	`json:"password"`
	NickName string		`json:"nickName" gorm:"default:'cute bee'"`
	HeaderImg string    `json:"headerImg" gorm:"default:'https://jdc.jd.com/img/200'"`
	AuthorityId string	`json:"authorityId"`
}

// 返回注册结果
type LoginResponse struct {
	User  		model.SysUser  `json:"user"`
	Token 		string		`json:"token"`
	ExpiresAt 	int64		`json:"expiresAt"`
}

type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

