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
		NextEpisodeNumber int    `json:"nextEpisodeNumber"`
		RssFeed           string `json:"rssFeed"`
	}

	userId := helpers.ReadCookieHandler(w, r)

	vars := mux.Vars(r)
	pId := vars["id"]

	c, err := models.GetEpisodesCountAndIncrement(pId, helpers.DbClient())
	podcast, err2 := models.GetPodcastById(pId, userId, helpers.DbClient())
	feed := helpers.CreateRssFeed(&podcast)

	if err != nil && err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		NextEpisodeNumber: c + 1,
		RssFeed:           feed,
	}

	json.NewEncoder(w).Encode(response)

}
