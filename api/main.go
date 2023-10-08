package main

import (
	"api/constants"
	"api/helpers"
	"api/routes"
	"log"
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	helpers.InitDb()

	store, err := pgstore.NewPGStore(constants.DbUrl, []byte("secret-key"))

	if err != nil {
		log.Fatalf(err.Error())
	}

	e := echo.New()

	e.Use(session.Middleware(store))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		ExposeHeaders:    []string{echo.HeaderContentType, echo.HeaderAccept, echo.HeaderCookie, echo.HeaderAccessControlAllowCredentials},
	}))

	e.POST("/api/episode/create", routes.CreateEpisode)
	e.POST("/api/episode/:id/edit", routes.UpdateEpisode)
	e.DELETE("/api/episode/:id/delete", routes.DeleteEpisode)
	e.GET("/api/episode/:id", routes.GetEpisode)
	e.POST("/api/auth/register", routes.AuthHandler)
	e.POST("/api/auth/login", routes.Login)

	e.GET("/api/user/me", routes.GetCurrentUser)

	e.GET("/api/user/:id/podcasts", routes.GetUserPodcasts)

	e.POST("/api/podcast/create", routes.CreatePodcast)
	e.GET("/api/podcast/:name", routes.GetPodcastSettings)
	e.GET("/api/podcast/:name/episodes", routes.GetPodcastEpisodes)
	e.POST("/api/podcast/:name/edit", routes.UpdatePodcast)

	e.GET("/api/podcast/:id/info", routes.InfoRoute)

	err = http.ListenAndServe(":8080", e)

	if err != nil {
		log.Fatal(err)
	}
}
