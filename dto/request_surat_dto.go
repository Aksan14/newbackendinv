package dto

type RequestSuratDTO struct {
	NIK              string `json:"nik" validate:"required"`
	JenisSurat       string `json:"jenis_surat" validate:"required"`
	NamaLengkap      string `json:"nama_lengkap" validate:"required"`
	TempatLahir      string `json:"tempat_lahir" validate:"required"`
	TanggalLahir     string `json:"tanggal_lahir" validate:"required"` // bisa pakai time.Time jika mau parsing
	JenisKelamin     string `json:"jenis_kelamin" validate:"required"`
	Agama            string `json:"agama" validate:"required"`
	Pekerjaan        string `json:"pekerjaan" validate:"required"`
	Alamat           string `json:"alamat" validate:"required"`

	// Field khusus SKTM
	Penghasilan       string `json:"penghasilan,omitempty"`

	// Field khusus Domisili
	StatusPernikahan  string `json:"status_pernikahan,omitempty"`
	LamaTinggal       string `json:"lama_tinggal,omitempty"`

	// Field khusus Usaha
	NamaUsaha         string `json:"nama_usaha,omitempty"`
	JenisUsaha        string `json:"jenis_usaha,omitempty"`
	AlamatUsaha       string `json:"alamat_usaha,omitempty"`

	// Field khusus Pindah
	AlamatTujuan      string `json:"alamat_tujuan,omitempty"`
	AlasanPindah      string `json:"alasan_pindah,omitempty"`
	KeperluanPindah   string `json:"keperluan_pindah,omitempty"`
	TujuanPindah      string `json:"tujuan_pindah,omitempty"`
}
