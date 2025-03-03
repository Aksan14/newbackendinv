package dto

type UserResponse struct {
	Id			string `json:"id"`
	Email		string `json:"email"`
	Nikadmin  	string `json:"nikadmin"`
	Pass		string `json:"password"`
}

type ForgotPasswordResponse struct {
	Message    string `json:"message"`
	ResetToken string `json:"reset_token,omitempty"`
}

type WargaResponse struct {
	ID          int    `json:"id"`
	NIK         string `json:"nik"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	JenisSurat  string `json:"jenis_surat"`
	Keterangan  string `json:"keterangan"`
	FileUpload  string `json:"file_upload"`
	NoHP        string `json:"no_hp"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}