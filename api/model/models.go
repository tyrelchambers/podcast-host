package model

import (
	"time"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Episode struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	URL           string `json:"url"`
	Image         string `json:"image"`
	PodcastId     string `json:"podcast_id"`
	Keywords      string `json:"keywords"`
	PublishDate   string `json:"publishDate"`
	Author        string `json:"author"`
	EpisodeNumber string `json:"episodeNumber"`
	Draft         bool   `json:"draft"`
}

type Podcast struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Thumbnail         string    `json:"thumbnail"`
	ExplicitContent   bool      `json:"explicit_content"`
	PrimaryCategory   string    `json:"primary_category"`
	SecondaryCategory string    `json:"secondary_category"`
	Author            string    `json:"author"`
	Copyright         string    `json:"copyright"`
	Keywords          string    `json:"keywords"`
	Website           string    `json:"website"`
	Language          string    `json:"language"`
	Timezone          string    `json:"timezone"`
	ShowOwner         string    `json:"show_owner"`
	OwnerEmail        string    `json:"owner_email"`
	DisplayEmailInRSS bool      `json:"display_email_in_rss"`
	UserID            string    `json:"user_id"`
	Episodes          []Episode `json:"episodes"`
}

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	SessionToken string    `json:"session_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Cookie struct {
	UserID       string
	SessionToken string
	ExpiresAt    time.Time
	Email        string
}

type PodcastEpisode struct {
	Podcast
	Episodes []Episode
}
