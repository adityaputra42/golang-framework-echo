package controller

import (
	"golang_framework_echo/helper"
	"golang_framework_echo/models/web"
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/service"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// Create implements UserController.
func (controller *UserControllerImpl) Create(c echo.Context) error {
	body := new(request.CreateUser)
	if err := c.Bind(body); err != nil {
		return err
	}
	result, err := controller.UserService.Create(*body)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success create user",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

// Delete implements UserController.
func (controller *UserControllerImpl) Delete(c echo.Context) error {
	pegawaiId := c.Param("userId")
	id, er := strconv.Atoi(pegawaiId)
	helper.PanicIfError(er)
	err := controller.UserService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Succes delete user",
	}
	return c.JSON(http.StatusOK, responseData)
}

// FetchUSer implements UserController.
func (controller *UserControllerImpl) FetchUSer(c echo.Context) error {
	pegawaiId := c.Param("pegawaiId")
	id, err := strconv.Atoi(pegawaiId)
	helper.PanicIfError(err)
	result, err := controller.UserService.FecthUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

// Login implements UserController.
func (controller *UserControllerImpl) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := controller.UserService.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if !res {
		return echo.ErrUnauthorized
	}
	
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    res,
	}

	return c.JSON(http.StatusCreated, responseData)

}

// Update implements UserController.
func (controller *UserControllerImpl) Update(c echo.Context) error {
	pegawaiId := c.Param("userId")
	id, err := strconv.Atoi(pegawaiId)
	helper.PanicIfError(err)
	body := new(request.UpdateUser)
	body.Id = id
	if err := c.Bind(body); err != nil {
		return err
	}
	result, err := controller.UserService.Update(*body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
