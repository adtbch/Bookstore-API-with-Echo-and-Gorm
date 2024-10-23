// pkg/utils/utils.go
package utils

import "github.com/labstack/echo/v4"

// ErrorResponse handles error response formatting
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, echo.Map{"error": message})
}
