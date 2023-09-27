package routes

import (
	sessions "api/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCurrentUser(c echo.Context) error {

	user := sessions.GetUserFromSession(c)

	return c.JSON(http.StatusOK, user)
}
