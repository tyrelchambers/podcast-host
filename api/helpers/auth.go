package helpers

import (
	"api/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/securecookie"
)

// Hash keys should be at least 32 bytes long
var hashKey = []byte("ikiQQVPZlgwH3J7c")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("ikiQQVPZlgwH3J7c")
var s = securecookie.New(hashKey, blockKey)

func SessionHandler(w http.ResponseWriter, r *http.Request, values models.Cookie) {
	store, err := pgstore.NewPGStore(DbUrl, []byte("secret-key"))

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer store.Close()

	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	session, _ := store.Get(r, "session-key")

	// Set some session values.
	session.Values["user_id"] = values.UserID
	// session.Options.Domain = "localhost"

	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SetCookieHandler(values.UserID, w)

}

func SetCookieHandler(value string, w http.ResponseWriter) {
	if encoded, err := s.Encode("session-key", value); err == nil {
		cookie := &http.Cookie{
			Name:   "session-key",
			Value:  encoded,
			Path:   "/",
			Secure: true,
			// HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) (userId string) {
	var str string
	if cookie, err := r.Cookie("session-key"); err == nil {
		err = s.Decode("session-key", cookie.Value, &str)

		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	return str
}
