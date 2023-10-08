package sessions

import (
	"api/helpers"
	"api/model"
	"api/models"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SessionHandler(c echo.Context, values model.Cookie) error {
	sess, _ := session.Get("session-key", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["user_id"] = values.UserID
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func GetUserFromSession(c echo.Context) *model.UserDTO {
	s, _ := session.Get("session-key", c)

	val := s.Values["user_id"]

	if val == nil || val.(string) == "" {
		fmt.Println("ERROR: ", "user_id is nil in cookie")
		return nil
	}

	user, err := models.GetUser(val.(string), helpers.DB())

	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil
	}

	return user
}
