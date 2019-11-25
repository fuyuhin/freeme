package services

import (
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
	GetAlbum(id int) (models.Album, error)
	GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error)
	Create(a *models.Album) error
	Update(id int, a models.Album) error
	Delete(id int) error
}

// AlbumService .
type AlbumService struct {
	Runtime freedom.Runtime
	// DefRepo *repositorys.AlbumRepository

	// albumRepoInterface AlbumRepoInterface
	// initiator.BindRepository AlbumRepository
	// [FTAL] The member variable must be publicly visible,
	// Its type is services.AlbumRepoInterface

	// AlbumRepoInterface AlbumRepoInterface

	AlbumRepoInterface
	// AlbumRepoInterface *AlbumRepoInterface
}

// GetAlbum .
func (s *AlbumService) GetAlbum(id int) (models.Album, error) {
	// return s.DefRepo.GetAlbum(id)
	// return (*s.AlbumRepoInterface).GetAlbum(id)
	return s.AlbumRepoInterface.GetAlbum(id)
}

func (s *AlbumService) GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error) {
	return s.AlbumRepoInterface.GetAlbums(page, perPage, filter)
	// return s.DefRepo.GetAlbums(page, perPage, filter)
}

// func (s *AlbumService) Create(a *models.Album) (*models.Album, error) {
// 	return s.DefRepo.Create(a)
// }

func (s *AlbumService) Update(id int, a models.Album) error {
	return s.AlbumRepoInterface.Update(id, a)
	// return s.DefRepo.Update(id, a)
}

func (s *AlbumService) Create(a *models.Album) error {
	return s.AlbumRepoInterface.Create(a)
	// return s.DefRepo.Create(a)
}

func (s *AlbumService) Delete(id int) error {
	// return s.DefRepo.Delete(id)
	return s.AlbumRepoInterface.Delete(id)
}

// GetAlbum .
func (s *AlbumService) Delete(id int) error {
	return s.DefRepo.Delete(id)
}

func (s *AlbumService) Update(id int, a models.Album) error {
	return s.DefRepo.Update(id, a)
}
