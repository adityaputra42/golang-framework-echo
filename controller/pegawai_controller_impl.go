package controller

import (
	// "golang_framework_echo/service"

	"golang_framework_echo/helper"
	"golang_framework_echo/models/web"
	"golang_framework_echo/models/web/request"
	"golang_framework_echo/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PegawaiControllerImpl struct {
	PegawaiService service.PegawaiService
}

// Create implements PegawaiController.
func (controller *PegawaiControllerImpl) Create(c echo.Context) error {
	body := new(request.CreatePegawai)
	if err := c.Bind(body); err != nil {
		return err
	}
	result, err := controller.PegawaiService.Create(*body)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success create pegawai",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

// Delete implements PegawaiController.
func (controller *PegawaiControllerImpl) Delete(c echo.Context) error {
	pegawaiId := c.Param("pegawaiId")
	id, er := strconv.Atoi(pegawaiId)
	helper.PanicIfError(er)
	err := controller.PegawaiService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusOK,
		Message: "Succes delete pegawai",
	}
	return c.JSON(http.StatusOK, responseData)
}

// FindAll implements PegawaiController.
func (controller *PegawaiControllerImpl) FindAll(c echo.Context) error {
	result, err := controller.PegawaiService.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

// FindById implements PegawaiController.
func (controller *PegawaiControllerImpl) FindById(c echo.Context) error {
	pegawaiId := c.Param("pegawaiId")
	id, err := strconv.Atoi(pegawaiId)
	helper.PanicIfError(err)
	result, err := controller.PegawaiService.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

// Update implements PegawaiController.
func (controller *PegawaiControllerImpl) Update(c echo.Context) error {
	pegawaiId := c.Param("pegawaiId")
	id, err := strconv.Atoi(pegawaiId)
	helper.PanicIfError(err)
	body := new(request.UpdatePegawai)
	body.Id = id
	if err := c.Bind(body); err != nil {
		return err
	}
	result, err := controller.PegawaiService.Update(*body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	responseData := web.BaseResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    result,
	}
	return c.JSON(http.StatusCreated, responseData)
}

func NewPegawaiController(pegawaiService service.PegawaiService) *PegawaiControllerImpl {
	return &PegawaiControllerImpl{PegawaiService: pegawaiService}
}
