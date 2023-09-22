package repository

import (
	"context"
	"database/sql"
	"golang_framework_echo/models/domain"
)

type PegawaiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai
	Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai
	Delete(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai)
	FindById(ctx context.Context, tx *sql.Tx, pegawaiId int) (domain.Pegawai, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pegawai
}
