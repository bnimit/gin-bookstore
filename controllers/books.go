package controllers

import (
	"gin-bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.statusOK, gin.H{
		"data": books,
	})
}
