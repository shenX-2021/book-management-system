package dto

type UserRegisterDto struct {
	Name         string `json:"name" binding:"required,min=1,max=12" validateErrMsg:"name必须为1~12位的字符串"`
	Account      string `json:"account" binding:"required,min=4,max=16" validateErrMsg:"account必须为4~16位的字符串"`
	Pwd          string `json:"pwd" binding:"required,min=6,max=32" validateErrMsg:"pwd必须为6~32位的字符串"`
	ConfirmedPwd string `json:"confirmedPwd" binding:"required,min=6,max=32,eqfield=Pwd" validateErrMsg:"两次密码不一致"`
}

type UserLoginDto struct {
	Account string `json:"account" binding:"required,min=4,max=16" validateErrMsg:"account必须为4~16位的字符串"`
	Pwd     string `json:"pwd" binding:"required,min=6,max=32" validateErrMsg:"pwd必须为6~32位的字符串"`
}
