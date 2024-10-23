// pkg/routes/bookstore-routes.go
package routes

import (
	"Weekly-Task3/pkg/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookstoreRoutes(e *echo.Echo, db *gorm.DB) {
	// Provide DB connection to handlers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Define routes
	e.GET("/books", controllers.GetBooks)
	e.POST("/books", controllers.CreateBook)
	e.GET("/books/:id", controllers.GetBookByID)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)
}
