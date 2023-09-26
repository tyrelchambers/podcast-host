package models

import (
	"api/helpers"
	"api/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lucsky/cuid"
)

func CreatePodcast(p *model.Podcast, db *sql.DB) error {
	var podcast model.Podcast

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

func GetUsersPodcasts(userId string, db *sql.DB) ([]model.Podcast, error) {

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

	parsedPodcasts := helpers.ParsePodcasts(rows)

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return parsedPodcasts, nil
}

func GetPodcastById(id string, userId string, db *sql.DB) (p model.Podcast, e error) {
	var podcast model.Podcast
	var episodeJSON []byte

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
				Podcasts.ID = $1 AND Podcasts.user_id = $2
		GROUP BY
				Podcasts.ID;
	`
	rows := db.QueryRow(cmd, id, userId)

	err := rows.Scan(
		&podcast.ID,
		&podcast.Title,
		&podcast.Description,
		&podcast.Thumbnail,
		&podcast.ExplicitContent,
		&podcast.PrimaryCategory,
		&podcast.SecondaryCategory,
		&podcast.Author,
		&podcast.Copyright,
		&podcast.Keywords,
		&podcast.Website,
		&podcast.Language,
		&podcast.Timezone,
		&podcast.ShowOwner,
		&podcast.OwnerEmail,
		&podcast.DisplayEmailInRSS,
		&podcast.UserID,
		&episodeJSON,
	)

	if err != nil {
		fmt.Println(err.Error())
		return podcast, err
	}

	if err := json.Unmarshal(episodeJSON, &podcast.Episodes); err != nil {
		return podcast, err
	}

	return podcast, nil
}

func GetPodcastByNameWithEpisodes(name string, userId string, db *sql.DB) (p model.Podcast, e error) {
	var podcast model.Podcast
	var episodeJSON []byte

	parsedName := strings.Replace(name, "-", " ", -1)

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
				Podcasts.user_id = $1 AND Podcasts.title = $2
		GROUP BY
				Podcasts.ID;
	`

	rows := db.QueryRow(cmd, userId, parsedName)

	err := rows.Scan(
		&podcast.ID,
		&podcast.Title,
		&podcast.Description,
		&podcast.Thumbnail,
		&podcast.ExplicitContent,
		&podcast.PrimaryCategory,
		&podcast.SecondaryCategory,
		&podcast.Author,
		&podcast.Copyright,
		&podcast.Keywords,
		&podcast.Website,
		&podcast.Language,
		&podcast.Timezone,
		&podcast.ShowOwner,
		&podcast.OwnerEmail,
		&podcast.DisplayEmailInRSS,
		&podcast.UserID,
		&episodeJSON,
	)

	if err != nil {
		fmt.Println(err.Error())
		return podcast, err
	}

	if err := json.Unmarshal(episodeJSON, &podcast.Episodes); err != nil {
		return podcast, err
	}

	return podcast, nil

}

func GetPodcastIdFromName(name string, db *sql.DB) (p model.Podcast, e error) {
	var podcast model.Podcast

	parsedName := strings.Replace(name, "-", " ", -1)

	cmd := `SELECT id FROM Podcasts WHERE title = $1`

	rows := db.QueryRow(cmd, parsedName)

	err := rows.Scan(
		&podcast.ID,
	)

	if err != nil {
		fmt.Println(err.Error())
		return podcast, err
	}

	return podcast, nil
}

func GetPodcastEpisodesById(id string, db *sql.DB) ([]model.Episode, error) {
	var episodes []model.Episode

	cmd := `SELECT id, title, description, url, keywords, publish_date, author, episode_number, draft FROM Episodes WHERE podcast_id = $1`

	rows, err := db.Query(cmd, id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var episode model.Episode

		err := rows.Scan(
			&episode.ID,
			&episode.Title,
			&episode.Description,
			&episode.URL,
			&episode.Keywords,
			&episode.PublishDate,
			&episode.Author,
			&episode.EpisodeNumber,
			&episode.Draft,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		episodes = append(episodes, episode)

	}

	return episodes, nil

}
