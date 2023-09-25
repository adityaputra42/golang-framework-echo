package service

import (
	"database/sql"
	"fmt"
	"golang_framework_echo/exception"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/models/web/response"
	"golang_framework_echo/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Create implements UserService.
func (service *UserServiceImpl) Create(req request.CreateUser) (response.UserResponse, error) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	user = service.UserRepository.Create(user)
	return helper.ToUserResponse(user), nil

}

// Delete implements UserService.
func (service *UserServiceImpl) Delete(userId int) error {
	user, err := service.UserRepository.FindById(userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.UserRepository.Delete(user)
	return nil
}

// FecthUser implements UserService.
func (service *UserServiceImpl) FecthUser(userId int) (response.UserResponse, error) {
	pegawai, err := service.UserRepository.FindById(userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToUserResponse(pegawai), nil
}

// Login implements UserService.
func (service *UserServiceImpl) Login(username string, password string) (bool, error) {
	user, err := service.UserRepository.FindByUsername(username)
	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}
	match, err := helper.CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("hash and password doesn't match")
		return false, err
	}
	return true, nil
}

// Update implements UserService.
func (service *UserServiceImpl) Update(req request.UpdateUser) (response.UserResponse, error) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	user, err := service.UserRepository.FindById(req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user.Username = req.Username
	user.Password = req.Password

	user = service.UserRepository.Update(user)
	return helper.ToUserResponse(user), nil

}

func NewUserRepository(userRespository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRespository, Validate: validate}
}
