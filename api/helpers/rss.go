package helpers

import (
	"api/model"
	"fmt"
	"log"

	"github.com/gorilla/feeds"
)

func CreateRssFeed(podcast *model.Podcast) string {
	feed := &feeds.Feed{
		Title:       podcast.Title,
		Link:        &feeds.Link{Href: "https://tyrelchambers.com"},
		Description: podcast.Description,
		Author:      &feeds.Author{Name: "Tyrel Chambers"},
	}

	if podcast.DisplayEmailInRSS {
		feed.Author.Email = podcast.OwnerEmail
	}

	var feedItems []*feeds.Item

	fmt.Println("here")

	for _, episode := range podcast.Episodes {
		fmt.Println(episode.PublishDate)

		feedItems = append(feedItems, &feeds.Item{
			Title:       episode.Title,
			Link:        &feeds.Link{Href: episode.URL},
			Description: episode.Description,
			Author:      &feeds.Author{Name: episode.Author},
			// Created:     "",
		})
	}

	json, err := feed.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	return json
}
