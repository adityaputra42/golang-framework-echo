package repository

import "golang_framework_echo/models/domain"

type UserRepository interface {
	Create(user domain.User) domain.User
	Update(user domain.User) domain.User
	Delete(user domain.User)
	FindById(UserId int) (domain.User, error)
	FindByUsername(Username string) (domain.User, error)
}
