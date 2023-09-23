package exception

import (
	"golang_framework_echo/helper"
	"golang_framework_echo/models/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}
	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		WebResponse := web.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Data:    exception.Error(),
		}
		helper.WriteToResponseBody(writer, WebResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		WebResponse := web.BaseResponse{
			Status:  http.StatusNotFound,
			Message: "Not Found",
			Data:    exception.Error,
		}
		helper.WriteToResponseBody(writer, WebResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	WebResponse := web.BaseResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Sever Error",
		Data:    err,
	}
	helper.WriteToResponseBody(writer, WebResponse)
}
