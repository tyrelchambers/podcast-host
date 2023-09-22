package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePodcast(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(3 << 20)

	var podcast model.Podcast

	userId := helpers.ReadCookieHandler(w, r)

	podcast.Thumbnail = r.FormValue("thumbnail")
	podcast.Title = r.FormValue("title")
	podcast.Description = r.FormValue("description")
	podcast.ExplicitContent = r.FormValue("explicitContent") == "true"
	podcast.PrimaryCategory = r.FormValue("primaryCategory")
	podcast.SecondaryCategory = r.FormValue("secondaryCategory")
	podcast.Author = r.FormValue("author")
	podcast.Copyright = r.FormValue("copyright")
	podcast.Keywords = r.FormValue("keywords")
	podcast.Website = r.FormValue("website")
	podcast.Language = r.FormValue("language")
	podcast.Timezone = r.FormValue("timezone")
	podcast.ShowOwner = r.FormValue("showOwner")
	podcast.OwnerEmail = r.FormValue("ownerEmail")
	podcast.DisplayEmailInRSS = r.FormValue("displayEmailInRSS") == "true"
	podcast.UserID = userId

	models.CreatePodcast(&podcast, helpers.DbClient())

	fmt.Println(podcast)
}

func GetUserPodcasts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userId := helpers.ReadCookieHandler(w, r)

	podcasts, err := models.GetUsersPodcasts(userId, helpers.DbClient())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(podcasts)
}

func GetPodcast(w http.ResponseWriter, r *http.Request) {
	userId := helpers.ReadCookieHandler(w, r)

	name := mux.Vars(r)["name"]

	podcast, err := models.GetPodcastByName(name, userId, helpers.DbClient())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(podcast)
}
