package service

import (
	"context"
	"godesaapps/dto"
)

type UserService interface {
    GetUserInfoByNikAdmin(ctx context.Context, nikadmin string) (dto.UserResponse, error)
    CreateUser(ctx context.Context, userRequest dto.CreateUserRequest) dto.UserResponse
    GenerateJWT(email, nikadmin, roleadmin, namalengkap string) (string, error)
    LoginUser(ctx context.Context, loginRequest dto.LoginUserRequest) (string, error)
    FindByNIK(ctx context.Context, nik string) (*dto.UserResponse, error)
    ForgotPassword(request dto.ForgotPasswordRequest) error
    ResetPassword(request dto.ResetPasswordRequest) error
    GetRoleByUserId(ctx context.Context, roleID string) (dto.RoleResponse, error) 
}

