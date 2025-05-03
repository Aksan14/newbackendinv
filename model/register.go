package model

import "time"

type User struct {
	Id       	string
	Email    	string
	Nikadmin 	string
	Password 	string
	NamaLengkap string
	RoleID		string
	ResetToken  string
	ResetExpiry int64  
}


type Admin struct {
    Id          int    `json:"id"`
    Email       string `json:"email"`
    NikAdmin    string `json:"nik_admin"`
    NamaLengkap string `json:"nama_lengkap"`
    RoleId      string `json:"role_id"`
    Pass        string `json:"pass"`
    Jabatan     string `json:"jabatan"` 
}


type CreateUserRequest struct {
    NikAdmin     string `json:"nikadmin" validate:"required"`
    Password     string `json:"password" validate:"required"`
    NamaLengkap  string `json:"nama_lengkap" validate:"required"`
    RoleID       int    `json:"role_id" validate:"required"`
}


type Warga struct {
	ID              int       `json:"id"`
	NIK             string    `json:"nik"`
	NamaLengkap     string    `json:"nama_lengkap"`
	Alamat          string    `json:"alamat_lengkap"`
	JenisSurat      string    `json:"jenis_surat"`
	Keterangan      string    `json:"keterangan"`
	FileUpload      string    `json:"file_upload"`
	NoHP            string    `json:"no_hp"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type DataWarga struct {
	ID					int    `json:"id"`
	NIK					string `json:"nik"`
	NamaLengkap			string `json:"nama_lengkap"`
	TempatLahir			string `json:"tempat_lahir"`
	TanggalLahir		string `json:"tanggal_lahir"`
	JenisKelamin		string `json:"jenis_kelamin"`
	Pendidikan			string `json:"pendidikan"`
	Pekerjaan			string `json:"pekerjaan"`
	Agama				string `json:"agama"`
	StatusPernikahan	string `json:"status_pernikahan"`
	Kewarganegaraan		string `json:"kewarganegaraan"`
	Alamat           	string `json:"alamat_lengkap"`
}


