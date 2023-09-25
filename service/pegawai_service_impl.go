package service

import (
	"golang_framework_echo/exception"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/models/web/response"
	"golang_framework_echo/repository"

	"github.com/go-playground/validator/v10"
)

type PegawaiServiceImpl struct {
	PegawaiRepository repository.PegawaiRepository
	Validate          *validator.Validate
}

// Create implements PegawaiService.
func (service *PegawaiServiceImpl) Create(req request.CreatePegawai) (response.PegawaiResponse, error) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	pegawai := domain.Pegawai{
		Nama:    req.Nama,
		Alamat:  req.Alamat,
		Telepon: req.Telepon,
	}
	pegawai = service.PegawaiRepository.Create(pegawai)
	return helper.ToPegawaiResponse(pegawai), nil

}

// Delete implements PegawaiService.
func (service *PegawaiServiceImpl) Delete(pegawaiId int) error {
	pegawai, err := service.PegawaiRepository.FindById(pegawaiId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PegawaiRepository.Delete(pegawai)
	return nil
}

// FindAll implements PegawaiService.
func (service *PegawaiServiceImpl) FindAll() ([]response.PegawaiResponse, error) {
	listPegawai := service.PegawaiRepository.FindAll()
	return helper.ToPegawaiResponses(listPegawai), nil
}

// FindById implements PegawaiService.
func (service *PegawaiServiceImpl) FindById(pegawaiId int) (response.PegawaiResponse, error) {
	pegawai, err := service.PegawaiRepository.FindById(pegawaiId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPegawaiResponse(pegawai), nil
}

// Update implements PegawaiService.
func (service *PegawaiServiceImpl) Update(req request.UpdatePegawai) (response.PegawaiResponse, error) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	pegawai, err := service.PegawaiRepository.FindById(req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pegawai.Nama = req.Nama
	pegawai.Alamat = req.Alamat
	pegawai.Telepon = req.Telepon

	pegawai = service.PegawaiRepository.Update(pegawai)
	return helper.ToPegawaiResponse(pegawai), nil

}

func NewPegawaiService(pegawaiRepository repository.PegawaiRepository,
	Validate *validator.Validate) PegawaiService {
	return &PegawaiServiceImpl{PegawaiRepository: pegawaiRepository, Validate: Validate}

}
