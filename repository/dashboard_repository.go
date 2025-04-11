package repository

import "database/sql"

type DashboardRepository interface {
	GetWargaStats() (total int, laki int, perempuan int, err error)
}

type dashboardRepositoryImpl struct {
	DB *sql.DB
}

func NewDashboardRepository(db *sql.DB) DashboardRepository {
	return &dashboardRepositoryImpl{DB: db}
}

func (r *dashboardRepositoryImpl) GetWargaStats() (int, int, int, error) {
	var total, laki, perempuan int

	query := `SELECT 
		COUNT(*) as total, 
		SUM(CASE WHEN jenis_kelamin = 'Laki-laki' THEN 1 ELSE 0 END) as laki,
		SUM(CASE WHEN jenis_kelamin = 'Perempuan' THEN 1 ELSE 0 END) as perempuan 
		FROM datawarga`

	err := r.DB.QueryRow(query).Scan(&total, &laki, &perempuan)
	if err != nil {
		return 0, 0, 0, err
	}

	return total, laki, perempuan, nil
}
