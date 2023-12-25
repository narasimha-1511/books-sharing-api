package models

import (
	"gorm.io/gorm"
)

type Book struct {
    gorm.Model
    BookID   uint64 `gorm:"type:bigint"`
    Name     string `json:"name"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
    Borrowed bool   `json:"borrowed" gorm:"default:false"`
}
