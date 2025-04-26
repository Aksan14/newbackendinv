package controller

import (
	"encoding/json"
	"net/http"
	"godesaapps/service"
	"github.com/julienschmidt/httprouter"
)

type AdminController struct {
	service service.AdminService
}

func NewAdminController(service service.AdminService) *AdminController {
	return &AdminController{service}
}

type CreateAdminRequest struct {
	ID          int    `json:"id"`
	NamaLengkap string `json:"namalengkap"`
	Password    string `json:"pass"`
	RoleId      string `json:"role_id"`
}

func (c *AdminController) CreateAdminFromPegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req CreateAdminRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := c.service.CopyPegawaiToAdmin(req.ID, req.NamaLengkap, req.Password, req.RoleId)
	if err != nil {
		http.Error(w, "Gagal membuat admin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin berhasil dibuat"})
}
