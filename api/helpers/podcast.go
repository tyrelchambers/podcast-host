package helpers

import (
	"api/model"
	"database/sql"
	"encoding/json"
	"fmt"
)

func ParsePodcasts(rows *sql.Rows) []model.Podcast {
	var podcasts []model.Podcast

	for rows.Next() {
		var p model.Podcast
		var episodeJSON []byte

		err := rows.Scan(
			&p.UUID,
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
			return nil
		}

		// Unmarshal the episodeJSON into the Episodes field of the Podcast struct
		if err := json.Unmarshal(episodeJSON, &p.Episodes); err != nil {
			fmt.Println(err.Error())
			return nil
		}

		// Append the podcast to the podcasts slice
		podcasts = append(podcasts, p)
	}

	return podcasts
}
