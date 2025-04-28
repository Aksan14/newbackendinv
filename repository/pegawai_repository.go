package repository

import (
    "context"
    "godesaapps/model"
)

type PegawaiRepository interface {
    CreatePegawai(ctx context.Context, pegawai model.Pegawai) error
    GetAllPegawai(ctx context.Context) ([]model.Pegawai, error)
    GetPegawaiByID(ctx context.Context, id int64) (model.Pegawai, error)
    UpdatePegawai(ctx context.Context, pegawai model.Pegawai) error
    DeletePegawai(ctx context.Context, id int64) error
    FindByNIP(ctx context.Context, nip string) (*model.Pegawai, error)
}