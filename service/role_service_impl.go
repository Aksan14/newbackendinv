package service

import (
	"context"
	"godesaapps/model"
)

type RoleService interface {
    GetAdminRole(ctx context.Context) (model.MstRole, error)
    GetRoleById(ctx context.Context, roleId string) (model.MstRole, error)
}
