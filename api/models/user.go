package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserExists(email string, db *sql.DB) (userExists bool) {
	cmd := `SELECT id, email, password FROM Users WHERE email = $1`

	row := db.QueryRow(cmd, email)

	var user User

	row.Scan(&user.ID, &user.Email, &user.Password)

	if user.ID == "" {
		return false
	}

	return true
}

func CreateUser(user *User, db *sql.DB) (u User, e error) {
	cmd := `INSERT INTO Users (id, email, password) VALUES ($1, $2, $3) RETURNING id, email`

	id := cuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
		return u, errors.New("Failed to hash password.")
	}

	var uR *sql.Rows

	uR, err = db.Query(cmd, id, user.Email, hashPassword)

	if err != nil {
		fmt.Println(err.Error())
		return u, errors.New("Failed to create user. User already exists.")
	}

	uR.Scan(&u.ID, &u.Email)

	fmt.Println("SUCCESS: new user crated")
	return u, nil
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
	cmd := `SELECT id, title, url, publishDate, episodeNumber FROM Episodes WHERE user_id = $1`

	rows, err := db.Query(cmd, id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Failed to get episodes.")
	}

	defer rows.Close()

	var episode Episode

	for rows.Next() {
		err := rows.Scan(&episode.ID, &episode.Title, &episode.URL, &episode.PublishDate, &episode.EpisodeNumber)

		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("Failed to get episode.")

		}

		episodes = append(episodes, episode)

	}

	return episodes, nil

}
