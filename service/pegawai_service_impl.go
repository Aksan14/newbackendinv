package service

import (
    "context"
    "errors"
    "godesaapps/model"
    "godesaapps/repository"
    "log"
)

// PegawaiServiceImpl adalah implementasi dari interface PegawaiService
type PegawaiServiceImpl struct {
    pegawaiRepo repository.PegawaiRepository
}

// NewPegawaiService membuat instance PegawaiServiceImpl
func NewPegawaiService(pegawaiRepo repository.PegawaiRepository) PegawaiService {
    return &PegawaiServiceImpl{pegawaiRepo: pegawaiRepo}
}

// ValidateNIP memeriksa apakah NIP sudah ada di database
func (s *PegawaiServiceImpl) ValidateNIP(ctx context.Context, nip string) (bool, error) {
    pegawaiList, err := s.pegawaiRepo.GetAllPegawai(ctx)
    if err != nil {
        log.Println("Error fetching all pegawai:", err)
        return false, err
    }

    for _, p := range pegawaiList {
        if p.NIP == nip {
            return true, nil // NIP sudah ada
        }
    }
    return false, nil // NIP tidak ditemukan
}

// CreatePegawai menambahkan pegawai baru ke dalam database
func (s *PegawaiServiceImpl) CreatePegawai(ctx context.Context, p model.Pegawai) error {
    // Validasi NIP, pastikan tidak ada yang duplikat
    isDuplicate, err := s.ValidateNIP(ctx, p.NIP)
    if err != nil {
        return err
    }
    if isDuplicate {
        return errors.New("NIP sudah terdaftar")
    }

    // Jika valid, lanjutkan ke repositori untuk menambahkan pegawai
    return s.pegawaiRepo.CreatePegawai(ctx, p)
}

// GetAllPegawai mengambil semua data pegawai dari database
func (s *PegawaiServiceImpl) GetAllPegawai(ctx context.Context) ([]model.Pegawai, error) {
    return s.pegawaiRepo.GetAllPegawai(ctx)
}

// GetPegawaiByID mengambil data pegawai berdasarkan ID
func (s *PegawaiServiceImpl) GetPegawaiByID(ctx context.Context, id int64) (model.Pegawai, error) {
    return s.pegawaiRepo.GetPegawaiByID(ctx, id)
}

func (s *PegawaiServiceImpl) UpdatePegawai(ctx context.Context, p model.Pegawai) error {
    if p.NIP != "" {
        isDuplicate, err := s.ValidateNIP(ctx, p.NIP)
        if err != nil {
            return err
        }
        if isDuplicate {
            return errors.New("NIP sudah terdaftar")
        }
    }

    return s.pegawaiRepo.UpdatePegawai(ctx, p)
}

func (s *PegawaiServiceImpl) DeletePegawai(ctx context.Context, id int64) error {
    return s.pegawaiRepo.DeletePegawai(ctx, id)
}
