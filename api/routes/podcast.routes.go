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
	var podcast model.Podcast

	user := sessions.GetUserFromSession(c)

	podcast.Thumbnail = c.FormValue("thumbnail")
	podcast.Title = c.FormValue("title")
	podcast.Description = c.FormValue("description")
	podcast.ExplicitContent = c.FormValue("explicitContent") == "true"
	podcast.PrimaryCategory = c.FormValue("primaryCategory")
	podcast.SecondaryCategory = c.FormValue("secondaryCategory")
	podcast.Author = c.FormValue("author")
	podcast.Copyright = c.FormValue("copyright")
	podcast.Keywords = c.FormValue("keywords")
	podcast.Website = c.FormValue("website")
	podcast.Language = c.FormValue("language")
	podcast.Timezone = c.FormValue("timezone")
	podcast.ShowOwner = c.FormValue("showOwner")
	podcast.OwnerEmail = c.FormValue("ownerEmail")
	podcast.DisplayEmailInRSS = c.FormValue("displayEmailInRSS") == "true"
	podcast.UserID = user.ID

	// err = models.CreatePodcast(&podcast, helpers.DbClient())

	return c.JSON(http.StatusOK, podcast)
}

func GetUserPodcasts(c echo.Context) error {
	user := sessions.GetUserFromSession(c)

	podcasts, err := models.GetUsersPodcasts(user.ID, helpers.DbClient())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get podcasts.")
	}

	return c.JSON(http.StatusOK, podcasts)
}

func GetPodcastSettings(c echo.Context) error {
	type Body struct {
		Podcast       model.Podcast `json:"podcast"`
		LatestEpisode model.Episode `json:"latestEpisode"`
	}

	var rBody Body

	user := sessions.GetUserFromSession(c)

	name := c.Param("name")

	podcast, err := models.GetPodcastByNameWithEpisodes(name, user.ID, helpers.DbClient())

	latestEpisodeData, err := models.GetLatestEpisodeByPodcast(podcast.ID, helpers.DbClient())

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

	podcast, err := models.GetPodcastIdFromName(pName, helpers.DbClient())

	episodes, err := models.GetPodcastEpisodesById(podcast.ID, helpers.DbClient())

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
