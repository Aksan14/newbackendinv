package repository

import (
    "context"
    "database/sql"
    "godesaapps/model"
)

type pegawaiRepositoryImpl struct {
    DB *sql.DB
}

func NewPegawaiRepository(db *sql.DB) PegawaiRepository {
    return &pegawaiRepositoryImpl{DB: db}
}

func (r *pegawaiRepositoryImpl) CreatePegawai(ctx context.Context, p model.Pegawai) error {
    query := "INSERT INTO pegawai (nip, namalengkap, email, jabatan, foto) VALUES (?, ?, ?, ?, ?)"
    _, err := r.DB.ExecContext(ctx, query, p.NIP, p.NamaLengkap, p.Email, p.Jabatan, p.Foto)
    return err
}

func (r *pegawaiRepositoryImpl) GetAllPegawai(ctx context.Context) ([]model.Pegawai, error) {
    rows, err := r.DB.QueryContext(ctx, "SELECT id, nip, namalengkap, email, jabatan, foto FROM pegawai")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pegawaiList []model.Pegawai
    for rows.Next() {
        var p model.Pegawai
        err := rows.Scan(&p.ID, &p.NIP, &p.NamaLengkap, &p.Email, &p.Jabatan, &p.Foto)
        if err != nil {
            return nil, err
        }
        pegawaiList = append(pegawaiList, p)
    }

    return pegawaiList, nil
}

func (r *pegawaiRepositoryImpl) GetPegawaiByID(ctx context.Context, id int64) (model.Pegawai, error) {
	
    query := "SELECT id, nip, namalengkap, email, jabatan, foto FROM pegawai WHERE id = ?"
    row := r.DB.QueryRowContext(ctx, query, id)

    var p model.Pegawai
    err := row.Scan(&p.ID, &p.NIP, &p.NamaLengkap, &p.Email, &p.Jabatan, &p.Foto)
    return p, err
}

func (r *pegawaiRepositoryImpl) UpdatePegawai(ctx context.Context, p model.Pegawai) error {
    query := "UPDATE pegawai SET nip=?, namalengkap=?, email=?, jabatan=?, foto=? WHERE id=?"
    _, err := r.DB.ExecContext(ctx, query, p.NIP, p.NamaLengkap, p.Email, p.Jabatan, p.Foto, p.ID)
    return err
}

func (r *pegawaiRepositoryImpl) DeletePegawai(ctx context.Context, id int64) error {
    _, err := r.DB.ExecContext(ctx, "DELETE FROM pegawai WHERE id=?", id)
    return err
}
