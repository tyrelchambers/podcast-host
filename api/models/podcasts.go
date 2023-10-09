package models

import (
	"api/helpers"
	"api/model"
	"fmt"
	"strings"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

func CreatePodcast(p *model.Podcast, db *gorm.DB) error {

	id := cuid.New()

	pd := &model.Podcast{
		UUID:              id,
		Title:             p.Title,
		Description:       p.Description,
		Thumbnail:         p.Thumbnail,
		ExplicitContent:   p.ExplicitContent,
		PrimaryCategory:   p.PrimaryCategory,
		SecondaryCategory: p.SecondaryCategory,
		Author:            p.Author,
		Copyright:         p.Copyright,
		Keywords:          p.Keywords,
		Website:           p.Website,
		Language:          p.Language,
		Timezone:          p.Timezone,
		ShowOwner:         p.ShowOwner,
		OwnerEmail:        p.OwnerEmail,
		DisplayEmailInRSS: p.DisplayEmailInRSS,
		UserID:            p.UserID,
	}

	db.Create(pd)

	if db.Error != nil {
		fmt.Println("ERROR: ", db.Error)
		return db.Error
	}

	fmt.Println("SUCCESS: new podcast created")

	return nil
}

func GetUsersPodcasts(userId string, db *gorm.DB) ([]model.PodcastDTO, error) {
	var podcasts []model.Podcast
	var podcastsDto []model.PodcastDTO

	db.Preload("Episodes").Find(&podcasts, "user_id = ?", userId)

	for _, v := range podcasts {
		var dto model.PodcastDTO

		dto = *v.ToDTO()

		podcastsDto = append(podcastsDto, dto)
	}

	return podcastsDto, nil
}

func GetPodcastById(id string, userId string, db *gorm.DB) (p *model.PodcastDTO, e error) {
	var podcast model.Podcast
	var episodesDto []*model.EpisodeDTO

	db.Table("podcasts").Preload("Episodes").First(&podcast, "uuid = ? AND user_id = ?", id, userId)

	podcastDto := podcast.ToDTO()

	for _, v := range podcast.Episodes {
		var dto model.EpisodeDTO

		dto = *v.ToDTO()

		episodesDto = append(episodesDto, &dto)
	}

	podcastDto.Episodes = episodesDto

	return podcastDto, nil
}

func GetEpisodes(id string, userId string, db *gorm.DB) (p model.Podcast, e error) {

	var podcast model.Podcast

	db.First(&podcast, "uuid = ? AND user_id = ?", id, userId)

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

func UpdatePodcast(podcast *model.Podcast, db *gorm.DB) error {

	db.Model(&podcast).Updates(podcast)

	fmt.Printf("SUCCESS: updated %d podcast\n", db.RowsAffected)

	return nil
}
