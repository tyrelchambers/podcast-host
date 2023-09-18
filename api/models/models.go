package models

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
	UserID        string `json:"user_id"`
	Keywords      string `json:"keywords"`
	PublishDate   string `json:"publishDate"`
	Author        string `json:"author"`
	EpisodeNumber string `json:"episodeNumber"`
}
