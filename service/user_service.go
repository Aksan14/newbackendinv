package service

import (
	"context"
	"godesaapps/dto"
	"godesaapps/model"
)

type UserService interface {
	CreateUser(ctx context.Context, userRequest dto.CreateUserRequest) dto.UserResponse
	LoginUser(ctx context.Context, loginRequest dto.LoginUserRequest) (string, error)
	GenerateJWT(email string) (string, error)
	ReadUser(ctx context.Context) []dto.UserResponse
	FindById(ctx context.Context, id string) (dto.UserResponse, error)
	GetUserInfoByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	ForgotPassword(request dto.ForgotPasswordRequest) error
	ResetPassword(request dto.ResetPasswordRequest) error
}


type WargaService interface {
	RegisterWarga(warga model.Warga) error
}
