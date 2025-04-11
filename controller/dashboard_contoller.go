package controller

import (
	"encoding/json"
	"godesaapps/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DashboardController interface {
	GetStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type dashboardControllerImpl struct {
	Service service.DashboardService
}

func NewDashboardController(s service.DashboardService) DashboardController {
	return &dashboardControllerImpl{Service: s}
}

func (c *dashboardControllerImpl) GetStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	stats, err := c.Service.GetStats()
	if err != nil {
		http.Error(w, "Gagal mengambil data statistik", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
