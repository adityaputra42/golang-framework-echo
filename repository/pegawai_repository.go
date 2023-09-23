package repository

import (
	"golang_framework_echo/models/request"
	"golang_framework_echo/models/web"
)

type PegawaiRepository interface {
	Create(pegawai request.CreatePegawai) (web.BaseResponse, error)
	Update(pegawai request.UpdatePegawai) (web.BaseResponse, error)
	Delete(pegawaiId int) (web.BaseResponse, error)
	FindById(pegawaiId int) (web.BaseResponse, error)
	FindAll() (web.BaseResponse, error)
}
