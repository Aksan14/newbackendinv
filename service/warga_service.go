package service

import "godesaapps/model"

type WargaService interface {
	RegisterWarga(warga model.Warga) error
	InsertDataWarga(warga model.DataWarga) error
	GetAllWarga() ([]model.DataWarga, error)
	UpdateWarga(id int, warga model.DataWarga) error
	DeleteWarga(id int) error
}
