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

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	db := helpers.DbClient()

	var u models.User

	json.NewDecoder(r.Body).Decode(&u)

	e := models.CreateUser(&u, db)

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

}
