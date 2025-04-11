package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"godesaapps/dto"
	"godesaapps/service"

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
	var req dto.WebsiteContentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Format data salah", http.StatusBadRequest)
		return
	}
	err = c.service.UpdateContent(&req)
	if err != nil {
		http.Error(w, "Gagal memperbarui konten", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Konten berhasil diperbarui"}`))
}
