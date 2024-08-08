package handler

import (
	"myapp/domain/model"
	"myapp/service"
)

type IUserService interface {
	GetUserService(id int) *model.User
	ListUserService() []*model.User
	CreateUserService(userRequest service.UserRequest) error
	UpdateUserService(id int, userRequest service.UserRequest) error
	DeleteUserService(id int) error
}
