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
        // Password:    request.Pass,
        NamaLengkap: request.NamaLengkap,
    }
}

func ToRoleResponse(role model.MstRole) dto.RoleResponse {
    return dto.RoleResponse{
        IdRole:   role.IdRole,
        RoleName: role.RoleName,
    }
}

func ToRoleListResponse(roles []model.MstRole) []dto.RoleResponse {
    var roleResponses []dto.RoleResponse
    for _, role := range roles {
        roleResponses = append(roleResponses, dto.RoleResponse{
            IdRole:   role.IdRole,
            RoleName: role.RoleName,
        })
    }
    return roleResponses
}

func ToRoleModel(request dto.RoleRequest) model.MstRole {
    return model.MstRole{
        IdRole:   request.IdRole,
        RoleName: request.RoleName,
    }
}

func ToUserResponseWithRole(user model.User, role model.MstRole) dto.UserResponse {
    return dto.UserResponse{
        Id:        user.Id,
        Nikadmin:  user.Nikadmin,
        Email:     user.Email,
        Role:      ToRoleResponse(role),
    }
}