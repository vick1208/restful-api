package exception

import (
	"net/http"
	"vick1208/restful-api/helper"
	"vick1208/restful-api/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(write http.ResponseWriter, req *http.Request, err interface{}) {

	if notFoundError(write, req, err) {
		return
	}
	if validationErrors(write, req, err) {
		return
	}

	internalServerError(write, req, err)
}

func validationErrors(write http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToRespBody(write, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(write http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusNotFound)

		webResp := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteToRespBody(write, webResp)
		return true

	} else {
		return false
	}
}

func internalServerError(write http.ResponseWriter, request *http.Request, err interface{}) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToRespBody(write, webResponse)
}
