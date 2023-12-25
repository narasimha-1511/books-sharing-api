package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/narasimha-1511/zolo-backend/config"
	"github.com/narasimha-1511/zolo-backend/models"
)

func GetBook(c *gin.Context){
    var book []models.Book

    config.DB.Find(&book)
    

    c.JSON(200, gin.H{
        "Books": book,
    })
}

func CreateBookParams(c *gin.Context){
    var book models.Book

    book.BookID = uint64(rand.Int63())
    book.Name = c.Param("name")
    book.Title = c.Param("title")
    book.Author = c.Param("author")

	config.DB.AutoMigrate(&models.Book{})

    config.DB.Create(&book)

    c.JSON(200, gin.H{
        "book_id": book.BookID,
		"name": book.Name,
		"title": book.Title,
		"author": book.Author,
		"status": "Book Created Successfully",
    })
}

func CreateBookPostForm(c *gin.Context){
	var book models.Book

	book.BookID = uint64(rand.Int63())
	book.Name = c.PostForm("name")
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")

	config.DB.Create(&book)

	c.JSON(200, gin.H{
		"book_id": book.BookID,
		"name": book.Name,
		"title": book.Title,
		"author": book.Author,
		"status": "Book Created Successfully",
	})
}

func BorrowBook(c *gin.Context){

	var book models.Book

	book_id := c.Param("book_id")

	result := config.DB.Where("book_id = ?", book_id).First(&book)
	if result.Error != nil {
   	 c.JSON(200, gin.H{
		"message": "Invalid Book ID",
	 })
	 return
	}

	//Check if the book is already borrowed
	if book.Borrowed == true {
		c.JSON(200, gin.H{
			"message": "Book is already borrowed",
		})	
		return
	}

	config.DB.Model(&book).Update("borrowed", true)

	var borrowed models.Borrowed

	borrowed.BorrowedID = uuid.New()

	time_needed_str := c.PostForm("borrow_period") // this is in days

	if time_needed_str == "" {
		time_needed_str = "7" // if no time is given then default is 7 days
	}

	time_needed, _ := strconv.Atoi(time_needed_str)

	borrowed.BookID = book.BookID
	borrowed.StartTime = time.Now()
	borrowed.EndTime = time.Now().AddDate(0, 0, time_needed)

	config.DB.Create(&borrowed)

	c.JSON(200, gin.H{
		"status": "Book Borrowed Successfully",
		"borrowed_id": borrowed.BorrowedID,
		"book_id": borrowed.BookID,
		"start_time": borrowed.StartTime,
		"end_time": borrowed.EndTime,
		"returned": borrowed.Returned,
	})
}

func GetBorrowedBooks(c *gin.Context){
	var borrowed []models.Borrowed

	result := config.DB.Find(&borrowed)

	if result.Error != nil {
		// Handle error
		fmt.Println(result.Error)
	}

	var borrowedData []gin.H

	for _, borrow := range borrowed {
		var book models.Book
		config.DB.Where("book_id = ?", borrow.BookID).First(&book)

		borrowedData = append(borrowedData, gin.H{
			"borrowed_id": borrow.BorrowedID,
			"book_id": borrow.BookID,
			"book_name": book.Name,
			"start_time": borrow.StartTime,
			"end_time": borrow.EndTime,
			"returned": borrow.Returned,
		})
	}

	c.JSON(200, gin.H{
		"BorrowedBooks": borrowedData,
	})

}

func ReturnBook(c *gin.Context){

	var borrowed models.Borrowed
	var book models.Book

	book_id := c.Param("book_id")
	borrow_id := c.Param("borrow_id")

	result := config.DB.Where("book_id = ?", book_id).First(&book)
	if result.Error != nil {
     c.JSON(200, gin.H{
		"message": "Invalid Book ID",
	 })
	 return
	}

	result = config.DB.Where("borrowed_id = ?", borrow_id).First(&borrowed)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"message": "Invalid Borrow ID",
		})	
		return
	}

	if borrowed.Returned == true {
		c.JSON(200, gin.H{
			"message": "Book is already returned",
		})	
		return
	}

	borrowed.ReturnedAt = time.Now()

	config.DB.Model(&book).Update("borrowed", false)
	config.DB.Model(&borrowed).Update("returned", true)

	c.JSON(200, gin.H{
		"status": "Book Returned Successfully",
		"borrowed_id": borrowed.BorrowedID,
		"book_id": borrowed.BookID,
	})

}
