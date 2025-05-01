package dto

type RequestSuratDTO struct {
	NIK              string  `json:"nik" validate:"required"`
	JenisSurat       string  `json:"jenis_surat" validate:"required"`
	NamaLengkap      string  `json:"nama_lengkap" validate:"required"`
	TempatLahir      string  `json:"tempat_lahir" validate:"required"`
	TanggalLahir     string  `json:"tanggal_lahir" validate:"required"`
	JenisKelamin     string  `json:"jenis_kelamin" validate:"required"`
	Agama            string  `json:"agama" validate:"required"`
	Pekerjaan        string  `json:"pekerjaan" validate:"required"`
	Alamat           string  `json:"alamat" validate:"required"`
	Pendidikan       string  `json:"pendidikan,omitempty"`
	Kewarganegaraan  string  `json:"kewarganegaraan,omitempty"`

	// Ubah penghasilan ke tipe float64 atau *float64
	Penghasilan       string `json:"penghasilan,omitempty"` // Gunakan pointer jika memungkinkan NULL

	// Field khusus Domisili
	StatusPernikahan  string  `json:"status_pernikahan,omitempty"`
	LamaTinggal       string  `json:"lama_tinggal,omitempty"`

	// Field khusus Usaha
	NamaUsaha         string  `json:"nama_usaha,omitempty"`
	JenisUsaha        string  `json:"jenis_usaha,omitempty"`
	AlamatUsaha       string  `json:"alamat_usaha,omitempty"`

	// Field khusus Pindah
	AlamatTujuan      string  `json:"alamat_tujuan,omitempty"`
	AlasanPindah      string  `json:"alasan_pindah,omitempty"`
	KeperluanPindah   string  `json:"keperluan_pindah,omitempty"`
	TujuanPindah      string  `json:"tujuan_pindah,omitempty"`

	// Field khusus Kelahiran
	NamaAyah          string  `json:"nama_ayah,omitempty"`
	NamaIbu           string  `json:"nama_ibu,omitempty"`

	// Field khusus Kematian
	TanggalKematian   string  `json:"tanggal_kematian,omitempty"`
	PenyebabKematian  string  `json:"penyebab_kematian,omitempty"`

	// Field umum
	NomorHP           string  `json:"nomor_hp,omitempty"`
	TujuanSurat       string  `json:"tujuan_surat,omitempty"`
}
