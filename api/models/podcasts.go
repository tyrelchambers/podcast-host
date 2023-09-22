package models

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lucsky/cuid"
)

func CreatePodcast(p *Podcast, db *sql.DB) error {
	var podcast Podcast

	cmd := `INSERT INTO Podcasts (
		id,
		title,
		description,
		thumbnail,
		explicit_content,
		primary_category,
		secondary_category,
		author,
		copyright,
		keywords,
		website,
		language,
		timezone,
		show_owner,
		owner_email,
		display_email_in_rss,
		user_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING *`

	stmt, err := db.Prepare(cmd)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	id := cuid.New()

	err = stmt.QueryRow(id, p.Title, p.Description, p.Thumbnail, p.ExplicitContent, p.PrimaryCategory, p.SecondaryCategory, p.Author, p.Copyright, p.Keywords, p.Website, p.Language, p.Timezone, p.ShowOwner, p.OwnerEmail, p.DisplayEmailInRSS, p.UserID).Scan(&podcast.ID, &podcast.Title, &podcast.Description, &podcast.Thumbnail, &podcast.ExplicitContent, &podcast.PrimaryCategory, &podcast.SecondaryCategory, &podcast.Author, &podcast.Copyright, &podcast.Keywords, &podcast.Website, &podcast.Language, &podcast.Timezone, &podcast.ShowOwner, &podcast.OwnerEmail, &podcast.DisplayEmailInRSS, &podcast.UserID)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("SUCCESS: new podcast created")

	return nil
}

func GetUsersPodcasts(userId string, db *sql.DB) ([]Podcast, error) {

	cmd := `
		SELECT
			Podcasts.*,
			(
				SELECT COALESCE(json_agg(Episodes.*), '[]'::json)
				FROM Episodes
				WHERE Podcasts.ID = Episodes.podcast_id
			) AS episodes
		FROM
				Podcasts
		LEFT JOIN
				Episodes
		ON
				Podcasts.ID = Episodes.podcast_id
		WHERE
				Podcasts.user_id = $1
		GROUP BY
				Podcasts.ID;
	`

	rows, err := db.Query(cmd, userId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var podcasts []Podcast

	for rows.Next() {
		var p Podcast
		var episodeJSON []byte

		// var es Episode

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Thumbnail,
			&p.ExplicitContent,
			&p.PrimaryCategory,
			&p.SecondaryCategory,
			&p.Author,
			&p.Copyright,
			&p.Keywords,
			&p.Website,
			&p.Language,
			&p.Timezone,
			&p.ShowOwner,
			&p.OwnerEmail,
			&p.DisplayEmailInRSS,
			&p.UserID,
			&episodeJSON,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		// Unmarshal the episodeJSON into the Episodes field of the Podcast struct
		if err := json.Unmarshal(episodeJSON, &p.Episodes); err != nil {
			return nil, err
		}

		// Append the podcast to the podcasts slice
		podcasts = append(podcasts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return podcasts, nil
}
