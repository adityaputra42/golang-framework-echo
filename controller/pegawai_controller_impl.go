package controller

import (
	// "golang_framework_echo/service"

	"golang_framework_echo/helper"
	"golang_framework_echo/models/request"
	"golang_framework_echo/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PegawaiControllerImpl struct {
	PegawaiRepository repository.PegawaiRepository
}

// Create implements PegawaiController.
func (controller *PegawaiControllerImpl) Create(c echo.Context) error {
	body := new(request.CreatePegawai)
	if err := c.Bind(body); err != nil {
		return err
	}
	result, err := controller.PegawaiRepository.Create(*body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, result)
}

// Delete implements PegawaiController.
func (controller *PegawaiControllerImpl) Delete(c echo.Context) error {
	pegawaiId := c.Param("pegawaiId")
	id, er := strconv.Atoi(pegawaiId)
	helper.PanicIfError(er)
	result, err := controller.PegawaiRepository.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// FindAll implements PegawaiController.
func (controller *PegawaiControllerImpl) FindAll(c echo.Context) error {
	result, err := controller.PegawaiRepository.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// FindById implements PegawaiController.
func (controller *PegawaiControllerImpl) FindById(c echo.Context) error {
	pegawaiId := c.QueryParam("pegawaiId")
	id, err := strconv.Atoi(pegawaiId)
	helper.PanicIfError(err)
	result, err := controller.PegawaiRepository.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
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
	result, err := controller.PegawaiRepository.Update(*body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func NewPegawaiController(pegawaiRepository repository.PegawaiRepository) PegawaiController {
	return &PegawaiControllerImpl{PegawaiRepository: pegawaiRepository}
}
