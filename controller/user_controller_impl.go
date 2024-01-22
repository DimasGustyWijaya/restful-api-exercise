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

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(user service.UserService) UserController {
	return &UserControllerImpl{
		UserService: user,
	}
}

func (controller UserControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}

	// Deocde data JSON dari request dan di kirim ke userCreateRequest
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(&userCreateRequest)
	helper.PanicIfError(errDecode)

	//Panggil layer Service Create
	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)

	//Inisialisasi objek untuk response web
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	//Objek response web yang sudah di inisialisasi di encode ke json dan dikirim ke writter
	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)

}

func (controller UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	errDelete := controller.UserService.Delete(request.Context(), userId)
	helper.PanicIfError(errDelete)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)

}

func (controller UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}

	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(&userUpdateRequest)
	helper.PanicIfError(errDecode)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)

}

func (controller UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	user := controller.UserService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	errEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errEncode)

}
