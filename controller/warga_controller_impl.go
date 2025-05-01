package controller

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/service"
	"godesaapps/util"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type wargaControllerImpl struct {
	WargaService service.WargaService
}

func NewWargaController(wargaService service.WargaService) WargaController {
	return &wargaControllerImpl{
		WargaService: wargaService,
	}
}

// responseError adalah struktur untuk respons error JSON
type responseError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// writeErrorResponse menulis respons error dalam format JSON
func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	response := responseError{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) RegisterWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Gagal membaca form data")
		return
	}

	var wargaRequest dto.WargaRequest
	wargaRequest.NIK = r.FormValue("nik")
	wargaRequest.NamaLengkap = r.FormValue("nama_lengkap")
	wargaRequest.Alamat = r.FormValue("alamat")
	wargaRequest.JenisSurat = r.FormValue("jenis_surat")
	wargaRequest.Keterangan = r.FormValue("keterangan")
	wargaRequest.NoHP = r.FormValue("no_hp")

	file, handler, err := r.FormFile("file_upload")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "File tidak ditemukan atau salah")
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("filewarga/%s", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Gagal menyimpan file")
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Gagal menyimpan file")
		return
	}

	wargaRequest.FileUpload = filePath

	wargaModel := model.Warga{
		NIK:         wargaRequest.NIK,
		NamaLengkap: wargaRequest.NamaLengkap,
		Alamat:      wargaRequest.Alamat,
		JenisSurat:  wargaRequest.JenisSurat,
		Keterangan:  wargaRequest.Keterangan,
		FileUpload:  wargaRequest.FileUpload,
		NoHP:        wargaRequest.NoHP,
	}

	if err = controller.WargaService.RegisterWarga(wargaModel); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Warga berhasil didaftarkan dan file berhasil diunggah",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) InsertDataWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var warga model.DataWarga

	if err := util.ReadFromRequestBody(r, &warga); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Data tidak valid")
		return
	}

	if err := controller.WargaService.InsertDataWarga(warga); err != nil {
		if err.Error() == "NIK sudah terdaftar" {
			writeErrorResponse(w, http.StatusConflict, "NIK sudah terdaftar")
		} else {
			writeErrorResponse(w, http.StatusInternalServerError, "Gagal menyimpan data warga: "+err.Error())
		}
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Data warga berhasil ditambahkan",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) GetAllWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wargas, err := controller.WargaService.GetAllWarga()
	if err != nil {
		log.Printf("Error saat mengambil data warga: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data warga: "+err.Error())
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Data warga ditemukan",
		Data:    wargas,
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) UpdateWarga(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var warga model.DataWarga
	if err := util.ReadFromRequestBody(r, &warga); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Data tidak valid")
		return
	}

	if err := controller.WargaService.UpdateWarga(id, warga); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Gagal update data: "+err.Error())
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Data berhasil diperbarui",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) DeleteWarga(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := controller.WargaService.DeleteWarga(id); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Gagal hapus data: "+err.Error())
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Data warga berhasil dihapus",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}