package main

import (
	"golang-web/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	router.POST("/books", handler.PostBooksHandler)

	router.Run()
}
