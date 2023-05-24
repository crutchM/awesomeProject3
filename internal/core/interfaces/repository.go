package interfaces

import "awesomeProject3/internal/core/entity"

type RepositoryManager interface {
	GetVideoRepository() VideoRepository
	GetPlaylistRepository() PlayListRepository
}

type VideoRepository interface {
	GetVideo(uuid string) entity.VideoDTO
	CreateVideo(video entity.VideoDTO) (entity.VideoDTO, error)
	GetVideoList(query ...string) ([]entity.VideoDTO, error)
	GetVideoMinimized(uuid string) (entity.VideoMinimized, error)
}

type PlayListRepository interface {
	CreatePlaylist(name string) (entity.PlaylistDTO, error)
	AddToPlaylist(videoUUID, playlistUUID string) error
	GetPlaylist(uuid string) (entity.PlaylistDTO, error)
	GetPlaylistsList(customerID string) ([]entity.PlaylistDTO, error)
}
