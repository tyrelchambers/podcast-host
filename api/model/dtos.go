package model

import (
	"time"

	"gorm.io/gorm"
)

type UserDTO struct {
	UUID      string     `gorm:"primaryKey"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Podcasts  []*Podcast `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type PodcastDTO struct {
	UUID              string `gorm:"primaryKey"`
	Title             string `gorm:"not null"`
	Description       string `gorm:"not null"`
	Thumbnail         string
	ExplicitContent   bool
	PrimaryCategory   string
	SecondaryCategory string
	Author            string `gorm:"not null" `
	Copyright         string
	Keywords          string
	Website           string
	Language          string
	Timezone          string
	ShowOwner         string
	OwnerEmail        string `gorm:"not null" `
	DisplayEmailInRSS bool
	UserID            string
	User              User
	Episodes          []*Episode `gorm:"foreignKey:PodcastId"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type EpisodeDTO struct {
	UUID          string `gorm:"primaryKey"`
	Title         string `gorm:"not null"`
	Description   string
	URL           string
	Image         string
	PodcastId     string
	Podcast       *Podcast `gorm:"foreignKey:UUID;references:PodcastId"`
	Keywords      string
	PublishDate   uint64
	Author        string
	EpisodeNumber uint64
	Draft         bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" `
}

type CookieDTO struct {
	UserID       string
	UUID         string `gorm:"primaryKey"`
	SessionToken string
	ExpiresAt    time.Time
}

type PodcastEpisodeDTO struct {
	UUID      string `gorm:"primaryKey"`
	Episodes  []Episode
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterBodyDTO struct {
	Email    string
	Password string
}
