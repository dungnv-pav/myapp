package handler

import (
	"log"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserHandler struct {
	Service IUserService
}

func NewUserHandle(s IUserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (u *UserHandler) GetUserHandle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
	}

	user := u.Service.GetUserService(id)
	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) ListUserHandle(c echo.Context) error {
	users := u.Service.ListUserService()
	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) CreateUserHandle(c echo.Context) error {
	var userRequest UserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := c.Validate(&userRequest); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	bodyParam := service.UserRequest{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	u.Service.CreateUserService(bodyParam)
	return c.JSON(http.StatusCreated, "created success")
}

func (u *UserHandler) UpdateUserHandle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
	}

	var userRequest UserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := c.Validate(&userRequest); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	bodyParam := service.UserRequest{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	u.Service.UpdateUserService(id, bodyParam)
	return c.JSON(http.StatusOK, "updated success")
}

func (u *UserHandler) DeleteUserHandle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
	}

	u.Service.DeleteUserService(id)
	return c.JSON(http.StatusOK, "deleted success")
}
