package dto

type BookDataDto struct {
	Name      string `json:"name" binding:"required" validateErrMsg:"缺少name字段"`
	Author    string `json:"author"  binding:"required" validateErrMsg:"缺少author字段"`
	ISBN      int64  `json:"ISBN" binding:"required,bookISBN" validateErrMsg:"ISBN为10或13位的数字"`
	Published string `json:"published" binding:"required" validateErrMsg:"缺少published字段"`
}

type EdtiBookDataDto struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	ISBN      int64  `json:"ISBN" binding:"bookISBN" validateErrMsg:"ISBN为10或13位的数字"`
	Published string `json:"published"`
}
