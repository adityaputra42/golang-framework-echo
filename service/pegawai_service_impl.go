package service

import (
	"context"
	"database/sql"
	"golang_framework_echo/exception"
	"golang_framework_echo/helper"
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/request"
	"golang_framework_echo/repository"

	"github.com/go-playground/validator/v10"
)

type PegawaiServiceImpl struct {
	PegawaiRepository repository.PegawaiRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

// Create implements PegawaiService.
func (service *PegawaiServiceImpl) Create(ctx context.Context, req request.CreatePegawai) domain.Pegawai {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	pegawai := domain.Pegawai{
		Name:    req.Name,
		Alamat:  req.Alamat,
		Telepon: req.Telepon,
	}
	pegawai = service.PegawaiRepository.Create(ctx, tx, pegawai)
	return pegawai
}

// Delete implements PegawaiService.
func (service *PegawaiServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	pegawai, err := service.PegawaiRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PegawaiRepository.Delete(ctx, tx, pegawai)
}

// FindAll implements PegawaiService.
func (service *PegawaiServiceImpl) FindAll(ctx context.Context) []domain.Pegawai {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	listPegawai := service.PegawaiRepository.FindAll(ctx, tx)
	return listPegawai
}

// FindById implements PegawaiService.
func (service *PegawaiServiceImpl) FindById(ctx context.Context, categoryId int) domain.Pegawai {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pegawai, err := service.PegawaiRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return pegawai
}

// Update implements PegawaiService.
func (service *PegawaiServiceImpl) Update(ctx context.Context, req request.UpdatePegawai) domain.Pegawai {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	pegawai, err := service.PegawaiRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pegawai.Name = req.Name
	pegawai.Alamat = req.Alamat
	pegawai.Telepon = req.Telepon

	pegawai = service.PegawaiRepository.Update(ctx, tx, pegawai)
	return pegawai
}

func NewPegawaiService(PegawaiRepository repository.PegawaiRepository, DB *sql.DB, Validate *validator.Validate) PegawaiService {
	return &PegawaiServiceImpl{PegawaiRepository: PegawaiRepository, DB: DB, Validate: Validate}
}
