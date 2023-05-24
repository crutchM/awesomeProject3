package entity

type Playlist struct {
	Id          string
	CustomerId  string
	Name        string
	Description string
	Videos      []VideoResponse
}

type PlaylistDTO struct {
	Id          string
	CustomerId  string
	Name        string
	Description string
	Videos      []VideoDTO
}
