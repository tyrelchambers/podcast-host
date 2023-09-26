package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

func CreateEpisode(c echo.Context) error {
	e := make(chan error)

	var episode model.Episode

	pId := c.FormValue("podcastId")

	episode.Title = c.FormValue("title")
	episode.Description = c.FormValue("description")
	episode.Keywords = c.FormValue("keywords")
	episode.PublishDate = c.FormValue("publishDate")
	episode.Author = c.FormValue("author")
	episode.EpisodeNumber = c.FormValue("episodeNumber")
	episode.PodcastId = pId
	episode.Draft = c.FormValue("draft") == "true"

	file, _ := c.FormFile("file")

	if file != nil {
		src, _ := file.Open()
		uploadPathUrl := fmt.Sprintf("/%s/%s.mp3", pId, file.Filename)

		go helpers.WriteFileAndUpload(c, src, pId, file.Filename)

		episode.URL = fmt.Sprintf("%s%s", helpers.BUNNY_URL_BASE, uploadPathUrl)

		defer src.Close()

	}

	newEpError := models.CreateEpisode(&episode, helpers.DbClient())

	if newEpError != nil {

		return echo.NewHTTPError(500, newEpError.Error())
	}

	close(e)

	return nil
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

func DeleteEpisode(c echo.Context) error {

	id := c.Param("id")

	err := models.DeleteEpisode(id, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.String(200, "Success")
}
