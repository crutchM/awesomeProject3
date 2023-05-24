package repository

import (
	"awesomeProject3/internal/core/interfaces"
	"awesomeProject3/internal/repository/internal/playlist"
	"awesomeProject3/internal/repository/internal/video"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type repositoryManagerImpl struct {
	DB *pg.DB
}

func NewRepoManager(db *pg.DB) interfaces.RepositoryManager {
	return &repositoryManagerImpl{db}
}

func (rm *repositoryManagerImpl) getConnect() orm.DB {
	return rm.DB
}

func (rm *repositoryManagerImpl) GetVideoRepository() interfaces.VideoRepository {
	return video.NewVideoRepository(rm.DB)
}

func (rm *repositoryManagerImpl) GetPlaylistRepository() interfaces.PlayListRepository {
	return playlist.NewPlaylistRepository(rm.DB)
}
