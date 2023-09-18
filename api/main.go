package main

import (
	"api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	SetupDb()

	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/api/episode/create", routes.CreateEpisode).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/episode/{id}", routes.GetEpisode).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/episode/{id}/edit", routes.UpdateEpisode).Methods(http.MethodPut, http.MethodOptions)

	r.HandleFunc("/api/auth/register", routes.AuthHandler).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/me", routes.GetCurrentUser).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/user/{id}/episodes", routes.GetUserEpisodes).Methods(http.MethodGet, http.MethodOptions)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
