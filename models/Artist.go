package models

import "github.com/8treenet/freedom"

// Artist .
type Artist struct {
	ArtistID int `gorm:"primary_key"`
	Name     string
}

func (m *Artist) TableName() string {
	return "Artist"
}

// FindArtistByPrimary .
func FindArtistByPrimary(rep freedom.GORMRepository, primary interface{}) (result Artist, e error) {
	e = rep.DB().Find(&result, primary).Error
	return
}

// FindArtistsByPrimarys .
func FindArtistsByPrimarys(rep freedom.GORMRepository, primarys ...interface{}) (results []Artist, e error) {
	e = rep.DB().Find(&results, primarys).Error
	return
}

// FindArtist .
func FindArtist(rep freedom.GORMRepository, query *Artist, builders ...freedom.QueryBuilder) (result Artist, e error) {
	if len(builders) == 0 {
		e = rep.DB().Where(query).Last(&result).Error
		return
	}

	e = rep.DB().Where(query).Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// FindArtistByWhere .
func FindArtistByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (result Artist, e error) {
	if len(builders) == 0 {
		e = rep.DB().Where(query, args...).Last(&result).Error
		return
	}

	e = rep.DB().Where(query, args...).Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// FindArtists .
func FindArtists(rep freedom.GORMRepository, query *Artist, builders ...freedom.QueryBuilder) (results []Artist, e error) {
	db := rep.DB()
	if len(builders) == 0 {
		e = db.Where(query).Find(&results).Error
		return
	}

	where := db.Where(query)
	e = builders[0].Execute(where, &results)
	return
}

// FindArtistsByWhere .
func FindArtistsByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (results []Artist, e error) {
	db := rep.DB()
	if len(builders) == 0 {
		e = db.Where(query, args...).Find(&results).Error
		return
	}

	where := db.Where(query, args...)
	e = builders[0].Execute(where, &results)
	return
}

// CreateArtist .
func CreateArtist(rep freedom.GORMRepository, entity *Artist) (rowsAffected int64, e error) {
	db := rep.DB().Create(entity)
	rowsAffected = db.RowsAffected
	e = db.Error
	return
}

// UpdateArtist .
func UpdateArtist(rep freedom.GORMRepository, entity *Artist, value Artist) (affected int64, e error) {
	db := rep.DB().Model(entity).Updates(value)
	e = db.Error
	affected = db.RowsAffected
	return
}

// FindToUpdateArtists .
func FindToUpdateArtists(rep freedom.GORMRepository, query *Artist, value Artist, builders ...freedom.QueryBuilder) (affected int64, e error) {
	db := rep.DB()
	if len(builders) > 0 {
		db = db.Model(&Artist{}).Where(query).Order(builders[0].Order()).Limit(builders[0].Limit()).Updates(value)
	} else {
		db = db.Model(&Artist{}).Where(query).Updates(value)
	}

	affected = db.RowsAffected
	e = db.Error
	return
}

// FindByWhereToUpdateArtists .
func FindByWhereToUpdateArtists(rep freedom.GORMRepository, query string, args []interface{}, value Artist, builders ...freedom.QueryBuilder) (affected int64, e error) {
	db := rep.DB()
	if len(builders) > 0 {
		db = db.Model(&Artist{}).Where(query, args...).Order(builders[0].Order()).Limit(builders[0].Limit()).Updates(value)
	} else {
		db = db.Model(&Artist{}).Where(query, args...).Updates(value)
	}

	affected = db.RowsAffected
	e = db.Error
	return
}
