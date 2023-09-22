package repository

import (
	"context"
	"database/sql"
	"golang_framework_echo/models/domain"
)

type PegawaiRepositoryImpl struct {
}

// Create implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai {
	panic("unimplemented")
}

// Delete implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) {
	panic("unimplemented")
}

// FindAll implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Pegawai {
	panic("unimplemented")
}

// FindById implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pegawaiId int) (domain.Pegawai, error) {
	panic("unimplemented")
}

// Update implements PegawaiRepository.
func (repository *PegawaiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai {
	panic("unimplemented")
}

func NewPegawaiRepository() PegawaiRepository {
	return &PegawaiRepositoryImpl{}
}
