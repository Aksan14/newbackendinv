package util

import (
	"godesaapps/dto"
	"godesaapps/model"
)

func ToUserResponse(user model.User) dto.UserResponse {
	return dto.UserResponse{
		Id:        user.Id,
		Nikadmin:  user.Nikadmin,
		Email:     user.Email,
		// Password dihapus dari response untuk alasan keamanan
	}
}

func ToUserListResponse(users []model.User) []dto.UserResponse {
	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserModel(request dto.CreateUserRequest) model.User {
	return model.User{
		Nikadmin: request.Nikadmin,
		Email:    request.Email,
		Password: request.Pass, // Pastikan di-hash sebelum disimpan ke DB
	}
}
