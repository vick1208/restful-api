package main

import (
	"net/http"
	"vick1208/restful-api/app"
	"vick1208/restful-api/controller"
	"vick1208/restful-api/helper"
	"vick1208/restful-api/middleware"
	"vick1208/restful-api/repository"
	"vick1208/restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDatabase()
	validate := validator.New()

	categoryRepo := repository.NewCategoryRepo()
	categoryService := service.NewCategoryService(categoryRepo, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIE(err)
}
