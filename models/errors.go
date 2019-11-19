package models

import "errors"

var Err_CreateEntityRowsAffected = errors.New("Data Inserted but RowsAffected is not 1.")
