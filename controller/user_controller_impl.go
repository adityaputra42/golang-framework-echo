package controller

import (
	"golang_framework_echo/helper"
	"golang_framework_echo/models/web"
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claim["username"].(string)
		result, err := controller.UserService.FecthUser(username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		responseData := web.BaseResponse{
			Status:  http.StatusCreated,
			Message: "Success",
			Data:    result,
		}
		return c.JSON(http.StatusCreated, responseData)
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

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

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["expired"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    t,
	}

	return c.JSON(http.StatusCreated, responseData)

}

// Update implements UserController.
func (controller *UserControllerImpl) UpdatePassword(c echo.Context) error {

	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claim["username"].(string)

		body := new(request.UpdateUser)
		if err := c.Bind(body); err != nil {
			return err
		}
		result, err := controller.UserService.UpdatePassword(*body, username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		responseData := web.BaseResponse{
			Status:  http.StatusCreated,
			Message: "Success",
			Data:    result,
		}
		return c.JSON(http.StatusCreated, responseData)
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid Token"})
	}

}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
