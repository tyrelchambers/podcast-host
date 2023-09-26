package routes

import (
	"api/helpers"
	"api/models"
	"net/http"

	"github.com/labstack/echo"
)

func InfoRoute(c echo.Context) error {
	type PodcastID struct {
		PodcastID string `json:"podcastId"`
	}

	type Response struct {
		NextEpisodeNumber int    `json:"nextEpisodeNumber"`
		RssFeed           string `json:"rssFeed"`
	}

	userId, err := helpers.ReadCookieHandler(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Please provide valid credentials")

	}

	pId := c.Param("podcastId")

	count, _ := models.GetEpisodesCountAndIncrement(pId, helpers.DbClient())
	podcast, err2 := models.GetPodcastById(pId, userId, helpers.DbClient())
	feed := helpers.CreateRssFeed(&podcast)

	if err != nil && err2 != nil {

		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get episode count.")
	}

	response := Response{
		NextEpisodeNumber: count + 1,
		RssFeed:           feed,
	}

	return c.JSON(http.StatusOK, response)

}
