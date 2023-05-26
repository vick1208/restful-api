package service

import (
	"context"
	"database/sql"
	"vick1208/restful-api/exception"
	"vick1208/restful-api/helper"
	"vick1208/restful-api/model/domain"
	"vick1208/restful-api/model/web"
	"vick1208/restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImp struct {
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImp{
		CategoryRepo: categoryRepo,
		DB:           DB,
		Validate:     validate,
	}
}

func (service *CategoryServiceImp) Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(req)
	helper.PanicIE(err)

	tx, err := service.DB.Begin()
	helper.PanicIE(err)
	defer helper.RollbackOrCommit(tx)

	category := domain.Category{
		Name: req.Name,
	}
	category = service.CategoryRepo.Insert(ctx, tx, category)
	return helper.ToCategoryResp(category)
}

func (service *CategoryServiceImp) Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(req)
	helper.PanicIE(err)

	tx, err := service.DB.Begin()
	helper.PanicIE(err)
	defer helper.RollbackOrCommit(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name

	category = service.CategoryRepo.Update(ctx, tx, category)
	return helper.ToCategoryResp(category)

}

func (service *CategoryServiceImp) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIE(err)
	defer helper.RollbackOrCommit(tx)
	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CategoryRepo.Delete(ctx, tx, category)
}

func (service *CategoryServiceImp) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIE(err)
	defer helper.RollbackOrCommit(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResp(category)
}

func (service *CategoryServiceImp) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIE(err)
	defer helper.RollbackOrCommit(tx)

	categories := service.CategoryRepo.FindAll(ctx, tx)

	return helper.ToCategoryResps(categories)
}
