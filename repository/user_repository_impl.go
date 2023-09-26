package repository

import (
	"errors"
	"golang_framework_echo/db"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
)

type UserRepositoryImpl struct {
}

// Create implements UserRepository.
func (*UserRepositoryImpl) Create(user domain.User) domain.User {
	con := db.CreateCon()
	hash, _ := helper.HashPassword(user.Password)
	SQL := "insert into users(username,password) values(?,?)"
	statement, err := con.Prepare(SQL)
	helper.PanicIfError(err)
	result, err := statement.Exec(user.Username, hash)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

// Delete implements UserRepository.
func (*UserRepositoryImpl) Delete(user domain.User) {
	con := db.CreateCon()
	SQL := "delete from users where id = ?"
	_, err := con.Exec(SQL, user.Id)
	helper.PanicIfError(err)
}

// FecthUser implements UserRepository.
func (*UserRepositoryImpl) FindById(UserId int) (domain.User, error) {
	con := db.CreateCon()
	SQL := "select * from users where id = ?"
	rows, err := con.Query(SQL, UserId)
	user := domain.User{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {

		return user, errors.New("user is not found")
	}
}

// Update implements UserRepository.
func (*UserRepositoryImpl) Update(user domain.User) domain.User {
	con := db.CreateCon()
	hash, _ := helper.HashPassword(user.Password)
	SQL := "update users set username = ? ,pasword =? where id = ?"
	_, err := con.Exec(SQL, user.Username, hash, user.Id)
	helper.PanicIfError(err)
	return user
}

// FindByUsername implements UserRepository.
func (*UserRepositoryImpl) FindByUsername(Username string) (domain.User, error) {
	con := db.CreateCon()
	SQL := "select * from users where username = ?"
	rows, err := con.Query(SQL, Username)
	user := domain.User{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {

		return user, errors.New("user is not found")
	}
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
