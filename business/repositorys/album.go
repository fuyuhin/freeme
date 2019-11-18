package repositorys

import (
	"freeme/models"

	"github.com/8treenet/freedom"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindRepository(func() *AlbumRepository {
			println("initiator.BindRepository AlbumRepository")
			return &AlbumRepository{}
		})
	})
}

// AlbumRepository .
type AlbumRepository struct {
	freedom.Repository
	models.Album
}

// GetAlbum .
func (repo *AlbumRepository) GetAlbum(id int) (models.Album, error) {
	return models.FindAlbumByPrimary(repo, id)
}
