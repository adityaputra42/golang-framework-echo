package repository

import (
	"golang_framework_echo/models/domain"
)

type PegawaiRepository interface {
	Create(pegawai domain.Pegawai) domain.Pegawai
	Update(pegawai domain.Pegawai) domain.Pegawai
	Delete(pegawai domain.Pegawai)
	FindById(pegawaiId int) (domain.Pegawai, error)
	FindAll() []domain.Pegawai
}
