package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string     `json:"uuid"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Podcasts []*Podcast `json:"podcasts" gorm:"foreignKey:UserID"`
}
type Podcast struct {
	gorm.Model
	UUID              string     `json:"uuid"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Thumbnail         string     `json:"thumbnail"`
	ExplicitContent   bool       `json:"explicit_content"`
	PrimaryCategory   string     `json:"primary_category"`
	SecondaryCategory string     `json:"secondary_category"`
	Author            string     `json:"author"`
	Copyright         string     `json:"copyright"`
	Keywords          string     `json:"keywords"`
	Website           string     `json:"website"`
	Language          string     `json:"language"`
	Timezone          string     `json:"timezone"`
	ShowOwner         string     `json:"show_owner"`
	OwnerEmail        string     `json:"owner_email"`
	DisplayEmailInRSS bool       `json:"display_email_in_rss"`
	UserID            string     `json:"user_id"`
	Episodes          []*Episode `gorm:"foreignKey:PodcastId" json:"episodes"`
}

type Episode struct {
	UUID          string   `json:"uuid"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	URL           string   `json:"url"`
	Image         string   `json:"image"`
	PodcastId     string   `json:"podcast_id"`
	Podcast       *Podcast `gorm:"foreignKey:UUID;references:PodcastId"`
	Keywords      string   `json:"keywords"`
	PublishDate   uint64   `json:"publish_date"`
	Author        string   `json:"author"`
	EpisodeNumber uint64   `json:"episode_number"`
	Draft         bool     `json:"draft"`
}

type Cookie struct {
	UserID       string
	UUID         string `gorm:"primaryKey"`
	SessionToken string
}

type PodcastEpisode struct {
	UUID      string `gorm:"primaryKey"`
	Episodes  []Episode
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
