package service

import (
	"book-management-system-server/dto"
	"book-management-system-server/global"
	"book-management-system-server/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookService struct{}

func (b BookService) Create(dto *dto.BookDataDto) error {
	if !errors.Is(global.DB.Where("ISBN = ?", dto.ISBN).First(&model.BookData{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同ISBN的书籍")
	}

	bookData := model.BookData{Name: dto.Name, Author: dto.Author, ISBN: dto.ISBN, Published: dto.Published}
	return global.DB.Create(&bookData).Error
}

func (b BookService) List(name string) ([]model.BookData, error) {
	var bookDataList []model.BookData

	db := global.DB.Model(&model.BookData{})

	if name != "" {
		db.Where("name LIKE ?", "%"+name+"%")
	}

	result := db.Find(&bookDataList)

	return bookDataList, result.Error
}

func (b BookService) Edit(id int64, dto *dto.EdtiBookDataDto) error {
	var bookData model.BookData
	tx := global.DB.Model(&model.BookData{}).First(id, &bookData)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在该图书")
	}
	fmt.Println("isbn:", dto.ISBN)
	if dto.ISBN != 0 {
		result := global.DB.Model(&model.BookData{}).Where("id !=? AND ISBN = ?", id, dto.ISBN).Find(&model.BookData{})
		if result.RowsAffected != 0 {
			return errors.New("ISBN和现有的图书冲突")
		}
		bookData.ISBN = dto.ISBN
	}
	if dto.Name != "" {
		bookData.Name = dto.Name
	}
	if dto.Author != "" {
		bookData.Author = dto.Author
	}
	if dto.Published != "" {
		bookData.Published = dto.Published
	}

	global.DB.Save(&bookData)

	return nil
}

func (b BookService) Delete(id int64) error {
	return global.DB.Model(&model.BookData{}).Delete(id).Error
}
