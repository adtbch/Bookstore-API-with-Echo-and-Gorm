// pkg/models/book.go
package models

import "gorm.io/gorm"

// Book represents a book in the system
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}
