package dto

type CreateUserRequest struct {
	Nikadmin	string `json:"nikadmin"`
	Email		string `json:"email"`
	Pass		string `json:"password"`
}

type LoginUserRequest struct {
	Nikadmin	string `json:"nikadmin"`
	Pass  		string `json:"password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
    Token    string `json:"token" validate:"required"`
    Password string `json:"password" validate:"required,min=6"`
}


type WargaRequest struct {
	NIK         string `json:"nik"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	JenisSurat  string `json:"jenis_surat"`
	Keterangan  string `json:"keterangan"`
	FileUpload  string `json:"file_upload"`
	NoHP        string `json:"no_hp"`
}