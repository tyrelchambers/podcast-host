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
var hashKey = []byte("very-secret")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("a-lot-secret")
var s = securecookie.New(hashKey, blockKey)

func SessionHandler(w http.ResponseWriter, r *http.Request, values models.Cookie) {
	store, err := pgstore.NewPGStore(DbUrl, []byte("secret-key"))

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer store.Close()

	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Set some session values.
	session.Values["session_token"] = values.SessionToken
	// session.Options.Domain = "localhost"

	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}
	}
}
