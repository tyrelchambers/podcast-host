package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID      string     `gorm:"primaryKey"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Podcasts  []*Podcast `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Podcast struct {
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
	User              *User      `gorm:"foreignKey:UUID;references:UserID"`
	Episodes          []*Episode `gorm:"foreignKey:PodcastId"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Episode struct {
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

type Cookie struct {
	UserID       string
	UUID         string `gorm:"primaryKey"`
	SessionToken string
}

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
