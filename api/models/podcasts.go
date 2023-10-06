package models

import (
	"api/helpers"
	"api/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func CreatePodcast(p *model.Podcast, db *gorm.DB) error {

	db.Create(p)

	fmt.Println("SUCCESS: new podcast created")

	return nil
}

func GetUsersPodcasts(userId string, db *gorm.DB) ([]model.PodcastDTO, error) {
	var podcasts []model.Podcast
	var podcastsDto []model.PodcastDTO

	db.Preload("Episodes").Find(&podcasts, "user_id = ?", userId)

	for _, v := range podcasts {
		var dto model.PodcastDTO

		helpers.ConvertToDto(v, &dto)

		podcastsDto = append(podcastsDto, dto)
	}

	fmt.Println(podcasts[0].Episodes)

	return podcastsDto, nil
}

func GetPodcastById(id string, userId string, db *gorm.DB) (p model.Podcast, e error) {
	var podcast model.Podcast

	db.First(&podcast, "id = ? AND user_id = ?", id, userId)

	return podcast, nil
}

func GetPodcastByNameWithEpisodes(name string, userId string, db *gorm.DB) (p model.Podcast, e error) {

	var podcast model.Podcast

	parsedName := strings.Replace(name, "-", " ", -1)

	db.First(&podcast, "title = ? AND user_id = ?", parsedName, userId)

	return podcast, nil

}

func GetPodcastIdFromName(name string, db *gorm.DB) (string, error) {
	var podcast model.Podcast

	parsedName := strings.Replace(name, "-", " ", -1)

	db.First(&podcast, "title = ?", parsedName)

	return podcast.UUID, nil
}

func GetPodcastEpisodesById(id string, db *gorm.DB) ([]model.EpisodeDTO, error) {
	var episodes []model.Episode
	var eDtos []model.EpisodeDTO

	db.Find(&episodes, "podcast_id = ?", id)

	for _, v := range episodes {
		var dto model.EpisodeDTO

		helpers.ConvertToDto(v, &dto)

		eDtos = append(eDtos, dto)
	}

	return eDtos, nil

}

func UpdatePodcast(podcast *string, db *gorm.DB) error {

	db.Model(&podcast).Updates(podcast)

	fmt.Printf("SUCCESS: updated %d podcast\n", db.RowsAffected)

	return nil
}
