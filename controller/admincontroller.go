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
	ID       int    `json:"id"`
	Password string `json:"pass"`
	RoleId   string `json:"role_id"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *AdminController) CreateAdminFromPegawai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request body
	var req CreateAdminRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, Response{Message: "Invalid request body"})
		return
	}

	// Validate request
	if req.ID <= 0 {
		sendJSONResponse(w, http.StatusBadRequest, Response{Message: "ID pegawai harus lebih besar dari 0"})
		return
	}
	if req.Password == "" {
		sendJSONResponse(w, http.StatusBadRequest, Response{Message: "Password tidak boleh kosong"})
		return
	}
	validRoles := map[string]bool{"ROLE000": true, "ROLE001": true, "ROLE002": true}
	if req.RoleId == "" || !validRoles[req.RoleId] {
		sendJSONResponse(w, http.StatusBadRequest, Response{Message: "Role ID tidak valid, harus salah satu dari: ROLE000, ROLE001, ROLE002"})
		return
	}

	// Call service to create admin
	err := c.service.CopyPegawaiToAdmin(req.ID, req.Password, req.RoleId)
	if err != nil {
		switch err.Error() {
		case "Pegawai tidak ditemukan":
			sendJSONResponse(w, http.StatusNotFound, Response{Message: err.Error()})
		case "Pegawai sudah terdaftar sebagai admin":
			sendJSONResponse(w, http.StatusConflict, Response{Message: err.Error()})
		default:
			sendJSONResponse(w, http.StatusInternalServerError, Response{Message: "Gagal membuat admin: " + err.Error()})
		}
		return
	}

	// Send success response
	response := Response{
		Message: "Admin berhasil dibuat",
	}
	sendJSONResponse(w, http.StatusCreated, response)
}

// sendJSONResponse mengirimkan respons JSON dengan status HTTP yang ditentukan
func sendJSONResponse(w http.ResponseWriter, status int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}