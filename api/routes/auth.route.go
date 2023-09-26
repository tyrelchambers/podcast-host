package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(c echo.Context) error {

	db := helpers.DbClient()

	var u model.User

	err := c.Bind(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userExists := models.CheckUserExists(u.Email, db)

	if userExists == true {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists.")
	}

	newUser, e := models.CreateUser(&u, db)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	cookieValues := model.Cookie{
		UserID: newUser.ID,
	}

	helpers.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, newUser)
}

func Login(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")

	userExists := models.CheckUserExists(email, helpers.DbClient())

	if userExists {
		return echo.NewHTTPError(http.StatusBadRequest, "User exists.")
	}

	user, err := models.FindUserByEmail(email, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	diffPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if diffPassword != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password.")
	}

	cookieValues := model.Cookie{
		UserID: user.ID,
	}

	helpers.SessionHandler(c, cookieValues)

	return c.JSON(http.StatusOK, user)
}
