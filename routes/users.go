package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Sign in")
}

func SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Sign up")
}
