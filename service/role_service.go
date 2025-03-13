package service

import (
    "context"
    "godesaapps/model"
    "godesaapps/repository"
)

// roleServiceImpl adalah implementasi dari RoleService.
type roleServiceImpl struct {
    roleRepo repository.RoleRepository
}

// NewRoleServiceImpl adalah constructor untuk membuat instance baru dari roleServiceImpl.
func NewRoleServiceImpl(roleRepo repository.RoleRepository) RoleService {
    return &roleServiceImpl{
        roleRepo: roleRepo,
    }
}

// GetAdminRole mengimplementasikan metode untuk mendapatkan role admin.
func (s *roleServiceImpl) GetAdminRole(ctx context.Context) (model.MstRole, error) {
    // Memanggil repository untuk mencari role admin
    role, err := s.roleRepo.FindAdminRole(ctx)
    if err != nil {
        return model.MstRole{}, err
    }
    return role, nil
}

// GetRoleById mengimplementasikan metode untuk mendapatkan role berdasarkan ID.
func (s *roleServiceImpl) GetRoleById(ctx context.Context, roleId string) (model.MstRole, error) {
    // Memanggil repository untuk mencari role berdasarkan ID
    role, err := s.roleRepo.FindMstRole(ctx, roleId)
    if err != nil {
        return model.MstRole{}, err
    }
    return role, nil
}