package controller

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/service"
	"godesaapps/util"
	"io"
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

// ========== Controller Register Warga dengan file upload ==========
func (controller *wargaControllerImpl) RegisterWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Gagal membaca form data", http.StatusBadRequest)
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
		http.Error(w, "File tidak ditemukan atau salah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("filewarga/%s", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	if err := controller.WargaService.InsertDataWarga(warga); err != nil {
		http.Error(w, "Gagal menyimpan data warga", http.StatusInternalServerError)
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

// ========== Controller Get All Warga ==========
func (controller *wargaControllerImpl) GetAllWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wargas, err := controller.WargaService.GetAllWarga()
	if err != nil {
		http.Error(w, "Gagal mengambil data warga", http.StatusInternalServerError)
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

// ========== Controller Update Data Warga ==========
func (controller *wargaControllerImpl) UpdateWarga(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var warga model.DataWarga
	if err := util.ReadFromRequestBody(r, &warga); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	if err := controller.WargaService.UpdateWarga(id, warga); err != nil {
		http.Error(w, "Gagal update data", http.StatusInternalServerError)
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

// ========== Controller Delete Data Warga ==========
func (controller *wargaControllerImpl) DeleteWarga(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	if err := controller.WargaService.DeleteWarga(id); err != nil {
		http.Error(w, "Gagal hapus data", http.StatusInternalServerError)
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
