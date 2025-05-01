package repository

import (
	"database/sql"
	"godesaapps/model"
	"log"
)

type requestSuratRepositoryImpl struct {
	db *sql.DB
}

func NewRequestSuratRepository(db *sql.DB) RequestSuratRepository {
	return &requestSuratRepositoryImpl{db}
}

func (r *requestSuratRepositoryImpl) FindByNik(nik string) (*model.DataWarga, error) {
	query := `SELECT nik, nama_lengkap, tempat_lahir, tanggal_lahir, jenis_kelamin,
	                 pendidikan, pekerjaan, agama, status_pernikahan, kewarganegaraan, alamat
	          FROM datawarga WHERE nik = ?`

	row := r.db.QueryRow(query, nik)

	var warga model.DataWarga
	err := row.Scan(&warga.NIK, &warga.NamaLengkap, &warga.TempatLahir, &warga.TanggalLahir,
		&warga.JenisKelamin, &warga.Pendidikan, &warga.Pekerjaan, &warga.Agama,
		&warga.StatusPernikahan, &warga.Kewarganegaraan, &warga.Alamat)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error fetching data:", err)
		return nil, err
	}

	return &warga, nil
}

func (r *requestSuratRepositoryImpl) FindDataWargaByNIK(nik string) (*model.DataWarga, error) {
	query := `SELECT id, nik, nama_lengkap, tempat_lahir, tanggal_lahir, jenis_kelamin, pendidikan,
			  pekerjaan, agama, status_pernikahan, kewarganegaraan, alamat FROM datawarga WHERE nik = ?`

	row := r.db.QueryRow(query, nik)

	var warga model.DataWarga
	err := row.Scan(
		&warga.ID,
		&warga.NIK,
		&warga.NamaLengkap,
		&warga.TempatLahir,
		&warga.TanggalLahir,
		&warga.JenisKelamin,
		&warga.Pendidikan,
		&warga.Pekerjaan,
		&warga.Agama,
		&warga.StatusPernikahan,
		&warga.Kewarganegaraan,
		&warga.Alamat,
	)

	if err != nil {
		return nil, err
	}

	return &warga, nil
}

func (r *requestSuratRepositoryImpl) InsertRequestSurat(req model.RequestSuratWarga) error {
	query := `INSERT INTO requestsuratwarga (
		id_warga, jenis_surat,
		nik, nama_lengkap, tempat_lahir, tanggal_lahir, jenis_kelamin, pendidikan,
		pekerjaan, agama, status_pernikahan, kewarganegaraan, alamat,
		penghasilan, lama_tinggal, nama_usaha, jenis_usaha, alamat_usaha,
		alamat_tujuan, alasan_pindah, keperluan_pindah, tujuan_pindah,
		nama_ayah, nama_ibu, nomor_hp, tgl_kematian, penyebab_kematian, tujuan
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`	

    _, err := r.db.Exec(query,
        req.IDWarga,
        req.JenisSurat,
        
        req.NIK,
        req.NamaLengkap,
        req.TempatLahir,
        req.TanggalLahir,
        req.JenisKelamin,
        req.Pendidikan,
        req.Pekerjaan,
        req.Agama,
        req.StatusPernikahan,
        req.Kewarganegaraan,
        req.Alamat,
        
        req.Penghasilan,
        req.LamaTinggal,
        req.NamaUsaha,
        req.JenisUsaha,
        req.AlamatUsaha,
        
        req.AlamatTujuan,
        req.AlasanPindah,
        req.KeperluanPindah,
        req.TujuanPindah,
        
        req.NamaAyah,
        req.NamaIbu,
        req.NomorHP,
        req.TanggalKematian,
        req.PenyebabKematian,
        req.TujuanSurat,
    )

    return err
}