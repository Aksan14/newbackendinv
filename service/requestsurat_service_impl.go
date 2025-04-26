package service

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
	"strconv"
)

type requestSuratServiceImpl struct {
	repo repository.RequestSuratRepository
}

func NewRequestSuratService(repo repository.RequestSuratRepository) RequestSuratService {
	return &requestSuratServiceImpl{repo}
}

func (s *requestSuratServiceImpl) FindByNik(nik string) (*model.DataWarga, error) {
	warga, err := s.repo.FindByNik(nik)
	if err != nil {
		return nil, fmt.Errorf("gagal mencari warga dengan NIK %s: %v", nik, err)
	}

	if warga == nil {
		return nil, fmt.Errorf("warga dengan NIK %s tidak ditemukan", nik)
	}

	return warga, nil
}

func (s *requestSuratServiceImpl) RequestSurat(input dto.RequestSuratDTO) error {
	warga, err := s.repo.FindDataWargaByNIK(input.NIK)
	if err != nil {
		return fmt.Errorf("data warga tidak ditemukan: %v", err)
	}

	lamaTinggalInt, err := strconv.Atoi(input.LamaTinggal)
	if err != nil {
		return fmt.Errorf("gagal mengonversi LamaTinggal ke int: %v", err)
	}

	penghasilanFloat, err := strconv.ParseFloat(input.Penghasilan, 64)
	if err != nil {
		return fmt.Errorf("gagal mengonversi Penghasilan ke float64: %v", err)
	}

	request := model.RequestSuratWarga{
		IDWarga:          warga.ID,
		JenisSurat:       input.JenisSurat,
		NIK:              warga.NIK,
		NamaLengkap:      warga.NamaLengkap,
		TempatLahir:      warga.TempatLahir,
		TanggalLahir:     warga.TanggalLahir,
		JenisKelamin:     warga.JenisKelamin,
		Pendidikan:       warga.Pendidikan,
		Pekerjaan:        warga.Pekerjaan,
		Agama:            warga.Agama,
		StatusPernikahan: warga.StatusPernikahan,
		Kewarganegaraan:  warga.Kewarganegaraan,
		Alamat:           warga.Alamat,
		Penghasilan:      penghasilanFloat,  
		LamaTinggal:      lamaTinggalInt,   
		NamaUsaha:        input.NamaUsaha,
		JenisUsaha:       input.JenisUsaha,
		AlamatUsaha:      input.AlamatUsaha,
		AlamatTujuan:     input.AlamatTujuan,
		AlasanPindah:     input.AlasanPindah,
		KeperluanPindah:  input.KeperluanPindah,
		TujuanPindah:     input.TujuanPindah,
	}

	err = s.repo.InsertRequestSurat(request)
	if err != nil {
		return fmt.Errorf("gagal menyimpan permintaan surat: %v", err)
	}

	return nil
}
