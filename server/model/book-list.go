package model

type BookData struct {
	BaseModel
	Name      string `json:"name" gorm:"index"`
	Author    string `json:"author"`
	ISBN      int64  `json:"ISBN"`
	Published string `json:"published"`
}

func (BookData) TableName() string {
	return "tb_book_data"
}

func NewBookData(name, author string, ISBN int64, published string) BookData {
	bookData := BookData{}
	bookData.Name = name
	bookData.Author = author
	bookData.ISBN = ISBN
	bookData.Published = published
	return bookData
}
