package controller

import (
	"encoding/json"
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/service"
	"godesaapps/util"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

type wargaControllerImpl struct {
	WargaService service.WargaService
}

func NewWargaController(wargaService service.WargaService) WargaController {
	return &wargaControllerImpl{
		WargaService: wargaService,
	}
}

func (controller *userControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	requestCreate := dto.CreateUserRequest{}
	util.ReadFromRequestBody(request, &requestCreate)

	responseDTO := controller.UserService.CreateUser(request.Context(), requestCreate)
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "create user successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginRequest dto.LoginUserRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := controller.UserService.LoginUser(r.Context(), loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    token,
		Message: "token generate successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) ReadUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseDTO := controller.UserService.ReadUser(request.Context())
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "read user successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[7:]
	claims := &service.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
		return
	}

	email := claims.Email
	userResponse, err := controller.UserService.GetUserInfoByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    userResponse,
		Message: "success login to user",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) ForgotPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req dto.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := controller.UserService.ForgotPassword(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    nil,
		Message: "Reset link dikirim ke email",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    queryParams := r.URL.Query()
    token := queryParams.Get("token")
    if token == "" {
        http.Error(w, "Token tidak ditemukan", http.StatusBadRequest)
        return
    }


    var req dto.ResetPasswordRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Set token ke request DTO
    req.Token = token

    err := controller.UserService.ResetPassword(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := dto.ResponseList{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Password berhasil direset",
    }

    w.Header().Set("Content-Type", "application/json")
    util.WriteToResponseBody(w, response)
}

func (controller *wargaControllerImpl) RegisterWarga(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse request body
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Gagal membaca form data", http.StatusBadRequest)
		return
	}

	// data warga body
	var wargaRequest dto.WargaRequest
	wargaRequest.NIK = r.FormValue("nik")
	wargaRequest.NamaLengkap = r.FormValue("nama_lengkap")
	wargaRequest.Alamat = r.FormValue("alamat")
	wargaRequest.JenisSurat = r.FormValue("jenis_surat")
	wargaRequest.Keterangan = r.FormValue("keterangan")
	wargaRequest.NoHP = r.FormValue("no_hp")

	// ambl file
	file, handler, err := r.FormFile("file_upload")
	if err != nil {
		http.Error(w, "File tidak ditemukan atau salah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// smpn file
	filePath := fmt.Sprintf("filewarga/%s", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
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

	// Simpan data
	err = controller.WargaService.RegisterWarga(wargaModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responseeee wargaaa
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Warga berhasil didaftarkan dan file berhasil diunggah",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}


