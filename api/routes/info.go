package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	sessions "api/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InfoRoute(c echo.Context) error {
	var latestEpisode *model.EpisodeDTO

	type Response struct {
		NextEpisodeNumber string            `json:"next_episode_number"`
		RssFeed           string            `json:"rss_feed"`
		EpisodeCount      int               `json:"episode_count"`
		DraftCount        int               `json:"draft_count"`
		ScheduledCount    string            `json:"scheduled_count"`
		LatestEpisode     *model.EpisodeDTO `json:"latest_episode"`
	}

	user := sessions.GetUserFromSession(c)

	pId := c.Param("id")

	count, _ := models.GetEpisodesCountAndIncrement(pId, helpers.DB())

	podcast, err := models.GetPodcastById(pId, user.UUID, helpers.DB())

	feed := helpers.CreateRssFeed(podcast)

	if len(podcast.Episodes) > 0 {
		latestEpisode = podcast.Episodes[len(podcast.Episodes)-1]
	} else {
		latestEpisode = nil
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get episode count.")
	}

	response := Response{
		NextEpisodeNumber: count,
		RssFeed:           feed,
		LatestEpisode:     latestEpisode,
		EpisodeCount:      len(podcast.GetPublishedEpisodes()),
		DraftCount:        len(podcast.GetDrafts()),
	}

	return c.JSON(http.StatusOK, response)

}
