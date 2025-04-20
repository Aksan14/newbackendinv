package service

import (
	"godesaapps/dto"
	"godesaapps/model"
)

type RequestSuratService interface {
	RequestSurat(input dto.RequestSuratDTO) error
	FindByNik(nik string) (*model.DataWarga, error)
}
