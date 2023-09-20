package routes

import (
	"api/helpers"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userId := helpers.ReadCookieHandler(w, r)

	db := helpers.DbClient()

	user, err := models.GetUser(&userId, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func GetUserEpisodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)

	id := vars["id"]

	db := helpers.DbClient()

	episodes, err := models.GetEpisodes(id, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(episodes)
}
