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

	podcast.UserID = user.UUID
	podcast.Title = c.FormValue("title")
	podcast.Description = c.FormValue("description")
	podcast.Thumbnail = c.FormValue("thumbnail")
	podcast.ExplicitContent = c.FormValue("explicit_content") == "true"
	podcast.PrimaryCategory = c.FormValue("primary_category")
	podcast.SecondaryCategory = c.FormValue("secondary_category")
	podcast.Author = c.FormValue("author")
	podcast.Copyright = c.FormValue("copyright")
	podcast.Keywords = c.FormValue("keywords")
	podcast.Website = c.FormValue("website")
	podcast.Language = c.FormValue("language")
	podcast.Timezone = c.FormValue("timezone")
	podcast.ShowOwner = c.FormValue("show_owner")
	podcast.OwnerEmail = c.FormValue("owner_email")
	podcast.DisplayEmailInRSS = c.FormValue("display_email_in_rss") == "true"

	err := models.CreatePodcast(&podcast, helpers.DB())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to create podcast.")
	}

	return c.JSON(http.StatusOK, podcast)
}

func GetUserPodcasts(c echo.Context) error {
	user := sessions.GetUserFromSession(c)

	podcasts, err := models.GetUsersPodcasts(user.UUID, helpers.DB())

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

	id := c.Param("id")

	podcast, err := models.GetEpisodes(id, user.UUID, helpers.DB())

	latestEpisodeData, err := models.GetLatestEpisodeByPodcast(id, helpers.DB())

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

	id := c.Param("id")

	episodes, err := models.GetPodcastEpisodesById(id, helpers.DB())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to get episodes.")
	}

	return c.JSON(http.StatusOK, episodes)
}

func UpdatePodcast(c echo.Context) error {
	var body model.PodcastDTO
	user := sessions.GetUserFromSession(c)

	body.UserID = user.UUID
	body.Title = c.FormValue("title")
	body.Description = c.FormValue("description")
	body.Thumbnail = c.FormValue("thumbnail")
	body.ExplicitContent = c.FormValue("explicit_content") == "true"
	body.PrimaryCategory = c.FormValue("primary_category")
	body.SecondaryCategory = c.FormValue("secondary_category")
	body.Author = c.FormValue("author")
	body.Copyright = c.FormValue("copyright")
	body.Keywords = c.FormValue("keywords")
	body.Website = c.FormValue("website")
	body.Language = c.FormValue("language")
	body.Timezone = c.FormValue("timezone")
	body.ShowOwner = c.FormValue("show_owner")
	body.OwnerEmail = c.FormValue("owner_email")
	body.DisplayEmailInRSS = c.FormValue("display_email_in_rss") == "true"
	body.UUID = c.Param("id")

	podcastModel := *body.ToEntity()

	podcast, err := models.UpdatePodcast(&podcastModel, user.UUID, helpers.DB())

	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Failed to update podcast.")
	}

	return c.JSON(http.StatusOK, podcast.ToDTO())

}
