package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	sessions "api/session"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(c echo.Context) error {

	db := helpers.DbClient()

	var body model.RegisterBody

	err := (&echo.DefaultBinder{}).BindBody(c, &body)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if body.Password == "" || body.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email or password.")
	}

	newUser, e := models.CreateUser(body, db)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	cookieValues := model.Cookie{
		UserID: newUser.UUID,
	}

	fmt.Println(cookieValues)

	sessions.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, newUser)
}

func Login(c echo.Context) error {

	var body *model.RegisterBody

	err := (&echo.DefaultBinder{}).BindBody(c, &body)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := models.FindUserByEmail(body.Email, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	diffPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if diffPassword != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password.")
	}

	cookieValues := model.Cookie{
		UserID: user.UUID,
	}

	sessions.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, "")
}
