package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
    gorm.Model
    BookID   uint64 `json:"book_id" gorm:"type:bigint"`
    Name     string `json:"name"`
    Title    string `json:"title"`
    Author   string  `json:"author"`
    Borrowed bool   `json:"borrowed" gorm:"default:false"`
}

type Borrowed struct {
    gorm.Model
    BorrowedID  uuid.UUID   `json:"borrowed_id" gorm:"type:uuid;"`
    BookID      uint64   `json:"book_id" gorm:"type:bigint"`
    StartTime   time.Time     `json:"startTimestamp"`
    EndTime     time.Time     `json:"endTimestamp"`
    Returned    bool       `json:"returned" gorm:"default:false"`
    ReturnedAt  time.Time     `json:"returnedAt"`
}