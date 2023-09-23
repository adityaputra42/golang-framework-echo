package service

import (
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/models/web/response"
)

type PegawaiService interface {
	Create(req request.CreatePegawai) (response.PegawaiResponse, error)
	Update(req request.UpdatePegawai) (response.PegawaiResponse, error)
	Delete(pegawaiId int) error
	FindById(pegawaiId int) (response.PegawaiResponse, error)
	FindAll() ([]response.PegawaiResponse, error)
}
