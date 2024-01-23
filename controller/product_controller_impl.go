package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"
	"restful-api/service"
	"strconv"
)

type ProductControllerImpl struct {
	service service.ProductService
}

func NewProductController(product service.ProductService) ProductController {
	return &ProductControllerImpl{
		service: product,
	}
}

func (controller ProductControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prodCreate := web.ProductCreateRequest{}

	errorDecode := json.NewDecoder(request.Body).Decode(&prodCreate)
	helper.PanicIfError(errorDecode)

	result := controller.service.Save(request.Context(), prodCreate)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}

	writter.Header().Add("Content-Type", "application/json")
	errorEncode := json.NewEncoder(writter).Encode(WebResponse)
	helper.PanicIfError(errorEncode)

}

func (controller ProductControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prodUpdate := web.ProductUpdateRequest{}

	errorDecode := json.NewDecoder(request.Body).Decode(&prodUpdate)
	helper.PanicIfError(errorDecode)

	prodId := params.ByName("productId")
	idProd, err := strconv.Atoi(prodId)

	helper.PanicIfError(err)
	prodUpdate.IdProduct = idProd

	result := controller.service.Update(request.Context(), prodUpdate)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	writter.Header().Add("Content-Type", "application/json")
	errEncode := json.NewEncoder(writter).Encode(webResponse)
	helper.PanicIfError(errEncode)
}

func (controller ProductControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prodId := params.ByName("productId")
	id, err := strconv.Atoi(prodId)
	helper.PanicIfError(err)

	errDelete := controller.service.Delete(request.Context(), id)
	helper.PanicIfError(errDelete)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writter.Header().Add("Content-Type", "application/json")
	errEncode := json.NewEncoder(writter).Encode(webResponse)
	helper.PanicIfError(errEncode)

}

func (controller ProductControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	prodId := params.ByName("productId")
	id, err := strconv.Atoi(prodId)
	helper.PanicIfError(err)

	result := controller.service.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	writter.Header().Add("Content-Type", "application/json")
	errEncode := json.NewEncoder(writter).Encode(webResponse)
	helper.PanicIfError(errEncode)
}
