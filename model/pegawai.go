package model


type Pegawai struct {
    ID      int64  `json:"id"`
    NIP     string `json:"nip"`
    Email   string `json:"email"`
    Jabatan string `json:"jabatan"`
    Foto    string `json:"foto"`
}
