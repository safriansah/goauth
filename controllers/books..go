package controllers

import (
	"fmt"
	"goauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	fmt.Println("masuk controller")
	var books []models.Book
	books = append(books, models.Book{
		ID:     1,
		Title:  "test",
		Author: "sepatu",
	})
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	book := models.Book{
		ID:     1,
		Title:  input.Title,
		Author: input.Author,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
