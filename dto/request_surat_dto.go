package dto

type RequestSuratDTO struct {
	NIK        string `json:"nik" validate:"required"`
	JenisSurat string `json:"jenis_surat" validate:"required"`
}
