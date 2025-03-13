package repository

import (
	"context"
	"godesaapps/model"
)

type RoleRepository interface {
    FindMstRole(ctx context.Context, roleId string) (model.MstRole, error)
    FindAdminRole(ctx context.Context) (model.MstRole, error)
}