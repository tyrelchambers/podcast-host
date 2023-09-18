package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
)

func CreateUser(user *User, db *sql.DB) (e error) {
	cmd := `INSERT INTO Users (id, email, password) VALUES ($1, $2, $3)`

	id := cuid.New()

	_, err := db.Exec(cmd, id, user.Email, user.Password)

	if err != nil {
		return errors.New("Failed to create user. User already exists.")
	}

	fmt.Println("SUCCESS: new user crated")
	return nil
}

func GetUser(email string, db *sql.DB) (user User, e error) {
	cmd := `SELECT id, email, password FROM Users WHERE email = $1`

	row := db.QueryRow(cmd, email)

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return user, errors.New("Failed to get user.")
	}

	return user, nil
}

func GetEpisodes(id string, db *sql.DB) (episodes []Episode, e error) {
	cmd := `SELECT id, title, description, url, keywords, publishDate, author, user_id, episodeNumber FROM Episodes WHERE user_id = $1`

	rows, err := db.Query(cmd, id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Failed to get episodes.")
	}

	defer rows.Close()

	var episode Episode

	for rows.Next() {
		err := rows.Scan(&episode.ID, &episode.Title, &episode.Description, &episode.URL, &episode.UserID, &episode.Keywords, &episode.PublishDate, &episode.Author, &episode.EpisodeNumber)

		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("Failed to get episode.")

		}

		episodes = append(episodes, episode)

	}

	return episodes, nil

}
