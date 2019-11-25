package models

import "errors"

var ErrCreateAlbumRowAffected = errors.New("Album create error: RowsAffected is not 1.")
var ErrUpdateAlbumRowAffected = errors.New("Album updated error: RowsAffected is not 1 or 0.")
var ErrUpdateAlbumNoAffected = errors.New("Album updated error: RowsAffected is 0.")
var ErrDeleteAlbumRowAffected = errors.New("Album deleted error: RowsAffected is not 1 or 0.")
var ErrDeleteAlbumNoAffected = errors.New("Album deleted error: RowsAffected is 0.")
