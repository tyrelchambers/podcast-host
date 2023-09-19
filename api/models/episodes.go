package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
)

func CreateEpisode(episode *Episode, db *sql.DB) (e error) {
	cmd := `INSERT INTO Episodes (id, title, description, url, user_id, keywords, publishDate, author, episodeNumber) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	id := cuid.New()

	_, err := db.Exec(cmd, id, episode.Title, episode.Description, episode.URL, episode.UserID, episode.Keywords, episode.PublishDate, episode.Author, episode.EpisodeNumber)

	if err != nil {
		return errors.New("Failed to create episode.")
	}

	fmt.Println("SUCCESS: new episode created")

	return

}

func GetEpisode(id string, db *sql.DB) (episode Episode, e error) {
	cmd := `SELECT id, title, description, url, user_id, keywords, publishDate, author, episodeNumber FROM Episodes WHERE id = $1`

	row := db.QueryRow(cmd, id)

	err := row.Scan(&episode.ID, &episode.Title, &episode.Description, &episode.URL, &episode.UserID, &episode.Keywords, &episode.PublishDate, &episode.Author, &episode.EpisodeNumber)

	if err != nil {
		println(err.Error())
		return episode, errors.New("Failed to get episode.")
	}

	return episode, nil
}

func UpdateEpisode(episode Episode, db *sql.DB) (e error) {

	if episode.ID == "" {
		return errors.New("Failed to update episode. Missing ID.")
	}

	cmd := `UPDATE Episodes SET title = $1, description = $2, url = $3, user_id = $4, keywords = $5, author = $6, episodeNumber = $7 WHERE id = $8`

	res, err := db.Exec(cmd, episode.Title, episode.Description, episode.URL, episode.UserID, episode.Keywords, episode.Author, episode.EpisodeNumber, episode.ID)

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
