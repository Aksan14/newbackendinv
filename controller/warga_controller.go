package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WargaController interface {
	RegisterWarga(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	InsertDataWarga(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAllWarga(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateWarga(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteWarga(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
