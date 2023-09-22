package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
)

type PegawaiRepositoryImpl struct {
}

// Create implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai {
	SQL := "insert into pegawai(name,alamat,telepon) values(?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, pegawai.Name, pegawai.Alamat, pegawai.Telepon)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	pegawai.Id = int(id)
	return pegawai
}

// Delete implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, pegawai.Id)
	helper.PanicIfError(err)

}

// FindAll implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Pegawai {
	SQL := "select * from pegawai"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var listPegawai []domain.Pegawai
	for rows.Next() {
		pegawai := domain.Pegawai{}
		err := rows.Scan(&pegawai.Id, &pegawai.Name, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		listPegawai = append(listPegawai, pegawai)
	}
	return listPegawai
}

// FindById implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pegawaiId int) (domain.Pegawai, error) {
	SQL := "select * from pegawai where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, pegawaiId)
	pegawai := domain.Pegawai{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&pegawai.Id, &pegawai.Name, &pegawai.Alamat, &pegawai.Telepon)
		helper.PanicIfError(err)
		return pegawai, nil

	} else {
		return pegawai, errors.New("pegawai is not found")
	}
}

// Update implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai {
	SQL := "update pegawai set name = ? ,alamat =?, telepon=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, pegawai.Name, pegawai.Alamat, pegawai.Telepon, pegawai.Id)
	helper.PanicIfError(err)
	return pegawai
}

func NewPegawaiRepository() PegawaiRepository {
	return &PegawaiRepositoryImpl{}
}
