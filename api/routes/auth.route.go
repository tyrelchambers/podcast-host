package routes

import (
	"api/helpers"
	"api/models"
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	db := helpers.DbClient()

	var u models.User

	json.NewDecoder(r.Body).Decode(&u)

	userExists := models.CheckUserExists(u.Email, db)

	if userExists == true {
		http.Error(w, "User already exist", http.StatusBadRequest)
		return
	}

	newUser, e := models.CreateUser(&u, db)
	// session, e := models.CreateSession(newUser, db)

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	cookieValues := models.Cookie{
		UserID: newUser.ID,
	}

	helpers.SessionHandler(w, r, cookieValues)

}
