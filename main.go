package main

import (
	"fmt"
	"golang-web/handler"
	"golang-web/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Debug()
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatal(err.Error())
	}

	// book := model.Book{
	// 	Id:          2,
	// 	Title:       "atomic habitc",
	// 	Description: "ini adalah buku yang sangat bagus dari eka kurniawan",
	// 	Price:       120000,
	// 	Rating:      5,
	// 	CreatedAt:   time.Now().UTC(),
	// }
	var book model.Book
	if result := db.First(&book); result.Error != nil {
		log.Fatal(result.Error.Error())
	} else {
		fmt.Println(book)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
