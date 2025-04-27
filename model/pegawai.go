package model


type Pegawai struct {
    ID              int64       `json:"id"`
    NIP             string      `json:"nip"`
    Email           string      `json:"email"`
	NamaLengkap     string      `json:"namalengkap"`
    Jabatan         string      `json:"jabatan"`
    Foto            string      `json:"foto"`
}
