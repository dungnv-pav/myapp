package router

import (
	"log"
	constant "myapp/constants"
	"myapp/database"
	"myapp/handler"
	"myapp/repository"
	"myapp/service"
	"myapp/validator"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewValidator()
	database.Connect()
	db, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	rep := repository.NewRepository()

	userService := service.NewUserService(db, rep)
	userHandler := handler.NewUserHandle(userService)

	e.GET(constant.UserDetailGetPath, userHandler.GetUserHandle)
	e.GET(constant.UserGetPath, userHandler.ListUserHandle)
	e.POST(constant.UserCreatePath, userHandler.CreateUserHandle)
	e.PUT(constant.UserUpdatePath, userHandler.UpdateUserHandle)
	e.DELETE(constant.UserDeletePath, userHandler.DeleteUserHandle)

	return e
}
