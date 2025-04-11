package model

type DashboardStats struct {
	TotalPenduduk     int `json:"total_penduduk"`
	LakiLaki          int `json:"laki_laki"`
	Perempuan         int `json:"perempuan"`
	PendudukDewasa    int `json:"penduduk_dewasa"`
	PendudukAnakAnak  int `json:"penduduk_anak_anak"`
}
