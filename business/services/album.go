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
	GetAlbum(id int) (*models.Album, error)
	GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error)
	Create(a *models.Album) error
	Update(id int, a models.Album) error
	Delete(id int) error
}

// AlbumService .
type AlbumService struct {
	Runtime freedom.Runtime
	AlbumRepoInterface
}

// GetAlbum .
func (s *AlbumService) GetAlbum(id int) (*models.Album, error) {
	return s.AlbumRepoInterface.GetAlbum(id)
}

func (s *AlbumService) GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error) {
	return s.AlbumRepoInterface.GetAlbums(page, perPage, filter)
}

func (s *AlbumService) Create(a *models.Album) error {
	return s.AlbumRepoInterface.Create(a)
}

func (s *AlbumService) Update(id int, a models.Album) error {
	return s.AlbumRepoInterface.Update(id, a)
}

func (s *AlbumService) Delete(id int) error {
	return s.AlbumRepoInterface.Delete(id)
}
