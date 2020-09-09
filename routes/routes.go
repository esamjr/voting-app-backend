package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Mockturnal API")
	})
	// e.POST("/login", Login)
	// e.POST("/register", SignUp)

	return e
}
