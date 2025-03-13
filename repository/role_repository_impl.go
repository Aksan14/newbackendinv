package repository

import (
    "context"
    "database/sql"
    "errors"
    "godesaapps/model"
)

type roleRepositoryImpl struct {
    DB *sql.DB
}

func NewRoleRepositoryImpl(db *sql.DB) RoleRepository {
    return &roleRepositoryImpl{
        DB: db,
    }
}

func (roleRepository roleRepositoryImpl) FindMstRole(ctx context.Context, roleId string) (model.MstRole, error) {
    query := "SELECT id, name FROM role_admin WHERE id = $1 LIMIT 1"
    row := roleRepository.DB.QueryRowContext(ctx, query, roleId)

    var role model.MstRole
    err := row.Scan(&role.IdRole, &role.RoleName)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.MstRole{}, errors.New("role not found")
        }
        return model.MstRole{}, err
    }

    return role, nil
}

func (roleRepository roleRepositoryImpl) FindAdminRole(ctx context.Context) (model.MstRole, error) {
    query := "SELECT id, name FROM role_admin WHERE is_admin = true LIMIT 1"
    row := roleRepository.DB.QueryRowContext(ctx, query)

    var role model.MstRole
    err := row.Scan(&role.IdRole, &role.RoleName)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.MstRole{}, errors.New("admin role not found")
        }
        return model.MstRole{}, err
    }

    return role, nil
}