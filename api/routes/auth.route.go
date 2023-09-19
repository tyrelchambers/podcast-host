package routes

import (
	"api/helpers"
	"api/models"
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")
	w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	db := helpers.DbClient()

	var u models.User

	userExists := models.CheckUserExists(r.FormValue("email"), db)

	if userExists {
		http.Error(w, "User already exist", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&u)

	newUser, e := models.CreateUser(&u, db)
	session, e := models.CreateSession(&newUser, db)

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	cookieValues := models.Cookie{
		SessionToken: session.SessionToken,
		ExpiresAt:    session.ExpiresAt,
	}

	helpers.SessionHandler(w, r, cookieValues)

}
