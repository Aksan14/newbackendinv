package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"godesaapps/model"
	"godesaapps/service"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type pegawaiControllerImpl struct {
	PegawaiService service.PegawaiService
}

func NewPegawaiController(s service.PegawaiService) PegawaiController {
	return &pegawaiControllerImpl{PegawaiService: s}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func (c *pegawaiControllerImpl) CreatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseMultipartForm(40 << 20)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "Gagal parsing form",
			"data":    nil,
		})
		return
	}

	nip := r.FormValue("nip")
	email := r.FormValue("email")
	jabatan := r.FormValue("jabatan")
	namalengkap := r.FormValue("namalengkap")

	file, handler, err := r.FormFile("foto")
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "Foto wajib diunggah",
			"data":    nil,
		})
		return
	}
	defer file.Close()

	filename := fmt.Sprintf("pegawai/%s", handler.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Gagal menyimpan file",
			"data":    nil,
		})
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	pegawai := model.Pegawai{
		NIP:         nip,
		Email:       email,
		Jabatan:     jabatan,
		Foto:        filename,
		NamaLengkap: namalengkap,
	}

	err = c.PegawaiService.CreatePegawai(context.Background(), pegawai)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "Created",
		"message": "Pegawai berhasil ditambahkan",
		"data":    pegawai,
	})
}

func (c *pegawaiControllerImpl) GetAllPegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pegawaiList, err := c.PegawaiService.GetAllPegawai(context.Background())
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Gagal mengambil data",
			"data":    nil,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "OK",
		"message": "Data pegawai berhasil diambil",
		"data":    pegawaiList,
	})
}

func (c *pegawaiControllerImpl) GetPegawaiByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "ID tidak valid",
			"data":    nil,
		})
		return
	}

	pegawai, err := c.PegawaiService.GetPegawaiByID(context.Background(), id)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"status":  "Not Found",
			"message": "Pegawai tidak ditemukan",
			"data":    nil,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "OK",
		"message": "Pegawai berhasil ditemukan",
		"data":    pegawai,
	})
}

func (c *pegawaiControllerImpl) UpdatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "Gagal parsing form",
			"data":    nil,
		})
		return
	}

	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "ID tidak valid",
			"data":    nil,
		})
		return
	}

	nip := r.FormValue("nip")
	email := r.FormValue("email")
	jabatan := r.FormValue("jabatan")
	namalengkap := r.FormValue("namalengkap")

	oldData, err := c.PegawaiService.GetPegawaiByID(context.Background(), id)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"status":  "Not Found",
			"message": "Pegawai tidak ditemukan",
			"data":    nil,
		})
		return
	}

	fotoPath := oldData.Foto

	file, handler, err := r.FormFile("foto")
	if err == nil {
		defer file.Close()

		os.MkdirAll("pegawai", os.ModePerm)

		filename := fmt.Sprintf("pegawai/%s", handler.Filename)
		dst, err := os.Create(filename)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"status":  "Internal Server Error",
				"message": "Gagal menyimpan file foto baru",
				"data":    nil,
			})
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		fotoPath = filename
	}

	pegawai := model.Pegawai{
		ID:          id,
		NIP:         nip,
		Email:       email,
		NamaLengkap: namalengkap,
		Jabatan:     jabatan,
		Foto:        fotoPath,
	}

	err = c.PegawaiService.UpdatePegawai(context.Background(), pegawai)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "OK",
		"message": "Pegawai berhasil diperbarui",
		"data":    pegawai,
	})
}

func (c *pegawaiControllerImpl) DeletePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": "ID tidak valid",
			"data":    nil,
		})
		return
	}

	err = c.PegawaiService.DeletePegawai(context.Background(), id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Gagal menghapus pegawai",
			"data":    nil,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "OK",
		"message": "Pegawai berhasil dihapus",
		"data":    nil,
	})
}