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

func CreateEpisode(w http.ResponseWriter, r *http.Request) {
	e := make(chan error)

	r.ParseMultipartForm(3 << 20)

	var episode model.Episode

	pId := r.FormValue("podcastId")

	episode.Title = r.FormValue("title")
	episode.Description = r.FormValue("description")
	episode.Keywords = r.FormValue("keywords")
	episode.PublishDate = r.FormValue("publishDate")
	episode.Author = r.FormValue("author")
	episode.EpisodeNumber = r.FormValue("episodeNumber")
	episode.PodcastId = pId

	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	uploadPathUrl := fmt.Sprintf("/%s/%s.mp3", pId, header.Filename)

	go helpers.WriteFileAndUpload(r, e, file, pId, header.Filename)

	episode.URL = fmt.Sprintf("%s%s", helpers.BUNNY_URL_BASE, uploadPathUrl)

	newEpError := models.CreateEpisode(&episode, helpers.DbClient())

	if newEpError != nil {
		http.Error(w, newEpError.Error(), http.StatusInternalServerError)
		return
	}

	close(e)
}

func GetEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	p := mux.Vars(r)

	id := p["id"]

	episode, err := models.GetEpisodeById(id, helpers.DbClient())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(episode)
}

func UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }

	// e := make(chan error)

	// r.ParseMultipartForm(3 << 20)

	// var episode model.Episode
	// var uploadPath = r.FormValue("url")

	// file, _, noFile := r.FormFile("file")

	// if noFile == nil {
	// 	p := helpers.WriteFileAndUpload(r, e, file)
	// 	uploadPath = p
	// }

	// episode.Title = r.FormValue("title")
	// episode.Description = r.FormValue("description")
	// episode.Keywords = r.FormValue("keywords")
	// episode.PublishDate = r.FormValue("publishDate")
	// episode.Author = r.FormValue("author")
	// episode.EpisodeNumber = r.FormValue("episodeNumber")
	// episode.ID = r.FormValue("id")
	// episode.URL = uploadPath
	// episode.PodcastId = r.FormValue("podcastId")

	// err := models.UpdateEpisode(episode, helpers.DbClient())

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// close(e)

}

func DeleteEpisode(w http.ResponseWriter, r *http.Request) {

	p := mux.Vars(r)

	id := p["id"]

	err := models.DeleteEpisode(id, helpers.DbClient())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
