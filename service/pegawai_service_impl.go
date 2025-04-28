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

// FindByNIP mencari pegawai berdasarkan NIP
func (s *PegawaiServiceImpl) FindByNIP(ctx context.Context, nip string) (*model.Pegawai, error) {
    pegawai, err := s.pegawaiRepo.FindByNIP(ctx, nip)
    if err != nil {
        log.Printf("Error finding pegawai by NIP %s: %v", nip, err)
        return nil, err
    }
    return pegawai, nil
}

// CreatePegawai menambahkan pegawai baru ke dalam database
func (s *PegawaiServiceImpl) CreatePegawai(ctx context.Context, p model.Pegawai) error {
    // Validasi NIP, pastikan tidak ada yang duplikat
    existingPegawai, err := s.FindByNIP(ctx, p.NIP)
    if err != nil {
        return err
    }
    if existingPegawai != nil {
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

// UpdatePegawai memperbarui data pegawai
func (s *PegawaiServiceImpl) UpdatePegawai(ctx context.Context, p model.Pegawai) error {
    if p.NIP == "" {
        return errors.New("NIP tidak boleh kosong")
    }

    // Ambil data pegawai asli berdasarkan ID
    oldPegawai, err := s.pegawaiRepo.GetPegawaiByID(ctx, p.ID)
    if err != nil {
        log.Printf("Error fetching pegawai by ID %d: %v", p.ID, err)
        return errors.New("Pegawai tidak ditemukan")
    }

    // Jika NIP berubah, periksa apakah NIP baru sudah digunakan oleh pegawai lain
    if p.NIP != oldPegawai.NIP {
        existingPegawai, err := s.FindByNIP(ctx, p.NIP)
        if err != nil {
            return err
        }
        if existingPegawai != nil && existingPegawai.ID != p.ID {
            return errors.New("NIP sudah terdaftar untuk pegawai lain")
        }
    }

    // Lanjutkan ke repositori untuk memperbarui pegawai
    return s.pegawaiRepo.UpdatePegawai(ctx, p)
}

// DeletePegawai menghapus pegawai berdasarkan ID
func (s *PegawaiServiceImpl) DeletePegawai(ctx context.Context, id int64) error {
    return s.pegawaiRepo.DeletePegawai(ctx, id)
}