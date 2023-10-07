package routes

import (
	"api/constants"
	"api/helpers"
	"api/model"
	"api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateEpisode(c echo.Context) error {
	e := make(chan error)

	var episode model.Episode

	pId := c.FormValue("podcastId")

	convertedDate := helpers.ConvertToUnix(c.FormValue("publishDate"))
	convertedEpNum, err := strconv.ParseUint(c.FormValue("episodeNumber"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	episode.Title = c.FormValue("title")
	episode.Description = c.FormValue("description")
	episode.Keywords = c.FormValue("keywords")
	episode.PublishDate = convertedDate
	episode.Author = c.FormValue("author")
	episode.EpisodeNumber = convertedEpNum
	episode.PodcastId = pId
	episode.Draft = c.FormValue("draft") == "true"

	file, _ := c.FormFile("file")

	if file != nil {
		src, _ := file.Open()
		uploadPathUrl := fmt.Sprintf("/%s/%s.mp3", pId, file.Filename)

		go helpers.WriteFileAndUpload(c, src, pId, file.Filename)

		episode.URL = fmt.Sprintf("%s%s", constants.BUNNY_URL_BASE, uploadPathUrl)

		defer src.Close()

	}

	newEpError := models.CreateEpisode(&episode, helpers.DbClient())

	if newEpError != nil {

		return echo.NewHTTPError(500, newEpError.Error())
	}

	close(e)

	return nil
}

func GetEpisode(c echo.Context) error {

	id := c.Param("id")

	episode, err := models.GetEpisodeById(id, helpers.DbClient())

	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusOK, episode)
}

func UpdateEpisode(c echo.Context) error {
	e := make(chan error)
	var episode model.Episode

	var uploadPath = c.FormValue("url")

	var convertedDate uint64

	if c.FormValue("publishDate") != "" {
		convertedDate = helpers.ConvertToUnix(c.FormValue("publishDate"))

	}

	convertedEpNum, err := strconv.ParseUint(c.FormValue("episodeNumber"), 10, 64)

	episode.Title = c.FormValue("title")
	episode.Description = c.FormValue("description")
	episode.Keywords = c.FormValue("keywords")
	episode.PublishDate = convertedDate
	episode.Author = c.FormValue("author")
	episode.EpisodeNumber = convertedEpNum
	episode.UUID = c.FormValue("id")
	episode.URL = uploadPath
	episode.PodcastId = c.FormValue("podcastId")

	file, _ := c.FormFile("file")

	if episode.PodcastId == "" {
		return echo.NewHTTPError(500, "Missing podcastId")
	}

	if file != nil {
		src, _ := file.Open()
		uploadPathUrl := fmt.Sprintf("/%s/%s", episode.PodcastId, helpers.ConvertToMp3(file.Filename))

		go helpers.WriteFileAndUpload(c, src, episode.PodcastId, file.Filename)

		episode.URL = fmt.Sprintf("%s%s", constants.BUNNY_URL_BASE, uploadPathUrl)

		defer src.Close()

	}

	err = models.UpdateEpisode(episode, helpers.DbClient())

	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(500, err.Error())
	}

	close(e)

	return nil
}

func DeleteEpisode(c echo.Context) error {

	id := c.Param("id")

	err := models.DeleteEpisode(id, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.String(200, "Success")
}
