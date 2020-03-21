package models

type RegisterInfo struct {
	Mobile string `json:"mobile"`
	Pwd string `json:"pwd"`
}


type Response struct {
	Code int 	`json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

