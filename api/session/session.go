package sessions

import (
	"api/model"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Hash keys should be at least 32 bytes long
var hashKey = []byte("ikiQQVPZlgwH3J7c")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("ikiQQVPZlgwH3J7c")
var s = securecookie.New(hashKey, blockKey)

func SessionHandler(c echo.Context, values model.Cookie) error {
	sess, _ := session.Get("session-key", c)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		// HttpOnly: true,
	}
	sess.Values["user_id"] = values.UserID
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func GetUserId(c echo.Context) error {
	sess, _ := session.Get("session-key", c)

	sess.Values["foo"] = "bar"
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func ReadCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}
