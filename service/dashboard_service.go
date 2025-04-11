package service

import "godesaapps/repository"

type DashboardStats struct {
	Total     int `json:"total"`
	Laki      int `json:"laki"`
	Perempuan int `json:"perempuan"`
}

type DashboardService interface {
	GetStats() (DashboardStats, error)
}

type dashboardServiceImpl struct {
	Repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardServiceImpl{Repo: repo}
}

func (s *dashboardServiceImpl) GetStats() (DashboardStats, error) {
	total, laki, perempuan, err := s.Repo.GetWargaStats()
	if err != nil {
		return DashboardStats{}, err
	}

	return DashboardStats{
		Total:     total,
		Laki:      laki,
		Perempuan: perempuan,
	}, nil
}
