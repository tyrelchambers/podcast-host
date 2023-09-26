package helpers

import (
	"api/model"
	"log"
	"net/http"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

// Hash keys should be at least 32 bytes long
var hashKey = []byte("ikiQQVPZlgwH3J7c")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("ikiQQVPZlgwH3J7c")
var s = securecookie.New(hashKey, blockKey)
var Hi = "sh"

func SessionHandler(c echo.Context, values model.Cookie) error {
	store, err := pgstore.NewPGStore(DbUrl, []byte("secret-key"))

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer store.Close()

	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	session, _ := store.Get(c.Request(), "session-key")

	// Set some session values.
	session.Values["user_id"] = values.UserID
	// session.Options.Domain = "localhost"

	// Save it before we write to the response/return from the handler.
	err = session.Save(c.Request(), c.Response())
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	SetCookieHandler(values.UserID, c)

	return nil
}

func SetCookieHandler(value string, c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func ReadCookieHandler(c echo.Context) (userId string, err error) {
	cookie, err := c.Cookie("session-key")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
