package services

import (
	"freeme/business/repositorys"
	"freeme/models"

	"github.com/8treenet/freedom"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindService(func() *AlbumService {
			return &AlbumService{}
		})
	})
}

// AlbumRepoInterface .
type AlbumRepoInterface interface {
	GetUA() string
}

// AlbumService .
type AlbumService struct {
	Runtime   freedom.Runtime
	DefRepo   *repositorys.AlbumRepository
	DefRepoIF AlbumRepoInterface
}

// GetAlbum .
func (s *AlbumService) GetAlbum(id int) (models.Album, error) {
	return s.DefRepo.GetAlbum(id)
}

func (s *AlbumService) GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error) {
	return s.DefRepo.GetAlbums(page, perPage, filter)
}

// func (s *AlbumService) Create(a *models.Album) (*models.Album, error) {
// 	return s.DefRepo.Create(a)
// }

func (s *AlbumService) Create(a *models.Album) error {
	return s.DefRepo.Create(a)
}
