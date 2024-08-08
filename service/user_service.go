package service

import (
	"database/sql"
	"log"
	"myapp/domain"
	"myapp/domain/model"
)

type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UserService struct {
	DB   *sql.DB
	Repo domain.IUserRepository
}

func NewUserService(db *sql.DB, rep domain.IUserRepository) *UserService {
	return &UserService{
		DB:   db,
		Repo: rep,
	}
}

func (g *UserService) GetUserService(id int) *model.User {
	user, err := g.Repo.GetUser(g.DB, id)
	if err != nil {
		log.Print(err)
	}

	return user
}

func (g *UserService) ListUserService() []*model.User {
	users, err := g.Repo.GetUserList(g.DB)
	if err != nil {
		log.Print(err)
	}

	return users
}

func (g *UserService) CreateUserService(userRequest UserRequest) error {
	bodyParam := &model.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
	err := g.Repo.RegisterUser(g.DB, bodyParam)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (g *UserService) UpdateUserService(id int, userRequest UserRequest) error {
	bodyParam := &model.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
	err := g.Repo.UpdateUser(g.DB, id, bodyParam)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (g *UserService) DeleteUserService(id int) error {
	err := g.Repo.DeleteUser(g.DB, id)
	if err != nil {
		log.Print(err)
	}
	return err
}
