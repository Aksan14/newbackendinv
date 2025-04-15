package repository

import "godesaapps/model"

type WargaRepository interface {
	InsertWarga(warga model.Warga) error
	InsertDataWarga(warga model.DataWarga) error
	GetAllWarga() ([]model.DataWarga, error)
	UpdateWarga(id int, warga model.DataWarga) error
	DeleteWarga(id int) error
	FindByNIK(nik string) (*model.DataWarga, error)
}

