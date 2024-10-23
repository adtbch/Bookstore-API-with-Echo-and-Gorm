// cmd/main.go
package main

import (
    "Weekly-Task3/pkg/config"
    "Weekly-Task3/pkg/routes"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // Inisialisasi database
    db := config.InitDB()

    // Inisialisasi Echo
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Inisialisasi routes
    routes.BookstoreRoutes(e, db)

    // Menjalankan server
    e.Logger.Fatal(e.Start(":8080"))
}
