package controller

import (
	"golang_framework_echo/service"

	echo "github.com/labstack/echo/v4"
)

type PegawaiControllerImpl struct {
	PegawaiService service.PegawaiService
}

// Create implements PegawaiController.
func (controller *PegawaiControllerImpl) Create(c *echo.Context) {
	// body := request.CreatePegawai{}
	// helper.ReadFromRequestBody(req, &body)
	// pegawai := controller.PegawaiService.Create(req.Context(), body)
	// response := web.BaseResponse{
	// 	Status:  200,
	// 	Message: "Ok",
	// 	Data:    pegawai,
	// }
	// helper.WriteToResponseBody(writer, response)

}

// Delete implements PegawaiController.
func (controller *PegawaiControllerImpl) Delete(c *echo.Context) {
	panic("unimplemented")
}

// FindAll implements PegawaiController.
func (controller *PegawaiControllerImpl) FindAll(c *echo.Context) {
	panic("unimplemented")
}

// FindById implements PegawaiController.
func (controller *PegawaiControllerImpl) FindById(c *echo.Context) {
	panic("unimplemented")
}

// Update implements PegawaiController.
func (controller *PegawaiControllerImpl) Update(c *echo.Context) {
	panic("unimplemented")
}

func NewPegawaiController(pegawaiService service.PegawaiService) PegawaiController {
	return &PegawaiControllerImpl{PegawaiService: pegawaiService}
}

// // Create implements PegawaiController.
// func (controller *PegawaiControllerImpl) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	body := request.CreatePegawai{}
// 	helper.ReadFromRequestBody(req, &body)
// 	pegawai := controller.PegawaiService.Create(req.Context(), body)
// 	response := web.BaseResponse{
// 		Status:  200,
// 		Message: "Ok",
// 		Data:    pegawai,
// 	}
// 	helper.WriteToResponseBody(writer, response)
// }

// // Delete implements PegawaiController.
// func (controller *PegawaiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	pegawaiId := params.ByName("pegawaiId")
// 	id, err := strconv.Atoi(pegawaiId)
// 	helper.PanicIfError(err)
// 	controller.PegawaiService.Delete(request.Context(), id)
// 	response := web.BaseResponse{
// 		Status:  200,
// 		Message: "Ok",
// 	}

// 	helper.WriteToResponseBody(writer, response)
// }

// // FindAll implements PegawaiController.
// func (controller *PegawaiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	listPegawai := controller.PegawaiService.FindAll(request.Context())
// 	webResponse := web.BaseResponse{
// 		Status:  200,
// 		Message: "Ok",
// 		Data:    listPegawai,
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)

// }

// // FindById implements PegawaiController.
// func (controller *PegawaiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	pegawaiId := params.ByName("pegawaiId")
// 	id, err := strconv.Atoi(pegawaiId)
// 	helper.PanicIfError(err)

// 	pegawai := controller.PegawaiService.FindById(request.Context(), id)
// 	webResponse := web.BaseResponse{
// 		Status:  200,
// 		Message: "Ok",
// 		Data:    pegawai,
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)
// }

// // Update implements PegawaiController.
// func (controller *PegawaiControllerImpl) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	body := request.UpdatePegawai{}
// 	helper.ReadFromRequestBody(req, &body)

// 	pegawaiId := params.ByName("pegawaiId")
// 	id, err := strconv.Atoi(pegawaiId)
// 	helper.PanicIfError(err)
// 	body.Id = id

// 	pegawai := controller.PegawaiService.Update(req.Context(), body)
// 	webResponse := web.BaseResponse{
// 		Status:  200,
// 		Message: "Ok",
// 		Data:    pegawai,
// 	}
// 	helper.WriteToResponseBody(writer, webResponse)
// }
