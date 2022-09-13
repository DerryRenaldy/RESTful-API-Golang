package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryServices service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryServices: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := c.CategoryServices.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	params := mux.Vars(request)
	categoryId, err := strconv.Atoi(params["categoryId"])
	helper.PanicIfError(err)
	fmt.Println(categoryId)

	categoryUpdateRequest.ID = categoryId

	categoryResponse := c.CategoryServices.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	categoryId, err := strconv.Atoi(params["categoryId"])
	helper.PanicIfError(err)

	c.CategoryServices.Delete(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	categoryId, err := strconv.Atoi(params["categoryId"])
	helper.PanicIfError(err)

	categoryResponse := c.CategoryServices.FindById(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	categoryResponse := c.CategoryServices.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
