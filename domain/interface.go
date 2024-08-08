package domain

import (
	"database/sql"
	"myapp/domain/model"
)

type IUserRepository interface {
	GetUser(db *sql.DB, userID int) (*model.User, error)
	GetUserList(db *sql.DB) ([]*model.User, error)
	RegisterUser(db *sql.DB, user *model.User) error
	UpdateUser(db *sql.DB, userID int, user *model.User) error
	DeleteUser(db *sql.DB, userID int) error
}
