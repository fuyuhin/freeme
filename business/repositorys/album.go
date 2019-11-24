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

// GetAlbums .
func (repo *AlbumRepository) GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error) {
	builder := repo.NewDescOrder("AlbumId").SetPager(page, perPage)
	return models.FindAlbums(repo, filter, builder)
}

// Create .
// func (repo *AlbumRepository) Create(album *models.Album) (*models.Album, error) {
func (repo *AlbumRepository) Create(album *models.Album) error {
	// affectedRows, err := models.CreateAlbum(repo, album)
	// if err != nil {
	// 	return nil, err
	// }
	// if affectedRows != 0 {
	// 	return nil, models.Err_CreateEntityRowsAffected
	// }
	// return album, nil
	affectedRows, err := models.CreateAlbum(repo, album)
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return models.ErrRowsAffected
	}
	return nil
}

func (repo *AlbumRepository) Update(id int, album models.Album) error {
	affectedRows, err := models.UpdateAlbum(repo, &models.Album{AlbumID: id}, album)
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return models.ErrRowsAffected
	}
	return nil
}

func (repo *AlbumRepository) Delete(id int) error {
	affectedRows, err := models.DeleteAlbum(repo, &models.Album{AlbumID: id})
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return models.ErrRowsAffected
	}
	return nil
}
