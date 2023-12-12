package handler

import (
	"fmt"
	"golang-web/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "fani alfirdaus",
		"bio":  "learn to be a full-stack developer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":     "hello world",
		"sub title": "learn golang with agung setiawan",
	})
}
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
func PostBooksHandler(c *gin.Context) {
	var bookInput model.BookInput

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		log.Println(err.Error())

		errMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			log.Println(errMessage)
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
