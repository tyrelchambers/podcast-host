package models

import (
	"api/model"
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

	var user model.User

	_ = row.Scan(&user.ID)

	if user.ID == "" {
		return false
	}

	return true
}

func CreateUser(user *model.User, db *sql.DB) (u *model.User, e error) {
	stmt, _ := db.Prepare(`INSERT INTO Users (id, email, password) VALUES ($1, $2, $3) RETURNING id, email`)

	id := cuid.New()

	var newUser model.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
		return u, errors.New("Failed to hash password.")
	}

	stmt.QueryRow(id, user.Email, hashPassword).Scan(&newUser.ID, &newUser.Email)

	fmt.Println("SUCCESS: new user created")
	return &newUser, nil
}

func GetUser(id *string, db *sql.DB) (user model.User, e error) {
	var u model.User
	cmd := `SELECT id, email FROM Users WHERE id = $1`

	row := db.QueryRow(cmd, *id)

	err := row.Scan(&u.ID, &u.Email)

	if err != nil {
		return u, errors.New("Failed to get user. Doesn't exist.")
	}

	return u, nil
}
