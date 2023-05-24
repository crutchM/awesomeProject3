package playlist

import (
	"awesomeProject3/internal/core/entity"
	"awesomeProject3/internal/core/interfaces"
	"github.com/go-pg/pg/orm"
)

type playlistRepository struct {
	db orm.DB
}

func NewPlaylistRepository(db orm.DB) interfaces.PlayListRepository {
	return &playlistRepository{db: db}
}

func (p playlistRepository) CreatePlaylist(name string) (entity.PlaylistDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (p playlistRepository) AddToPlaylist(videoUUID, playlistUUID string) error {
	//TODO implement me
	panic("implement me")
}

func (p playlistRepository) GetPlaylist(uuid string) (entity.PlaylistDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (p playlistRepository) GetPlaylistsList(customerID string) ([]entity.PlaylistDTO, error) {
	//TODO implement me
	panic("implement me")
}
