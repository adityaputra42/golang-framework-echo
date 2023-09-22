package service

import (
	"context"
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/request"
)

type PegawaiService interface {
	Create(ctx context.Context, req request.CreatePegawai) domain.Pegawai
	Update(ctx context.Context, req request.UpdatePegawai) domain.Pegawai
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) domain.Pegawai
	FindAll(ctx context.Context) []domain.Pegawai
}
