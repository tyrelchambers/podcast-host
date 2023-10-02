package routes

import (
	"api/helpers"
	"api/model"
	"api/models"
	sessions "api/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePodcast(c echo.Context) error {
	var podcast model.PodcastDTO

	user := sessions.GetUserFromSession(c)

	err := (&echo.DefaultBinder{}).BindBody(c, &podcast)

	podcast.UserID = user.UUID

	// err = models.CreatePodcast(&podcast, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to create podcast.")
	}

	return c.JSON(http.StatusOK, podcast)
}

func GetUserPodcasts(c echo.Context) error {
	user := sessions.GetUserFromSession(c)

	podcasts, err := models.GetUsersPodcasts(user.UUID, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get podcasts.")
	}

	return c.JSON(http.StatusOK, podcasts)
}

func GetPodcastSettings(c echo.Context) error {
	type Body struct {
		Podcast       model.PodcastDTO `json:"podcast"`
		LatestEpisode model.Episode    `json:"latestEpisode"`
	}

	var rBody Body

	user := sessions.GetUserFromSession(c)

	name := c.Param("name")

	podcast, err := models.GetPodcastByNameWithEpisodes(name, user.UUID, helpers.DbClient())

	latestEpisodeData, err := models.GetLatestEpisodeByPodcast(podcast.UUID, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get latest episode.")
	}

	rBody = Body{
		Podcast:       podcast,
		LatestEpisode: latestEpisodeData,
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get latest episode.")
	}

	return c.JSON(http.StatusOK, rBody)
}

func GetPodcastEpisodes(c echo.Context) error {

	pName := c.Param("name")

	podcastId, err := models.GetPodcastIdFromName(pName, helpers.DbClient())

	episodes, err := models.GetPodcastEpisodesById(podcastId, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get episodes.")
	}

	return c.JSON(http.StatusOK, episodes)
}

func UpdatePodcast(c echo.Context) error {
	pName := c.Param("name")

	podcast, err := models.GetPodcastIdFromName(pName, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get podcast.")
	}

	err = models.UpdatePodcast(&podcast, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to update podcast.")
	}

	return c.JSON(http.StatusOK, "")

}
