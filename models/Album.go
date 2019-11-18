package models

import "github.com/8treenet/freedom"

// Album .
//type Album struct {
//	AlbumId  int `gorm:"primary_key"`
//	Title    string
//	ArtistId int
//}
//Album test
type Album struct {
	AlbumID  int    `gorm:"column:AlbumId;primary_key" json:"id"`
	Title    string `gorm:"column:Title" json:"title"`
	ArtistID int    `gorm:"column:ArtistId" json:"artist_id"`
}

//TableName
func (m *Album) TableName() string {
	//return "Album"
	return "Albums"
}

// FindAlbumByPrimary .
func FindAlbumByPrimary(rep freedom.GORMRepository, primary interface{}) (result Album, e error) {
	e = rep.DB().Find(&result, primary).Error
	return
}

// FindAlbumsByPrimarys .
func FindAlbumsByPrimarys(rep freedom.GORMRepository, primarys ...interface{}) (results []Album, e error) {
	e = rep.DB().Find(&results, primarys).Error
	return
}

// FindAlbum .
func FindAlbum(rep freedom.GORMRepository, query *Album, builders ...freedom.QueryBuilder) (result Album, e error) {
	if len(builders) == 0 {
		e = rep.DB().Where(query).Last(&result).Error
		return
	}

	e = rep.DB().Where(query).Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// FindAlbumByWhere .
func FindAlbumByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (result Album, e error) {
	if len(builders) == 0 {
		e = rep.DB().Where(query, args...).Last(&result).Error
		return
	}

	e = rep.DB().Where(query, args...).Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// FindAlbums .
func FindAlbums(rep freedom.GORMRepository, query *Album, builders ...freedom.QueryBuilder) (results []Album, e error) {
	db := rep.DB()
	if len(builders) == 0 {
		e = db.Where(query).Find(&results).Error
		return
	}

	where := db.Where(query)
	e = builders[0].Execute(where, &results)
	return
}

// FindAlbumsByWhere .
func FindAlbumsByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (results []Album, e error) {
	db := rep.DB()
	if len(builders) == 0 {
		e = db.Where(query, args...).Find(&results).Error
		return
	}

	where := db.Where(query, args...)
	e = builders[0].Execute(where, &results)
	return
}

// CreateAlbum .
func CreateAlbum(rep freedom.GORMRepository, entity *Album) (*Album, error) {
	db := rep.DB().Create(entity)
	// rowsAffected = db.RowsAffected
	if db.Error != nil {
		return nil, db.Error
	}
	return entity, nil
}

// UpdateAlbum .
func UpdateAlbum(rep freedom.GORMRepository, entity *Album, value Album) (affected int64, e error) {
	db := rep.DB().Model(entity).Updates(value)
	e = db.Error
	affected = db.RowsAffected
	return
}

// FindToUpdateAlbums .
func FindToUpdateAlbums(rep freedom.GORMRepository, query *Album, value Album, builders ...freedom.QueryBuilder) (affected int64, e error) {
	db := rep.DB()
	if len(builders) > 0 {
		db = db.Model(&Album{}).Where(query).Order(builders[0].Order()).Limit(builders[0].Limit()).Updates(value)
	} else {
		db = db.Model(&Album{}).Where(query).Updates(value)
	}

	affected = db.RowsAffected
	e = db.Error
	return
}

// FindByWhereToUpdateAlbums .
func FindByWhereToUpdateAlbums(rep freedom.GORMRepository, query string, args []interface{}, value Album, builders ...freedom.QueryBuilder) (affected int64, e error) {
	db := rep.DB()
	if len(builders) > 0 {
		db = db.Model(&Album{}).Where(query, args...).Order(builders[0].Order()).Limit(builders[0].Limit()).Updates(value)
	} else {
		db = db.Model(&Album{}).Where(query, args...).Updates(value)
	}

	affected = db.RowsAffected
	e = db.Error
	return
}
