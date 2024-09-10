package data

import (
	"database/sql"
	"errors"
)

type Models struct {
	Movies      MovieModel
	Tokens      TokenModel
	Permissions PermissionsModel
	Users       UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionsModel{DB: db},
		Users:       UserModel{DB: db},
	}
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflic")
)
