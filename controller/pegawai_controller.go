package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PegawaiController interface {
	CreatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetAllPegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetPegawaiByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeletePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
