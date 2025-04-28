package repository

import (
    "context"
    "database/sql"
    "fmt"
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
    if err != nil {
        return fmt.Errorf("error creating pegawai: %v", err)
    }
    return nil
}

func (r *pegawaiRepositoryImpl) GetAllPegawai(ctx context.Context) ([]model.Pegawai, error) {
    rows, err := r.DB.QueryContext(ctx, "SELECT id, nip, namalengkap, email, jabatan, foto FROM pegawai")
    if err != nil {
        return nil, fmt.Errorf("error querying all pegawai: %v", err)
    }
    defer rows.Close()

    var pegawaiList []model.Pegawai
    for rows.Next() {
        var p model.Pegawai
        err := rows.Scan(&p.ID, &p.NIP, &p.NamaLengkap, &p.Email, &p.Jabatan, &p.Foto)
        if err != nil {
            return nil, fmt.Errorf("error scanning pegawai: %v", err)
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
    if err != nil {
        if err == sql.ErrNoRows {
            return model.Pegawai{}, fmt.Errorf("pegawai with ID %d not found", id)
        }
        return model.Pegawai{}, fmt.Errorf("error querying pegawai by ID: %v", err)
    }
    return p, nil
}

func (r *pegawaiRepositoryImpl) UpdatePegawai(ctx context.Context, p model.Pegawai) error {
    query := "UPDATE pegawai SET nip=?, namalengkap=?, email=?, jabatan=?, foto=? WHERE id=?"
    _, err := r.DB.ExecContext(ctx, query, p.NIP, p.NamaLengkap, p.Email, p.Jabatan, p.Foto, p.ID)
    if err != nil {
        return fmt.Errorf("error updating pegawai: %v", err)
    }
    return nil
}

func (r *pegawaiRepositoryImpl) DeletePegawai(ctx context.Context, id int64) error {
    _, err := r.DB.ExecContext(ctx, "DELETE FROM pegawai WHERE id=?", id)
    if err != nil {
        return fmt.Errorf("error deleting pegawai: %v", err)
    }
    return nil
}

func (r *pegawaiRepositoryImpl) FindByNIP(ctx context.Context, nip string) (*model.Pegawai, error) {
    query := "SELECT id, nip, namalengkap, email, jabatan, foto FROM pegawai WHERE nip = ?"
    row := r.DB.QueryRowContext(ctx, query, nip)

    var p model.Pegawai
    err := row.Scan(&p.ID, &p.NIP, &p.NamaLengkap, &p.Email, &p.Jabatan, &p.Foto)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil 
        }
        return nil, fmt.Errorf("error querying pegawai by NIP: %v", err)
    }
    return &p, nil
}