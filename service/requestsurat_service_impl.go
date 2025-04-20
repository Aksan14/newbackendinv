package service

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
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


// Service function untuk menyalin data warga ke request surat
func (s *requestSuratServiceImpl) RequestSurat(input dto.RequestSuratDTO) error {
	// Mendapatkan data warga berdasarkan NIK
	warga, err := s.repo.FindDataWargaByNIK(input.NIK)
	if err != nil {
		return fmt.Errorf("data warga tidak ditemukan: %v", err)
	}

	// Membuat objek request surat berdasarkan data warga
	request := model.RequestSuratWarga{
		IDWarga:          warga.ID,               // Mengambil ID warga
		JenisSurat:       input.JenisSurat,       // Jenis surat diinput dari request                 // Tanggal selesai, bisa diset nanti
		NIK:              warga.NIK,              // NIK warga yang diambil dari data warga
		NamaLengkap:      warga.NamaLengkap,      // Nama lengkap warga yang diambil dari data warga
		TempatLahir:      warga.TempatLahir,      // Tempat lahir warga
		TanggalLahir:     warga.TanggalLahir,     // Tanggal lahir warga
		JenisKelamin:     warga.JenisKelamin,     // Jenis kelamin warga
		Pendidikan:       warga.Pendidikan,       // Pendidikan warga
		Pekerjaan:        warga.Pekerjaan,        // Pekerjaan warga
		Agama:            warga.Agama,            // Agama warga
		StatusPernikahan: warga.StatusPernikahan, // Status pernikahan warga
		Kewarganegaraan:  warga.Kewarganegaraan,  // Kewarganegaraan warga
		Alamat:           warga.Alamat,           // Alamat warga
	}

	// Menyimpan data request surat ke database
	err = s.repo.InsertRequestSurat(request)
	if err != nil {
		return fmt.Errorf("gagal menyimpan permintaan surat: %v", err)
	}

	return nil
}




