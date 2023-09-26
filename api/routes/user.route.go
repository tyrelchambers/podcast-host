package routes

import (
	"api/helpers"
	"api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetCurrentUser(c echo.Context) error {

	userId, err := helpers.ReadCookieHandler(c)

	if err != nil {
		return err
	}

	db := helpers.DbClient()

	user, err := models.GetUser(&userId, db)

	if err != nil {

		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	return c.JSON(http.StatusOK, user)

}
