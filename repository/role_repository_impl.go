package repository

import (
	"context"
	"database/sql"
	"errors"
	"godesaapps/model"
)

func (r *userRepositoryImpl) FindRoleById(ctx context.Context, tx *sql.Tx, roleId string) (model.MstRole, error) {
	query := "SELECT id, name FROM role_admin WHERE id = ?"

	row := tx.QueryRowContext(ctx, query, roleId)

	var role model.MstRole
	err := row.Scan(&role.IdRole, &role.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.MstRole{}, errors.New("role tidak ditemukan")
		}
		return model.MstRole{}, err
	}

	return role, nil
}
