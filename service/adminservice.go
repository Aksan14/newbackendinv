package service

import (
	"errors"
	"godesaapps/model"
	"godesaapps/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CopyPegawaiToAdmin(idPegawai int, pass string, roleId string) error
}

type adminServiceImpl struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminServiceImpl{repo}
}

func (s *adminServiceImpl) CopyPegawaiToAdmin(idPegawai int, pass string, roleId string) error {
	roleExists, err := s.repo.RoleExists(roleId)
	if err != nil {
		return err
	}
	if !roleExists {
		return errors.New("invalid roleId")
	}

	pegawai, err := s.repo.FindPegawaiById(idPegawai)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newAdmin := model.Admin{
		Id:          pegawai.Id,
		Email:       pegawai.Email,
		NikAdmin:    pegawai.NamaLengkap, 
		NamaLengkap: pegawai.NikAdmin,    
		RoleId:      roleId,
		Pass:        string(hashedPassword),
	}

	return s.repo.InsertAdmin(newAdmin)
}