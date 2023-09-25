package routes

import (
	"api/helpers"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func InfoRoute(w http.ResponseWriter, r *http.Request) {
	type PodcastID struct {
		PodcastID string `json:"podcastId"`
	}

	type Response struct {
		NextEpisodeNumber int `json:"nextEpisodeNumber"`
	}

	vars := mux.Vars(r)
	pId := vars["id"]

	c, err := models.GetEpisodesCountAndIncrement(pId, helpers.DbClient())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		NextEpisodeNumber: c + 1,
	}

	json.NewEncoder(w).Encode(response)

}
