package model

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updateAt`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleteAt"`
}

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = gorm.Open(sqlite.Open("book.db"))
	if err != nil {
		panic(fmt.Sprintf("open sqlite file(book.db) failed: %s", err.Error()))
	}

	initDB()

	return db
}
