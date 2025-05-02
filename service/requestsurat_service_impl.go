package service

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
	"strconv"
	"time"
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

	var lamaTinggal *int
	if input.JenisSurat == "Domisili" {
		if input.LamaTinggal == "" {
			return fmt.Errorf("lama tinggal harus diisi untuk surat domisili")
		}
		lamaTinggalInt, err := strconv.Atoi(input.LamaTinggal)
		if err != nil {
			return fmt.Errorf("gagal mengonversi LamaTinggal ke int: %v", err)
		}
		if lamaTinggalInt < 6 {
			return fmt.Errorf("lama tinggal minimal 6 bulan untuk surat domisili")
		}
		lamaTinggal = &lamaTinggalInt
	} else {
		if input.LamaTinggal != "" {
			lamaTinggalInt, err := strconv.Atoi(input.LamaTinggal)
			if err != nil {
				return fmt.Errorf("gagal mengonversi LamaTinggal ke int: %v", err)
			}
			lamaTinggal = &lamaTinggalInt
		} else {
			lamaTinggal = nil		}
	}

	var penghasilan float64
	if input.Penghasilan != "" {
		p, err := strconv.ParseFloat(input.Penghasilan, 64)
		if err != nil {
			return fmt.Errorf("gagal mengonversi Penghasilan ke float64: %v", err)
		}
		penghasilan = p
	} else {
		penghasilan = 0.0
	}

	var tanggalKematian *string
	if input.JenisSurat == "Kematian" {
		if input.TanggalKematian == "" {
			return fmt.Errorf("tanggal kematian harus diisi untuk surat kematian")
		}
		_, err := time.Parse("2006-01-02", input.TanggalKematian)
		if err != nil {
			return fmt.Errorf("format tanggal kematian tidak valid: %v", err)
		}
		tanggalKematian = &input.TanggalKematian
	} else {
		tanggalKematian = nil
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
		Penghasilan:      penghasilan,
		LamaTinggal:      lamaTinggal,
		NamaUsaha:        input.NamaUsaha,
		JenisUsaha:       input.JenisUsaha,
		AlamatUsaha:      input.AlamatUsaha,
		AlamatTujuan:     input.AlamatTujuan,
		AlasanPindah:     input.AlasanPindah,
		KeperluanPindah:  input.KeperluanPindah,
		TujuanPindah:     input.TujuanPindah,
		NamaAyah:         input.NamaAyah,
		NamaIbu:          input.NamaIbu,
		NomorHP:          input.NomorHP,
		TanggalKematian:  tanggalKematian,
		PenyebabKematian: input.PenyebabKematian,
		TujuanSurat:      input.TujuanSurat,
	}

	err = s.repo.InsertRequestSurat(request)
	if err != nil {
		return fmt.Errorf("gagal menyimpan permintaan surat: %v", err)
	}

	return nil
}
