package repository

import (
	"errors"
	"golang_framework_echo/db"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
)

type PegawaiRepositoryImpl struct {
}

// Create implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Create(pegawai domain.Pegawai) domain.Pegawai {
	con := db.CreateCon()
	SQL := "insert into pegawai(nama,alamat,telepon) values(?,?,?)"
	statement, err := con.Prepare(SQL)
	helper.PanicIfError(err)
	result, err := statement.Exec(pegawai.Nama, pegawai.Alamat, pegawai.Telepon)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	pegawai.Id = int(id)
	return pegawai
}

// Delete implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Delete(pegawai domain.Pegawai) {
	con := db.CreateCon()
	SQL := "delete from pegawai where id = ?"
	_, err := con.Exec(SQL, pegawai.Id)
	helper.PanicIfError(err)

}

// FindById implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindById(pegawaiId int) (domain.Pegawai, error) {
	con := db.CreateCon()
	SQL := "select * from pegawai where id = ?"
	rows, err := con.Query(SQL, pegawaiId)
	pegawai := domain.Pegawai{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		return pegawai, nil
	} else {

		return pegawai, errors.New("pegawai is not found")
	}
}

// Update implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Update(pegawai domain.Pegawai) domain.Pegawai {
	con := db.CreateCon()
	SQL := "update pegawai set nama = ? ,alamat =?, telepon=? where id = ?"
	_, err := con.Exec(SQL, pegawai.Nama, pegawai.Alamat, pegawai.Telepon, pegawai.Id)
	helper.PanicIfError(err)

	return pegawai
}

// FindAll implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindAll() []domain.Pegawai {
	con := db.CreateCon()
	SQL := "select * from pegawai"
	rows, err := con.Query(SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var listPegawai []domain.Pegawai
	for rows.Next() {
		pegawai := domain.Pegawai{}
		err := rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		listPegawai = append(listPegawai, pegawai)
	}

	return listPegawai
}

func NewPegawaiRepository() PegawaiRepository {
	return &PegawaiRepositoryImpl{}
}
