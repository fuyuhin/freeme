package repositorys

import (
	"freeme/models"

	"github.com/8treenet/freedom"
	"github.com/jinzhu/gorm"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindRepository(func() *AlbumRepository {
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
func (repo *AlbumRepository) GetAlbum(id int) (*models.Album, error) {
	album, err := models.FindAlbumByPrimary(repo, id)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAlbums .
func (repo *AlbumRepository) GetAlbums(page, perPage int, filter *models.Album) ([]models.Album, error) {
	builder := repo.NewDescOrder("AlbumId").NewPager(page, perPage)
	return models.FindAlbums(repo, filter, builder)
}

// Create .
// func (repo *AlbumRepository) Create(album *models.Album) (*models.Album, error) {
func (repo *AlbumRepository) Create(album *models.Album) error {
	affectedRows, err := models.CreateAlbum(repo, album)
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return models.ErrCreateAlbumRowAffected
	}
	return nil
}

func (repo *AlbumRepository) Update(id int, a models.Album) (err error) {
	a.AlbumID = id
	affectedRows, err := models.UpdateAlbum(repo, &a, a)
	if err != nil {
		return err
	}
	switch affectedRows {
	case 1:
		return nil
	case 0:
		return models.ErrUpdateAlbumNoAffected
	default:
		return models.ErrUpdateAlbumRowAffected
	}
}

func (repo *AlbumRepository) Delete(id int) error {
	affectedRows, err := models.DeleteAlbum(repo, &models.Album{AlbumID: id})
	if err != nil {
		return err
	}
	switch affectedRows {
	case 1:
		return nil
	case 0:
		return models.ErrDeleteAlbumNoAffected
	default:
		return models.ErrDeleteAlbumRowAffected
	}
}
