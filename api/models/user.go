package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserExists(email string, db *sql.DB) (userExists bool) {
	if email == "" {
		panic(errors.New("Email cannot be empty."))
	}
	cmd := `SELECT id FROM Users WHERE email = $1`

	row := db.QueryRow(cmd, email)

	var user User

	_ = row.Scan(&user.ID)

	if user.ID == "" {
		return false
	}

	return true
}

func CreateUser(user *User, db *sql.DB) (u *User, e error) {
	stmt, _ := db.Prepare(`INSERT INTO Users (id, email, password) VALUES ($1, $2, $3) RETURNING id, email`)

	id := cuid.New()

	var newUser User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
		return u, errors.New("Failed to hash password.")
	}

	stmt.QueryRow(id, user.Email, hashPassword).Scan(&newUser.ID, &newUser.Email)

	fmt.Println("SUCCESS: new user created")
	return &newUser, nil
}

func GetUser(id *string, db *sql.DB) (user User, e error) {
	var u User
	cmd := `SELECT id, email FROM Users WHERE id = $1`

	row := db.QueryRow(cmd, *id)

	err := row.Scan(&u.ID, &u.Email)

	if err != nil {
		fmt.Println(err.Error())
		return u, errors.New("Failed to get user.")
	}

	return u, nil
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
