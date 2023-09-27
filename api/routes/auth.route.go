package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	sessions "api/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(c echo.Context) error {

	db := helpers.DbClient()

	var u model.User

	err := c.Bind(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newUser, e := models.CreateUser(&u, db)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	cookieValues := model.Cookie{
		UserID: newUser.ID,
	}

	sessions.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, newUser)
}

func Login(c echo.Context) error {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Body
	err := c.Bind(&body)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
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
		UserID: user.ID,
	}

	sessions.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, "")
}
