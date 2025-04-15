package service

import (
	"errors"
	"fmt"
	"godesaapps/model"
	"godesaapps/repository"
)

type wargaServiceImpl struct {
	repo repository.WargaRepository
}

func NewWargaService(repo repository.WargaRepository) WargaService {
	return &wargaServiceImpl{repo: repo}
}

// Validasi input sebelum insert
func (s *wargaServiceImpl) RegisterWarga(warga model.Warga) error {
	if warga.NIK == "" || warga.NamaLengkap == "" || warga.Alamat == "" || warga.JenisSurat == "" || warga.NoHP == "" {
		return errors.New("semua field wajib diisi")
	}
	return s.repo.InsertWarga(warga)
}

func (s *wargaServiceImpl) InsertDataWarga(warga model.DataWarga) error {
	existing, err := s.repo.FindByNIK(warga.NIK)
	if err != nil {
		return err
	}
	if existing != nil {
		return fmt.Errorf("NIK sudah terdaftar")
	}

	return s.repo.InsertDataWarga(warga)
}


func (s *wargaServiceImpl) GetAllWarga() ([]model.DataWarga, error) {
	return s.repo.GetAllWarga()
}

func (s *wargaServiceImpl) UpdateWarga(id int, warga model.DataWarga) error {
	return s.repo.UpdateWarga(id, warga)
}

func (s *wargaServiceImpl) DeleteWarga(id int) error {
	return s.repo.DeleteWarga(id)
}
