package routes

import (
	"api/helpers"
	"api/models"
	sessions "api/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InfoRoute(c echo.Context) error {
	type PodcastID struct {
		PodcastID string `json:"podcastId"`
	}

	type Response struct {
		NextEpisodeNumber string `json:"next_episode_number"`
		RssFeed           string `json:"rss_feed"`
	}

	user := sessions.GetUserFromSession(c)

	pId := c.Param("id")

	count, _ := models.GetEpisodesCountAndIncrement(pId, helpers.DB())

	podcast, err := models.GetPodcastById(pId, user.UUID, helpers.DB())

	feed := helpers.CreateRssFeed(podcast)

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get episode count.")
	}

	response := Response{
		NextEpisodeNumber: count,
		RssFeed:           feed,
	}

	return c.JSON(http.StatusOK, response)

}
