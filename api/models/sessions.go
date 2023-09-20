package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lucsky/cuid"
)

func CreateSession(user *User, db *sql.DB) (s Session, e error) {
	cmd := `INSERT INTO Sessions (id, user_id, session_token, expires_at) VALUES ($1, $2, $3, $4) RETURNING *`
	stmt, err := db.Prepare(cmd)

	if err != nil {
		fmt.Println(err.Error())
		return s, errors.New("Failed to create session.")
	}

	var session Session

	session_token := cuid.New()

	id := cuid.New()
	expires_at := time.Now().Add(24 * time.Hour)

	err = stmt.QueryRow(id, user.ID, session_token, expires_at).Scan(&session.ID, &session.UserID, &session.SessionToken, &session.ExpiresAt)

	if err != nil {
		println(err.Error())
		return session, errors.New("Failed to create session.")
	}

	return session, nil
}

func GetSession(sessionToken string, db *sql.DB) (u User, e error) {
	var session Session
	cmd := `SELECT user_id FROM Sessions WHERE session_token = $1`

	rows, err := db.Query(cmd, sessionToken)

	if err != nil {
		println(err.Error())
		return u, errors.New("Failed to get session.")
	}

	defer rows.Close()

	rows.Scan(&session.UserID)

	cmd = `SELECT id, email FROM Users WHERE id = $1`

	rows, err = db.Query(cmd, session.UserID)

	if err != nil {
		println(err.Error())
		return u, errors.New("Failed to get user from session.")
	}

	defer rows.Close()

	rows.Scan(&u.ID, &u.Email)

	return u, nil
}
