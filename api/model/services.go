package model

func (p *PodcastDTO) ToEntity() *Podcast {
	pDto := Podcast{
		UUID:              p.UUID,
		Title:             p.Title,
		Description:       p.Description,
		Thumbnail:         p.Thumbnail,
		ExplicitContent:   p.ExplicitContent,
		PrimaryCategory:   p.PrimaryCategory,
		SecondaryCategory: p.SecondaryCategory,
		Author:            p.Author,
		Copyright:         p.Copyright,
		Keywords:          p.Keywords,
		Website:           p.Website,
		Language:          p.Language,
		Timezone:          p.Timezone,
		ShowOwner:         p.ShowOwner,
		OwnerEmail:        p.OwnerEmail,
		DisplayEmailInRSS: p.DisplayEmailInRSS,
		UserID:            p.UserID,
	}

	return &pDto
}

func (p *Podcast) ToDTO() *PodcastDTO {
	pDto := PodcastDTO{
		UUID:              p.UUID,
		Title:             p.Title,
		Description:       p.Description,
		Thumbnail:         p.Thumbnail,
		ExplicitContent:   p.ExplicitContent,
		PrimaryCategory:   p.PrimaryCategory,
		SecondaryCategory: p.SecondaryCategory,
		Author:            p.Author,
		Copyright:         p.Copyright,
		Keywords:          p.Keywords,
		Website:           p.Website,
		Language:          p.Language,
		Timezone:          p.Timezone,
		ShowOwner:         p.ShowOwner,
		OwnerEmail:        p.OwnerEmail,
		DisplayEmailInRSS: p.DisplayEmailInRSS,
		UserID:            p.UserID,
	}

	return &pDto
}

func (p *PodcastDTO) GetPublishedEpisodes() []*EpisodeDTO {
	var episodes []*EpisodeDTO

	for _, episode := range p.Episodes {
		if !episode.Draft {
			episodes = append(episodes, episode)
		}
	}

	return episodes
}

func (p *PodcastDTO) GetDrafts() []*EpisodeDTO {
	var episodes []*EpisodeDTO

	for _, episode := range p.Episodes {
		if episode.Draft {
			episodes = append(episodes, episode)
		}
	}

	return episodes
}

func (e *EpisodeDTO) ToEntity() *Episode {
	eDto := Episode{
		UUID:          e.UUID,
		Title:         e.Title,
		Description:   e.Description,
		URL:           e.URL,
		Image:         e.Image,
		PodcastId:     e.PodcastId,
		Keywords:      e.Keywords,
		PublishDate:   e.PublishDate,
		Author:        e.Author,
		EpisodeNumber: e.EpisodeNumber,
		Draft:         e.Draft,
		IsScheduled:   e.IsScheduled,
	}

	return &eDto
}

func (e *Episode) ToDTO() *EpisodeDTO {
	eDto := EpisodeDTO{
		UUID:          e.UUID,
		Title:         e.Title,
		Description:   e.Description,
		URL:           e.URL,
		Image:         e.Image,
		PodcastId:     e.PodcastId,
		Keywords:      e.Keywords,
		PublishDate:   e.PublishDate,
		Author:        e.Author,
		EpisodeNumber: e.EpisodeNumber,
		Draft:         e.Draft,
		IsScheduled:   e.IsScheduled,
	}

	return &eDto
}

func (u *User) ToDTO() *UserDTO {
	uDto := UserDTO{
		UUID:     u.UUID,
		Email:    u.Email,
		Podcasts: make([]*PodcastDTO, len(u.Podcasts)),
	}

	for i, podcast := range u.Podcasts {
		uDto.Podcasts[i] = &PodcastDTO{
			UUID:              podcast.UUID,
			Title:             podcast.Title,
			Description:       podcast.Description,
			Thumbnail:         podcast.Thumbnail,
			ExplicitContent:   podcast.ExplicitContent,
			PrimaryCategory:   podcast.PrimaryCategory,
			SecondaryCategory: podcast.SecondaryCategory,
			Author:            podcast.Author,
			Copyright:         podcast.Copyright,
			Keywords:          podcast.Keywords,
			Website:           podcast.Website,
			Language:          podcast.Language,
			Timezone:          podcast.Timezone,
			ShowOwner:         podcast.ShowOwner,
			OwnerEmail:        podcast.OwnerEmail,
			DisplayEmailInRSS: podcast.DisplayEmailInRSS,
			UserID:            podcast.UserID,
		}
	}

	return &uDto
}
