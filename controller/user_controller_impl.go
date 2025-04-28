package controller

import (
	"encoding/json"
	"godesaapps/dto"
	"godesaapps/service"
	"godesaapps/util"
	"net/http"
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


func (controller *userControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	users, err := controller.UserService.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "Gagal mengambil data user", http.StatusInternalServerError)
		return
	}

	for i := range users {
		users[i].Password = "********"
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    users,
		Message: "Data semua user berhasil diambil",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) DeleteUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")

    err := controller.UserService.DeleteUser(r.Context(), id)
    if err != nil {
        response := dto.ResponseList{
            Code:    http.StatusNotFound,
            Status:  "Not Found",
            Message: "User not found or failed to delete",
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        util.WriteToResponseBody(w, response)
        return
    }

    response := dto.ResponseList{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "User deleted successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    util.WriteToResponseBody(w, response)
}


// POST /api/user/sign-up
func (controller *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestCreate dto.CreateUserRequest
	util.ReadFromRequestBody(r, &requestCreate)

	existingUser, err := controller.UserService.FindByNIK(r.Context(), requestCreate.Nikadmin)
	if err == nil && existingUser != nil {
		response := dto.ResponseList{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "NIK sudah terdaftar",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		util.WriteToResponseBody(w, response)
		return
	}

	responseDTO := controller.UserService.CreateUser(r.Context(), requestCreate)
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "User berhasil dibuat",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginRequest dto.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		response := dto.ResponseList{
			Code:    http.StatusBadRequest,
			Status:  "FAILED",
			Message: "Invalid input",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		util.WriteToResponseBody(w, response)
		return
	}

	token, err := controller.UserService.LoginUser(r.Context(), loginRequest)
	if err != nil {
		response := dto.ResponseList{
			Code:    http.StatusUnauthorized,
			Status:  "FAILED",
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		util.WriteToResponseBody(w, response)
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    token,
		Message: "Token generated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 {
		http.Error(w, "Missing or Invalid Authorization Header", http.StatusUnauthorized)
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

	if claims.Nikadmin == "" {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userResponse, err := controller.UserService.GetUserInfoByNikAdmin(r.Context(), claims.Nikadmin)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	role, err := controller.UserService.GetRoleByUserId(r.Context(), userResponse.Role.IdRole)
	if err != nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	roleResponse := dto.RoleResponse{
		IdRole:   role.IdRole,
		RoleName: role.RoleName,
	}

	response := dto.ResponseList{
		Code:   http.StatusOK,
		Status: "OK",
		Data: dto.UserResponse{
			Id:          userResponse.Id,
			Nikadmin:    userResponse.Nikadmin,
			Email:       userResponse.Email,
			NamaLengkap: userResponse.NamaLengkap,
			Role:        roleResponse,
		},
		Message: "Success fetching user information",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	util.WriteToResponseBody(w, response)
}

// POST /api/user/forgot-password
func (controller *userControllerImpl) ForgotPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req dto.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := dto.ResponseList{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		util.WriteToResponseBody(w, response)
		return
	}

	err := controller.UserService.ForgotPassword(req)
	if err != nil {
		response := dto.ResponseList{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		util.WriteToResponseBody(w, response)
		return
	}

	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Reset link dikirim ke email",
	}
	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) ResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    token := r.URL.Query().Get("token")

    if token == "" {
        response := dto.ResponseList{
            Code:    http.StatusBadRequest,
            Status:  "Bad Request",
            Message: "Token tidak ditemukan",
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        util.WriteToResponseBody(w, response)
        return
    }

    var req dto.ResetPasswordRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response := dto.ResponseList{
            Code:    http.StatusBadRequest,
            Status:  "Bad Request",
            Message: "Request tidak valid",
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        util.WriteToResponseBody(w, response)
        return
    }

    req.Token = token
    err := controller.UserService.ResetPassword(req)
    if err != nil {
        response := dto.ResponseList{
            Code:    http.StatusBadRequest,
            Status:  "Bad Request",
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        util.WriteToResponseBody(w, response)
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