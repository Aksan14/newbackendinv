package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"godesaapps/dto"
	"godesaapps/service"
	"github.com/julienschmidt/httprouter"
)

type response struct {
	Message string `json:"message"`
}

type RequestSuratController struct {
	service service.RequestSuratService
}

func NewRequestSuratController(service service.RequestSuratService) *RequestSuratController {
	return &RequestSuratController{service}
}

func (c *RequestSuratController) sendResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := response{
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Gagal encode JSON: %v\n", err)
	}
}

func (c *RequestSuratController) FindWargaByNik(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	nik := ps.ByName("nik")
	if nik == "" {
		c.sendResponse(w, http.StatusBadRequest, "NIK tidak boleh kosong")
		return
	}

	warga, err := c.service.FindByNik(nik)
	if err != nil {
		c.sendResponse(w, http.StatusInternalServerError, fmt.Sprintf("Gagal mencari data warga: %v", err))
		return
	}

	if warga == nil {
		c.sendResponse(w, http.StatusNotFound, "Data warga tidak ditemukan")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(warga); err != nil {
		fmt.Printf("Gagal encode JSON: %v\n", err)
	}
}

func (c *RequestSuratController) CreateRequestSurat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input dto.RequestSuratDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.sendResponse(w, http.StatusBadRequest, "Input tidak valid")
		return
	}

	err := c.service.RequestSurat(input)
	if err != nil {
		c.sendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.sendResponse(w, http.StatusCreated, "Permintaan surat berhasil dikirim")
}