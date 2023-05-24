package interfaces

import "awesomeProject3/internal/core/entity"

type VideoUseCase interface {
	GetVideo(uuid string) (entity.VideoResponse, error)
	CreateVideo(video entity.PostVideoRequest) (entity.VideoResponse, error)
	GetVideoList(query ...string) ([]entity.VideoMinimized, error)
	GetVideoMinimized(uuid string) (entity.VideoMinimized, error)
}

type PlayListUseCase interface {
	CreatePlaylist(name string) (entity.Playlist, error)
	AddToPlaylist(videoUUID, playlistUUID string) error
	GetPlaylist(uuid string) (entity.Playlist, error)
	GetPlaylistsList(customerID string) ([]entity.Playlist, error)
}
