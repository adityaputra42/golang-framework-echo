package service

import (
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/models/web/response"
)

type UserService interface {
	Create(req request.CreateUser) (response.UserResponse, error)
	Login(username, password string) (bool, error)
	UpdatePassword(req request.UpdateUser, username string) (response.UserResponse, error)
	Delete(userId int) error
	FecthUser(username string) (response.UserResponse, error)
}
