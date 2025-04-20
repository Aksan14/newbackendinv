package dto

import "time"

type WargaRequest struct {
	NIK              string     `json:"nik"`
	JenisSurat      string     `json:"jenis_surat"`
	Keterangan      string     `json:"keterangan"`
	NoHP            string     `json:"no_hp"`              
	Alamat          string     `json:"alamat"`            
	FileUpload      string     `json:"file_upload"`       
	TanggalSelesai  *time.Time `json:"tanggal_selesai"`

	// Data Warga
	NamaLengkap     string     `json:"nama_lengkap"`
	TempatLahir     string     `json:"tempat_lahir"`
	TanggalLahir    string     `json:"tanggal_lahir"`
	JenisKelamin    string     `json:"jenis_kelamin"`
	Pendidikan      string     `json:"pendidikan"`
	Pekerjaan       string     `json:"pekerjaan"`
	Agama           string     `json:"agama"`
	StatusPernikahan string    `json:"status_pernikahan"`
	Kewarganegaraan string     `json:"kewarganegaraan"`

	// Data Tambahan
	Penghasilan     float64    `json:"penghasilan,omitempty"`    
	LamaTinggal     int        `json:"lama_tinggal,omitempty"`  
	NamaUsaha       string     `json:"nama_usaha,omitempty"`
	JenisUsaha      string     `json:"jenis_usaha,omitempty"`
	AlamatUsaha     string     `json:"alamat_usaha,omitempty"`
	AlamatTujuan    string     `json:"alamat_tujuan,omitempty"`
	AlasanPindah    string     `json:"alasan_pindah,omitempty"`
	KeperluanPindah string     `json:"keperluan_pindah,omitempty"`
	TujuanPindah    string     `json:"tujuan_pindah,omitempty"`
}
