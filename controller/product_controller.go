package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
