package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lucsky/cuid"
)

func CreateSession(user *User, db *sql.DB) (s Session, e error) {
	cmd := `INSERT INTO Sessions (id, email, user_id, session_token, expires_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	stmt, _ := db.Prepare(cmd)

	var session Session

	session_token := cuid.New()

	id := cuid.New()
	expires_at := time.Now().Add(24 * time.Hour)

	err := stmt.QueryRow(id, user.Email, user.ID, session_token, expires_at).Scan(&session.ID, &session.Email, &session.UserID, &session.SessionToken, &session.ExpiresAt)

	if err != nil {
		println(err.Error())
		return session, errors.New("Failed to create session.")
	}

	return session, nil
}
