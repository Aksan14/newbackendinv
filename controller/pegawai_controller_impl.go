package controller

import (
    "context"
    "encoding/json"
    "fmt"
    "godesaapps/model"
    "godesaapps/service"
    "io"
    "net/http"
    "os"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

type pegawaiControllerImpl struct {
    PegawaiService service.PegawaiService
}

func NewPegawaiController(s service.PegawaiService) PegawaiController {
    return &pegawaiControllerImpl{PegawaiService: s}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(payload)
}

func (c *pegawaiControllerImpl) CreatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    err := r.ParseMultipartForm(40 << 20) 
    if err != nil {
        http.Error(w, "Gagal parsing form", http.StatusBadRequest)
        return
    }

    nip := r.FormValue("nip")
    email := r.FormValue("email")
    jabatan := r.FormValue("jabatan")

    // Handle file upload
    file, handler, err := r.FormFile("foto")
    if err != nil {
        http.Error(w, "Foto wajib diunggah", http.StatusBadRequest)
        return
    }
    defer file.Close()

    filename := fmt.Sprintf("pegawai/%s", handler.Filename)
    dst, err := os.Create(filename)
    if err != nil {
        http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    io.Copy(dst, file)

    pegawai := model.Pegawai{
        NIP:     nip,
        Email:   email,
        Jabatan: jabatan,
        Foto:    filename,
    }

    err = c.PegawaiService.CreatePegawai(context.Background(), pegawai)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Pegawai berhasil ditambahkan"})
}


func (c *pegawaiControllerImpl) GetAllPegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    pegawaiList, err := c.PegawaiService.GetAllPegawai(context.Background())
    if err != nil {
        http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusOK, pegawaiList)
}

func (c *pegawaiControllerImpl) GetPegawaiByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idStr := ps.ByName("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "ID tidak valid", http.StatusBadRequest)
        return
    }

    pegawai, err := c.PegawaiService.GetPegawaiByID(context.Background(), id)
    if err != nil {
        http.Error(w, "Pegawai tidak ditemukan", http.StatusNotFound)
        return
    }

    respondWithJSON(w, http.StatusOK, pegawai)
}

func (c *pegawaiControllerImpl) UpdatePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Gagal parsing form", http.StatusBadRequest)
        return
    }

    idStr := r.FormValue("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "ID tidak valid", http.StatusBadRequest)
        return
    }

    nip := r.FormValue("nip")
    email := r.FormValue("email")
    jabatan := r.FormValue("jabatan")

    oldData, err := c.PegawaiService.GetPegawaiByID(context.Background(), id)
    if err != nil {
        http.Error(w, "Pegawai tidak ditemukan", http.StatusNotFound)
        return
    }

    fotoPath := oldData.Foto

    file, handler, err := r.FormFile("foto")
    if err == nil {
        defer file.Close()
        os.MkdirAll("uploads/pegawai", os.ModePerm)

        filename := fmt.Sprintf("uploads/pegawai/%s", handler.Filename)
        dst, err := os.Create(filename)
        if err != nil {
            http.Error(w, "Gagal menyimpan file foto baru", http.StatusInternalServerError)
            return
        }
        defer dst.Close()
        io.Copy(dst, file)

        fotoPath = filename
    }

    pegawai := model.Pegawai{
        ID:      id,
        NIP:     nip,
        Email:   email,
        Jabatan: jabatan,
        Foto:    fotoPath,
    }

    err = c.PegawaiService.UpdatePegawai(context.Background(), pegawai)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"message": "Pegawai berhasil diperbarui"})
}

func (c *pegawaiControllerImpl) DeletePegawai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idStr := ps.ByName("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "ID tidak valid", http.StatusBadRequest)
        return
    }

    err = c.PegawaiService.DeletePegawai(context.Background(), id)
    if err != nil {
        http.Error(w, "Gagal menghapus pegawai", http.StatusInternalServerError)
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"message": "Pegawai berhasil dihapus"})
}
