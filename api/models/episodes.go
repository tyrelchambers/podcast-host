package models

import (
	"api/helpers"
	"api/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

func CreateEpisode(episode *model.Episode, db *gorm.DB) (e error) {

	if episode.PodcastId == "" {
		return errors.New("Podcast ID is required")
	}

	id := cuid.New()

	ep := &model.Episode{
		UUID:          id,
		Title:         episode.Title,
		Description:   episode.Description,
		Keywords:      episode.Keywords,
		PublishDate:   episode.PublishDate,
		Author:        episode.Author,
		Draft:         episode.Draft,
		URL:           episode.URL,
		Image:         episode.Image,
		EpisodeNumber: episode.EpisodeNumber,
		PodcastId:     episode.PodcastId,
		IsScheduled:   episode.IsScheduled,
	}

	db.Create(ep)

	fmt.Println("SUCCESS: new episode created")

	return
}

func GetEpisodeById(id string, db *gorm.DB) (model.EpisodeDTO, error) {
	var episode model.Episode
	var eDto model.EpisodeDTO

	db.First(&episode, "id = ?", id)

	helpers.ConvertToDto(episode, &eDto)

	return eDto, nil
}

func UpdateEpisode(episode model.Episode, db *gorm.DB) (e error) {

	res := db.Model(&episode).Updates(episode)

	fmt.Printf("SUCCESS: updated %d episode\n", res.RowsAffected)
	return

}

func DeleteEpisode(id string, db *gorm.DB) (e error) {

	res := db.Delete(&model.Episode{}, id)

	fmt.Printf("SUCCESS: deleted %d episode\n", res.RowsAffected)
	return
}

func GetLatestEpisodeByPodcast(podcastID string, db *gorm.DB) (model.Episode, error) {
	var podcast model.Podcast
	var episode model.Episode

	db.First(&podcast, "id = ?", podcastID)

	db.Where("podcast_id = ?", podcastID).Order("episode_number DESC").First(&episode)

	return episode, nil
}

func GetEpisodesCountAndIncrement(podcastId string, db *gorm.DB) (string, error) {

	var count int64

	db.Model(&model.Episode{}).Where("podcast_id = ?", podcastId).Count(&count)

	return strconv.FormatInt(count+1, 10), nil
}
