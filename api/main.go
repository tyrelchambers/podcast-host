package main

import (
	"api/routes"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	SetupDb()

	e := echo.New()

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
	// r.HandleFunc("/api/episode/{id}/edit", routes.UpdateEpisode).Methods(http.MethodPost, http.MethodOptions)
	e.DELETE("/api/episode/:id/delete", routes.DeleteEpisode)
	// r.HandleFunc("/api/episode/{id}", routes.GetEpisode).Methods(http.MethodGet, http.MethodOptions)

	e.POST("/api/auth/register", routes.AuthHandler)
	e.POST("/api/auth/login", routes.Login)

	e.GET("/api/user/me", routes.GetCurrentUser)

	e.GET("/api/user/:id/podcasts", routes.GetUserPodcasts)

	e.POST("/api/podcast/create", routes.CreatePodcast)
	e.GET("/api/podcast/:name", routes.GetPodcastSettings)
	e.GET("/api/podcast/:name/episodes", routes.GetPodcastEpisodes)

	e.GET("/api/podcast/:id/info", routes.InfoRoute)

	err := http.ListenAndServe(":8080", e)

	if err != nil {
		log.Fatal(err)
	}
}
