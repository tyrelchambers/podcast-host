package models

import (
	"api/model"
	"fmt"

	"gorm.io/gorm"
)

func CreateEpisode(episode *model.Episode, db *gorm.DB) (e error) {
	db.Create(episode)

	fmt.Println("SUCCESS: new episode created")

	return

}

func GetEpisodeById(id string, db *gorm.DB) (model.Episode, error) {
	var episode model.Episode

	db.First(&episode, "id = ?", id)

	return episode, nil
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
	var podcast model.PodcastDTO
	var episode model.Episode

	db.First(&podcast, "id = ?", podcastID)

	db.Where("podcast_id = ?", podcastID).Order("episode_number DESC").First(&episode)

	return episode, nil
}

func GetEpisodesCountAndIncrement(podcastId string, db *gorm.DB) (int64, error) {

	var count int64

	db.Model(&model.Episode{}).Where("podcast_id = ?", podcastId).Count(&count)

	return count, nil
}
