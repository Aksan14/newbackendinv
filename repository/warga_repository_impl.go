package repository

import (
	"database/sql"
	"godesaapps/model"
)

type wargaRepositoryImpl struct {
	db *sql.DB
}

func NewWargaRepository(db *sql.DB) WargaRepository {
	return &wargaRepositoryImpl{db: db}
}

func (r *wargaRepositoryImpl) InsertWarga(warga model.Warga) error {
	query := `INSERT INTO warga (nik, nama_lengkap, alamat, jenis_surat, keterangan, file_upload, no_hp, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err := r.db.Exec(query, warga.NIK, warga.NamaLengkap, warga.Alamat, warga.JenisSurat, warga.Keterangan, warga.FileUpload, warga.NoHP)
	return err
}

func (r *wargaRepositoryImpl) InsertDataWarga(w model.DataWarga) error {
	_, err := r.db.Exec(`
		INSERT INTO datawarga 
		(nik, nama_lengkap, tempat_lahir, tanggal_lahir, jenis_kelamin, pendidikan, pekerjaan, agama, status_pernikahan, kewarganegaraan, alamat) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		w.NIK, w.NamaLengkap, w.TempatLahir, w.TanggalLahir, w.JenisKelamin, w.Pendidikan, w.Pekerjaan, w.Agama, w.StatusPernikahan, w.Kewarganegaraan, w.Alamat,
	)
	return err
}

func (r *wargaRepositoryImpl) GetAllWarga() ([]model.DataWarga, error) {
	rows, err := r.db.Query(`
		SELECT id, nik, nama_lengkap, tempat_lahir, tanggal_lahir, jenis_kelamin, pendidikan, pekerjaan, agama, status_pernikahan, kewarganegaraan, alamat
		FROM datawarga`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wargas []model.DataWarga
	for rows.Next() {
		var warga model.DataWarga
		err := rows.Scan(
			&warga.ID, &warga.NIK, &warga.NamaLengkap, &warga.TempatLahir, &warga.TanggalLahir,
			&warga.JenisKelamin, &warga.Pendidikan, &warga.Pekerjaan, &warga.Agama,
			&warga.StatusPernikahan, &warga.Kewarganegaraan, &warga.Alamat,
		)
		if err != nil {
			return nil, err
		}
		wargas = append(wargas, warga)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return wargas, nil
}


func (r *wargaRepositoryImpl) UpdateWarga(id int, w model.DataWarga) error {
	_, err := r.db.Exec("UPDATE datawarga SET nik=?, nama_lengkap=?, tempat_lahir=?, tanggal_lahir=?, jenis_kelamin=?, pendidikan=?, pekerjaan=?, agama=?, status_pernikahan=?, kewarganegaraan=?, alamat=? WHERE id=?",
		w.NIK, w.NamaLengkap, w.TempatLahir, w.TanggalLahir, w.JenisKelamin, w.Pendidikan, w.Pekerjaan, w.Agama, w.StatusPernikahan, w.Kewarganegaraan, w.Alamat, id)
	return err
}

func (r *wargaRepositoryImpl) FindByNIK(nik string) (*model.DataWarga, error) {
	row := r.db.QueryRow("SELECT id, nik FROM datawarga WHERE nik = ?", nik)

	var warga model.DataWarga
	err := row.Scan(&warga.ID, &warga.NIK)
	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}
	return &warga, nil
}


func (r *wargaRepositoryImpl) DeleteWarga(id int) error {
	_, err := r.db.Exec("DELETE FROM datawarga WHERE id=?", id)
	return err
}

