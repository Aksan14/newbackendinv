package controller

import (
	"encoding/json"
	"fmt"
	"godesaapps/dto"
	"godesaapps/service"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type WebsiteContentController struct {
	service service.WebsiteContentService
}

func NewWebsiteContentController(service service.WebsiteContentService) *WebsiteContentController {
	return &WebsiteContentController{service}
}

func (c *WebsiteContentController) GetContent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content, err := c.service.GetContent()
	if err != nil {
		log.Println("Gagal ambil konten:", err)
		http.Error(w, "Gagal mengambil konten", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}

func (c *WebsiteContentController) UpdateContent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseMultipartForm(10 << 20) // Maks 10MB
	if err != nil {
		http.Error(w, "Gagal parsing form", http.StatusBadRequest)
		return
	}

	var req dto.WebsiteContentRequest
	req.Title = r.FormValue("title")
	req.Description = r.FormValue("description")
	req.Address = r.FormValue("address")
	req.Email = r.FormValue("email")
	req.Phone = r.FormValue("phone")

	logoPath := r.FormValue("logo")

	file, handler, err := r.FormFile("logo")
	if err == nil {
		defer file.Close()

		uploadDir := "kontenwebsite"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.MkdirAll(uploadDir, os.ModePerm)
		}

		filePath := fmt.Sprintf("%s/%s", uploadDir, handler.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Gagal menyalin file", http.StatusInternalServerError)
			return
		}

		req.Logo = filePath
	} else {
		req.Logo = logoPath
	}

	err = c.service.UpdateContent(&req)
	if err != nil {
		http.Error(w, "Gagal memperbarui konten", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Konten berhasil diperbarui",
	})
}
