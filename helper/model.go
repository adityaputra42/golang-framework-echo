package helper

import (
	"golang_framework_echo/models/domain"
	"golang_framework_echo/models/web/response"
)

func ToPegawaiResponse(pegawai domain.Pegawai) response.PegawaiResponse {
	return response.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Alamat:  pegawai.Alamat,
		Telepon: pegawai.Telepon,
	}

}

func ToPegawaiResponses(listPegawai []domain.Pegawai) []response.PegawaiResponse {
	var listPegawaiResponse []response.PegawaiResponse
	for _, pegawai := range listPegawai {
		listPegawaiResponse = append(listPegawaiResponse, ToPegawaiResponse(pegawai))
	}
	return listPegawaiResponse
}
