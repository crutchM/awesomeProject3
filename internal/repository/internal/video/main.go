package video

import (
	"awesomeProject3/internal/core/entity"
	"awesomeProject3/internal/core/interfaces"
	"github.com/go-pg/pg/orm"
)

type videoRepositoryImpl struct {
	db orm.DB
}

func NewVideoRepository(db orm.DB) interfaces.VideoRepository {
	return &videoRepositoryImpl{db: db}
}

func (v *videoRepositoryImpl) GetVideo(uuid string) entity.VideoDTO {
	//TODO implement me
	panic("implement me")
}

func (v *videoRepositoryImpl) CreateVideo(video entity.VideoDTO) (entity.VideoDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoRepositoryImpl) GetVideoList(query ...string) ([]entity.VideoDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoRepositoryImpl) GetVideoMinimized(uuid string) (entity.VideoMinimized, error) {
	//TODO implement me
	panic("implement me")
}
