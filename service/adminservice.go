package service

import (
	"errors"
	"godesaapps/model"
	"godesaapps/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CopyPegawaiToAdmin(idPegawai int, namalengkap string, pass string, roleId string) error
}

type adminServiceImpl struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminServiceImpl{repo}
}

func (s *adminServiceImpl) CopyPegawaiToAdmin(idPegawai int, namalengkap string, pass string, roleId string) error {
	// Mengecek apakah roleId valid menggunakan query ke database
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

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newAdmin := model.Admin{
		Id:          pegawai.Id,
		Email:       pegawai.Email,
		NikAdmin:    pegawai.NikAdmin,
		NamaLengkap: namalengkap,
		RoleId:      roleId, // gunakan roleId yang valid
		Pass:        string(hashedPassword),
	}

	return s.repo.InsertAdmin(newAdmin)
}
