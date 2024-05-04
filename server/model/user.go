package model

type User struct {
	BaseModel
	Account string `json:"account" gorm:"unique"`
	Name    string `json:"name" gorm:"index"`
	Pwd     string `json:"pwd"`
}

func (User) TableName() string {
	return "tb_user"
}
