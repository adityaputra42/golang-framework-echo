package repository

import (
	"errors"
	"golang_framework_echo/db"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/request"
	"golang_framework_echo/models/web"
	"net/http"
)

type PegawaiRepositoryImpl struct {
}

// Create implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Create(pegawai request.CreatePegawai) (web.BaseResponse, error) {
	con := db.CreateCon()
	SQL := "insert into pegawai(name,alamat,telepon) values(?,?,?)"
	result, err := con.Exec(SQL, pegawai.Name, pegawai.Alamat, pegawai.Telepon)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	resPegawai := domain.Pegawai{
		Id:      int(id),
		Name:    pegawai.Name,
		Alamat:  pegawai.Alamat,
		Telepon: pegawai.Telepon,
	}

	response := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    resPegawai,
	}
	return response, nil
}

// Delete implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Delete(pegawaiId int) (web.BaseResponse, error) {
	con := db.CreateCon()
	SQL := "delete from pegawai where id = ?"
	_, err := con.Exec(SQL, pegawaiId)
	helper.PanicIfError(err)
	response := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}
	return response, nil

}

// FindById implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindById(pegawaiId int) (web.BaseResponse, error) {
	con := db.CreateCon()
	SQL := "select * from pegawai where id = ?"
	rows, err := con.Query(SQL, pegawaiId)
	pegawai := domain.Pegawai{}
	helper.PanicIfError(err)
	defer rows.Close()
	response := web.BaseResponse{}
	if rows.Next() {
		err := rows.Scan(&pegawai.Id, &pegawai.Name, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = pegawai
		return response, nil

	} else {
		response.Status = http.StatusNotFound
		response.Message = "pegawai is not found"
		response.Data = pegawai
		return response, errors.New("pegawai is not found")
	}
}

// Update implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Update(pegawai request.UpdatePegawai) (web.BaseResponse, error) {
	con := db.CreateCon()
	SQL := "update pegawai set name = ? ,alamat =?, telepon=? where id = ?"
	_, err := con.Exec(SQL, pegawai.Name, pegawai.Alamat, pegawai.Telepon, pegawai.Id)
	helper.PanicIfError(err)

	response := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    pegawai,
	}
	return response, nil
}

// FindAll implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindAll() (web.BaseResponse, error) {
	con := db.CreateCon()
	SQL := "select * from pegawai"
	rows, err := con.Query(SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var listPegawai []domain.Pegawai
	for rows.Next() {
		pegawai := domain.Pegawai{}
		err := rows.Scan(&pegawai.Id, &pegawai.Name, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		listPegawai = append(listPegawai, pegawai)
	}
	response := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    listPegawai,
	}
	return response, nil
}

func NewPegawaiRepository() PegawaiRepository {
	return &PegawaiRepositoryImpl{}
}
