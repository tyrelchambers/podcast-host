package routes

import (
	sessions "api/session"

	"github.com/labstack/echo/v4"
)

func GetCurrentUser(c echo.Context) error {

	sessions.ReadCookie(c)

	// userId, err := sessions.ReadCookieHandler(c)

	// if err != nil {
	// 	return err
	// }

	// db := helpers.DbClient()

	// fmt.Println("---->", userId)

	// user, err := models.GetUser(&userId, db)

	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Can't find user from cookie")
	// }

	// return c.JSON(http.StatusOK, user)
	return nil
}
