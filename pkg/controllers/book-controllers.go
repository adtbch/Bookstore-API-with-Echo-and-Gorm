// pkg/controllers/book-controllers.go
package controllers

import (
	"net/http"
	"Weekly-Task3/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetBooks retrieves all books
func GetBooks(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	return c.JSON(http.StatusOK, books)
}

// CreateBook creates a new book
func CreateBook(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	db.Create(&book)
	return c.JSON(http.StatusOK, book)
}

// GetBookByID retrieves a book by ID
func GetBookByID(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book
func UpdateBook(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	db.Save(&book)
	return c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book
func DeleteBook(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}
	db.Delete(&book)
	return c.JSON(http.StatusOK, echo.Map{"message": "Book deleted successfully"})
}
