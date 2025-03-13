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
    }
}

func ToUserListResponse(users []model.User) []dto.UserResponse {
    var userResponses []dto.UserResponse
    for _, user := range users {
        userResponses = append(userResponses, dto.UserResponse{
            Id:       user.Id,
            Nikadmin: user.Nikadmin,
            Email:    user.Email,
        })
    }
    return userResponses
}

func ToUserModel(request dto.CreateUserRequest) model.User {
    return model.User{
        Nikadmin:    request.Nikadmin,
        Email:       request.Email,
        Password:    request.Pass,
        NamaLengkap: request.NamaLengkap,
    }
}

func ToRoleResponse(role model.MstRole) dto.RoleResponse {
    return dto.RoleResponse{
        IdRole:   role.IdRole,
        RoleName: role.RoleName,
        IsAdmin:  role.IsAdmin,
    }
}

// ToRoleListResponse mengonversi slice model.MstRole ke slice dto.RoleResponse
func ToRoleListResponse(roles []model.MstRole) []dto.RoleResponse {
    var roleResponses []dto.RoleResponse
    for _, role := range roles {
        roleResponses = append(roleResponses, dto.RoleResponse{
            IdRole:   role.IdRole,
            RoleName: role.RoleName,
            IsAdmin:  role.IsAdmin,
        })
    }
    return roleResponses
}

// ToRoleModel mengonversi dto.RoleRequest ke model.MstRole
func ToRoleModel(request dto.RoleRequest) model.MstRole {
    return model.MstRole{
        IdRole:   request.IdRole,
        RoleName: request.RoleName,
        IsAdmin:  request.IsAdmin,
    }
}

// ToUserResponseWithRole mengonversi model.User dan model.MstRole ke dto.UserResponse dengan informasi role
func ToUserResponseWithRole(user model.User, role model.MstRole) dto.UserResponse {
    return dto.UserResponse{
        Id:        user.Id,
        Nikadmin:  user.Nikadmin,
        Email:     user.Email,
        Role:      ToRoleResponse(role),
    }
}

// IsAdminRole memeriksa apakah suatu role adalah role admin
func IsAdminRole(role model.MstRole) bool {
    return role.IsAdmin
}