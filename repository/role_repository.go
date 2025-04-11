package repository

import (
	"context"
	"database/sql"
	"godesaapps/model"
)

type RoleRepository interface {
	FindRoleById(ctx context.Context, tx *sql.Tx, roleId string) (model.MstRole, error)
}
