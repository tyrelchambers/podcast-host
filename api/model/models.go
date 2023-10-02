package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID      string         `gorm:"primaryKey" json:"uuid"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	Podcasts  []Podcast      `json:"podcasts"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Podcast struct {
	UUID              string         `gorm:"primaryKey" json:"uuid"`
	Title             string         `gorm:"not null" json:"title"`
	Description       string         `gorm:"not null" json:"description"`
	Thumbnail         string         `json:"thumbnail"`
	ExplicitContent   bool           `json:"explicit_content"`
	PrimaryCategory   string         `json:"primary_category"`
	SecondaryCategory string         `json:"secondary_category"`
	Author            string         `gorm:"not null" json:"author"`
	Copyright         string         `json:"copyright"`
	Keywords          string         `json:"keywords"`
	Website           string         `json:"website"`
	Language          string         `json:"language"`
	Timezone          string         `json:"timezone"`
	ShowOwner         string         `json:"show_owner"`
	OwnerEmail        string         `gorm:"not null" json:"owner_email"`
	DisplayEmailInRSS bool           `json:"display_email_in_rss"`
	UserID            string         `json:"user_id"`
	Episodes          []Episode      `json:"episodes"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

type Episode struct {
	UUID          string         `gorm:"primaryKey" json:"uuid"`
	Title         string         `gorm:"not null" json:"title"`
	Description   string         `json:"description"`
	URL           string         `json:"url"`
	Image         string         `json:"image"`
	PodcastId     string         `json:"podcast_id"`
	Keywords      string         `json:"keywords"`
	PublishDate   uint64         `json:"publish_date"`
	Author        string         `json:"author"`
	EpisodeNumber uint64         `json:"episode_number"`
	Draft         bool           `json:"draft"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type Cookie struct {
	UserID       string
	UUID         string `gorm:"primaryKey"`
	SessionToken string
	ExpiresAt    time.Time
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
