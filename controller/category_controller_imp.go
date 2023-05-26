package controller

import (
	"net/http"
	"strconv"
	"vick1208/restful-api/helper"
	"vick1208/restful-api/model/web"
	"vick1208/restful-api/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(catSer service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: catSer,
	}
}

func (controller *CategoryControllerImpl) Create(write http.ResponseWriter, req *http.Request, params httprouter.Params) {

	cCR := web.CategoryCreateRequest{}
	helper.ReadFromReqBody(req, &cCR)

	categoryResp := controller.CategoryService.Create(req.Context(), cCR)
	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResp,
	}
	helper.WriteToRespBody(write, webResp)
}

func (controller *CategoryControllerImpl) Update(write http.ResponseWriter, req *http.Request, params httprouter.Params) {

	cUR := web.CategoryUpdateRequest{}
	helper.ReadFromReqBody(req, &cUR)

	cId := params.ByName("categoryId")
	id, err := strconv.Atoi(cId)
	helper.PanicIE(err)
	cUR.Id = id

	categoryResp := controller.CategoryService.Update(req.Context(), cUR)
	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResp,
	}
	helper.WriteToRespBody(write, webResp)
}

func (controller *CategoryControllerImpl) Delete(write http.ResponseWriter, req *http.Request, params httprouter.Params) {

	cId := params.ByName("categoryId")
	id, err := strconv.Atoi(cId)
	helper.PanicIE(err)

	controller.CategoryService.Delete(req.Context(), id)
	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToRespBody(write, webResp)
}

func (controller *CategoryControllerImpl) FindById(write http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cId := params.ByName("categoryId")
	id, err := strconv.Atoi(cId)
	helper.PanicIE(err)

	categoryResp := controller.CategoryService.FindById(req.Context(), id)
	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResp,
	}
	helper.WriteToRespBody(write, webResp)
}

func (controller *CategoryControllerImpl) FindAll(write http.ResponseWriter, req *http.Request, params httprouter.Params) {

	categoryResps := controller.CategoryService.FindAll(req.Context())
	webResp := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResps,
	}
	helper.WriteToRespBody(write, webResp)
}
