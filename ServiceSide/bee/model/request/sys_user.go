package request

import "bee/model"

// 接收注册数据结构体
type RegisterStruct struct {
	Username string 	`json:"userName"`
	Password string 	`json:"password"`
	NickName string		`json:"nickName" gorm:"default:'cute bee'"`
	HeaderImg string    `json:"headerImg" gorm:"default:'https://jdc.jd.com/img/200'"`
	AuthorityId string	`json:"authorityId"`
}

// 返回注册结果
type SysUserResponse struct {
	User model.SysUser  `json:"user"`
}

type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

