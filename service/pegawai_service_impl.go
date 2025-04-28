package service

import (
    "context"
    "errors"
    "godesaapps/model"
    "godesaapps/repository"
    "log"
)

type PegawaiServiceImpl struct {
    pegawaiRepo repository.PegawaiRepository
}

func NewPegawaiService(pegawaiRepo repository.PegawaiRepository) PegawaiService {
    return &PegawaiServiceImpl{pegawaiRepo: pegawaiRepo}
}

func (s *PegawaiServiceImpl) FindByNIP(ctx context.Context, nip string) (*model.Pegawai, error) {
    pegawai, err := s.pegawaiRepo.FindByNIP(ctx, nip)
    if err != nil {
        log.Printf("Error finding pegawai by NIP %s: %v", nip, err)
        return nil, err
    }
    return pegawai, nil
}

func (s *PegawaiServiceImpl) CreatePegawai(ctx context.Context, p model.Pegawai) error {
    existingPegawai, err := s.FindByNIP(ctx, p.NIP)
    if err != nil {
        return err
    }
    if existingPegawai != nil {
        return errors.New("NIP sudah terdaftar")
    }

    return s.pegawaiRepo.CreatePegawai(ctx, p)
}

func (s *PegawaiServiceImpl) GetAllPegawai(ctx context.Context) ([]model.Pegawai, error) {
    return s.pegawaiRepo.GetAllPegawai(ctx)
}

func (s *PegawaiServiceImpl) GetPegawaiByID(ctx context.Context, id int64) (model.Pegawai, error) {
    return s.pegawaiRepo.GetPegawaiByID(ctx, id)
}

func (s *PegawaiServiceImpl) UpdatePegawai(ctx context.Context, p model.Pegawai) error {
    if p.NIP == "" {
        return errors.New("NIP tidak boleh kosong")
    }

    oldPegawai, err := s.pegawaiRepo.GetPegawaiByID(ctx, p.ID)
    if err != nil {
        log.Printf("Error fetching pegawai by ID %d: %v", p.ID, err)
        return errors.New("Pegawai tidak ditemukan")
    }

    if p.NIP != oldPegawai.NIP {
        existingPegawai, err := s.FindByNIP(ctx, p.NIP)
        if err != nil {
            return err
        }
        if existingPegawai != nil && existingPegawai.ID != p.ID {
            return errors.New("NIP sudah terdaftar untuk pegawai lain")
        }
    }

    return s.pegawaiRepo.UpdatePegawai(ctx, p)
}

func (s *PegawaiServiceImpl) DeletePegawai(ctx context.Context, id int64) error {
    return s.pegawaiRepo.DeletePegawai(ctx, id)
}