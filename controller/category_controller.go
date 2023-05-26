package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(write http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(write http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(write http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(write http.ResponseWriter, req *http.Request, params httprouter.Params)
}
