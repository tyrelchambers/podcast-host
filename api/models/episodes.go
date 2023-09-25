package models

import (
	"api/model"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
)

func CreateEpisode(episode *model.Episode, db *sql.DB) (e error) {
	cmd := `INSERT INTO Episodes (id, title, description, url, keywords, publish_date, author, episode_number, podcast_id,draft) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	id := cuid.New()

	_, err := db.Exec(cmd, id, episode.Title, episode.Description, episode.URL, episode.Keywords, episode.PublishDate, episode.Author, episode.EpisodeNumber, episode.PodcastId, episode.Draft)

	if err != nil {
		println(err.Error())
		return errors.New("Failed to create episode.")
	}

	fmt.Println("SUCCESS: new episode created")

	return

}

func GetEpisodeById(id string, db *sql.DB) (episode model.Episode, e error) {
	cmd := `SELECT id, title, description, url, podcast_id, keywords, publishDate, author, episode_number FROM Episodes WHERE id = $1`

	row := db.QueryRow(cmd, id)

	err := row.Scan(&episode.ID, &episode.Title, &episode.Description, &episode.URL, &episode.PodcastId, &episode.Keywords, &episode.PublishDate, &episode.Author, &episode.EpisodeNumber)

	if err != nil {
		println(err.Error())
		return episode, errors.New("Failed to get episode.")
	}

	return episode, nil
}

func UpdateEpisode(episode model.Episode, db *sql.DB) (e error) {

	if episode.ID == "" {
		return errors.New("Failed to update episode. Missing ID.")
	}

	cmd := `UPDATE Episodes SET title = $1, description = $2, url = $3, keywords = $5, author = $6, episode_number = $7 WHERE id = $8`

	res, err := db.Exec(cmd, episode.Title, episode.Description, episode.URL, episode.Keywords, episode.Author, episode.EpisodeNumber, episode.ID)

	if err != nil {
		println(err.Error())
		return errors.New("Failed to update episode.")
	}

	count, err := res.RowsAffected()
	if err != nil {

		return errors.New(fmt.Sprintf("Failed to update episode. Error: %s", err.Error()))
	}

	fmt.Printf("SUCCESS: updated %d episode\n", count)
	return

}

func DeleteEpisode(id string, db *sql.DB) (e error) {
	cmd := `DELETE FROM Episodes WHERE id = $1`

	res, err := db.Exec(cmd, id)

	if err != nil {
		println(err.Error())
		return errors.New("Failed to delete episode.")
	}

	count, err := res.RowsAffected()

	if err != nil {
		return errors.New(fmt.Sprintf("Failure reading rows. Error: %s", err.Error()))
	}

	fmt.Printf("SUCCESS: deleted %d episode\n", count)
	return
}

func GetLatestEpisodeByPodcast(podcastID string, db *sql.DB) (episode model.Episode, e error) {
	cmd := `SELECT id, title, url, publish_date, episode_number FROM Episodes WHERE podcast_id = $1 ORDER BY publish_date DESC LIMIT 1`

	row := db.QueryRow(cmd, podcastID)

	row.Scan(&episode.ID, &episode.Title, &episode.URL, &episode.PublishDate, &episode.EpisodeNumber)

	return episode, nil
}

func GetEpisodesCountAndIncrement(podcastId string, db *sql.DB) (c int, e error) {

	cmd := `SELECT COUNT(*) FROM Episodes WHERE podcast_id = $1`

	row := db.QueryRow(cmd, podcastId)

	var count int

	err := row.Scan(&count)

	if err != nil {
		println(err.Error())
		return count, errors.New("Failed to get episode count.")
	}

	return count, nil
}
