package model

type RequestSuratWarga struct {
	IDWarga          int
	JenisSurat       string
	NIK              string
	NamaLengkap      string
	TempatLahir      string
	TanggalLahir     string
	JenisKelamin     string
	Pendidikan       string
	Pekerjaan        string
	Agama            string
	StatusPernikahan string
	Kewarganegaraan  string
	Alamat           string
	Penghasilan      float64
	LamaTinggal      *int
	NamaUsaha        string
	JenisUsaha       string
	AlamatUsaha      string
	AlamatTujuan     string
	AlasanPindah     string
	KeperluanPindah  string
	TujuanPindah     string
	NamaAyah         string
	NamaIbu          string
	NomorHP          string
	TanggalKematian  *string 
	PenyebabKematian string
	TujuanSurat      string
}