package controller

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/narasimha-1511/zolo-backend/config"
	"github.com/narasimha-1511/zolo-backend/models"
)

func GetBook(c *gin.Context){
    var book []models.Book

    config.DB.Find(&book)
    

    c.JSON(200, gin.H{
        "data": book,
    })
}

func CreateBookParams(c *gin.Context){
    var book models.Book

    book.BookID = uint64(rand.Int63())
    book.Name = c.Param("name")
    book.Title = c.Param("title")
    book.Author = c.Param("author")
    book.Quantity = 1

	config.DB.AutoMigrate(&models.Book{})

    config.DB.Create(&book)

    c.JSON(200, gin.H{
        "data": book,
    })
}

func CreateBookPostForm(c *gin.Context){
	var book models.Book

	book.BookID = uint64(rand.Int63())
	book.Name = c.PostForm("name")
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Quantity = 1

	config.DB.Create(&book)

	c.JSON(200, gin.H{
		"data": book,
	})
}