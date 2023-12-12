package model

import "time"

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

type Book struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rating      float64
	CreatedAt   time.Time
	UpdateAt    time.Time
}
