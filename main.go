package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	validate := validator.New()
	db, err := app.NewDB()
	helper.PanicIfError(err)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := mux.NewRouter()
	r.HandleFunc("/api/categories", categoryController.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/api/categories/{categoryId}", categoryController.FindById).Methods(http.MethodGet)
	r.HandleFunc("/api/categories/{categoryId}", categoryController.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/categories", categoryController.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/categories/{categoryId}", categoryController.Delete).Methods(http.MethodDelete)

	err = http.ListenAndServe("localhost:9000", r)
	if err != nil {
		fmt.Println("Error ListenAndServe")
		return
	}

	fmt.Println("Server Up and Ready...")
}
