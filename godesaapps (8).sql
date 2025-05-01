-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 01 Bulan Mei 2025 pada 12.39
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `godesaapps`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `admin`
--

CREATE TABLE `admin` (
  `id` char(36) NOT NULL,
  `email` varchar(100) NOT NULL,
  `nikadmin` varchar(100) NOT NULL,
  `namalengkap` varchar(255) NOT NULL,
  `role_id` varchar(100) DEFAULT NULL,
  `pass` varchar(100) NOT NULL,
  `reset_token` text DEFAULT NULL,
  `reset_expiry` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `admin`
--

INSERT INTO `admin` (`id`, `email`, `nikadmin`, `namalengkap`, `role_id`, `pass`, `reset_token`, `reset_expiry`) VALUES
('32', 'muhammadaksan263@gmail.com', '11111111', 'Sekretaris Desa', 'ROLE002', '$2a$10$cMniYYC87DpierOeHmf/9eXVvoVQ7tFiDYFR8cRujkQctr4pJPzZe', NULL, NULL),
('34', 'Kepaladesa@gmail.com', '22222222', 'Kepala Desa', 'ROLE000', '$2a$10$VVa17B4dQ4al9N9aHlISKuAfJ/vtBvI3sP6CGCehECa6vc8PDaWyi', NULL, NULL),
('35', 'Bendaharadesa@gmaial.com', '33333333', 'Bendahara Desa', 'ROLE001', '$2a$10$uX4ncNodeqnGY.haLLeA7uP8/pEqJnfM1QsqJngPWNeI8EX7cD9uG', NULL, NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `datawarga`
--

CREATE TABLE `datawarga` (
  `id` int(11) NOT NULL,
  `nik` varchar(20) NOT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `tempat_lahir` varchar(50) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `jenis_kelamin` enum('Laki-laki','Perempuan') DEFAULT NULL,
  `pendidikan` varchar(50) DEFAULT NULL,
  `pekerjaan` varchar(50) DEFAULT NULL,
  `agama` text DEFAULT NULL,
  `status_pernikahan` text DEFAULT NULL,
  `kewarganegaraan` text DEFAULT NULL,
  `alamat` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `datawarga`
--

INSERT INTO `datawarga` (`id`, `nik`, `nama_lengkap`, `tempat_lahir`, `tanggal_lahir`, `jenis_kelamin`, `pendidikan`, `pekerjaan`, `agama`, `status_pernikahan`, `kewarganegaraan`, `alamat`) VALUES
(10, '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 'SMA', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf'),
(11, '7371131908050005', 'iKAN', 'p', '2005-04-04', 'Perempuan', 'SMA', 'kuli', 'Konghucu', 'Cerai', 'WNA', 'kjnknda'),
(16, '7371131908050008', 'aksan', 'Makassar', '7777-07-07', 'Perempuan', 'S1', 'Programmer', 'Islam', 'Belum Menikah', 'WNA', ''),
(17, '7371131908055555', 'Budi Santoso', 'Makassar', '2005-08-19', 'Laki-laki', 'S1', 'CEO', 'Islam', 'Belum Menikah', 'WNI', '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `pegawai`
--

CREATE TABLE `pegawai` (
  `id` int(11) NOT NULL,
  `nip` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `jabatan` varchar(100) NOT NULL,
  `foto` varchar(255) NOT NULL,
  `namalengkap` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `pegawai`
--

INSERT INTO `pegawai` (`id`, `nip`, `email`, `jabatan`, `foto`, `namalengkap`) VALUES
(32, '11111111', 'muhammadaksan263@gmail.com', 'Sekretaris Desa', 'pegawai/WhatsApp Image 2025-04-27 at 19.56.25.jpeg', 'Sekretaris Desa'),
(34, '22222222', 'Kepaladesa@gmail.com', 'Kepala Desa', 'pegawai/Screenshot from 2025-03-06 02-13-15.png', 'Kepala Desa'),
(35, '33333333', 'Bendaharadesa@gmaial.com', 'Bendahara Desa', 'pegawai/Screenshot from 2025-04-28 23-33-42.png', 'Bendahara Desa');

-- --------------------------------------------------------

--
-- Struktur dari tabel `requestsuratwarga`
--

CREATE TABLE `requestsuratwarga` (
  `id_pengajuan` int(11) NOT NULL,
  `id_warga` int(11) NOT NULL,
  `jenis_surat` varchar(100) DEFAULT NULL,
  `nik` varchar(20) DEFAULT NULL,
  `nama_lengkap` varchar(100) DEFAULT NULL,
  `tempat_lahir` varchar(100) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `jenis_kelamin` varchar(20) DEFAULT NULL,
  `pendidikan` varchar(50) DEFAULT NULL,
  `pekerjaan` varchar(100) DEFAULT NULL,
  `agama` varchar(50) DEFAULT NULL,
  `status_pernikahan` varchar(50) DEFAULT NULL,
  `kewarganegaraan` varchar(50) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `penghasilan` double DEFAULT NULL,
  `lama_tinggal` int(11) DEFAULT NULL,
  `nama_usaha` varchar(100) DEFAULT NULL,
  `jenis_usaha` varchar(100) DEFAULT NULL,
  `alamat_usaha` text DEFAULT NULL,
  `alamat_tujuan` text DEFAULT NULL,
  `alasan_pindah` text DEFAULT NULL,
  `keperluan_pindah` text DEFAULT NULL,
  `tujuan_pindah` text DEFAULT NULL,
  `nama_ayah` varchar(100) DEFAULT NULL,
  `nama_ibu` varchar(100) DEFAULT NULL,
  `nomor_hp` varchar(15) DEFAULT NULL,
  `tgl_kematian` date DEFAULT NULL,
  `penyebab_kematian` varchar(100) DEFAULT NULL,
  `tujuan` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `requestsuratwarga`
--

INSERT INTO `requestsuratwarga` (`id_pengajuan`, `id_warga`, `jenis_surat`, `nik`, `nama_lengkap`, `tempat_lahir`, `tanggal_lahir`, `jenis_kelamin`, `pendidikan`, `pekerjaan`, `agama`, `status_pernikahan`, `kewarganegaraan`, `alamat`, `penghasilan`, `lama_tinggal`, `nama_usaha`, `jenis_usaha`, `alamat_usaha`, `alamat_tujuan`, `alasan_pindah`, `keperluan_pindah`, `tujuan_pindah`, `nama_ayah`, `nama_ibu`, `nomor_hp`, `tgl_kematian`, `penyebab_kematian`, `tujuan`) VALUES
(95, 10, 'SKTM', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 900, NULL, '', '', '', '', '', '', '', '', '', '', NULL, '', ''),
(96, 10, 'Domisili', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, 9, '', '', '', '', '', '', '', '', '', '', NULL, '', ''),
(97, 10, 'Usaha', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, NULL, 'usaha ', 'Perdagangan', '323232', '', '', '', '', '', '', '', NULL, '', ''),
(98, 10, 'Pindah', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, NULL, '', '', '', 'asasd', 'Pendidikan', 'dsada', 'dsad', '', '', '', NULL, '', ''),
(99, 10, 'Pengantar', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, NULL, '', '', '', '', '', '', '', '', '', '', NULL, '', ''),
(100, 10, 'Kelahiran', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, NULL, '', '', '', '', '', '', '', '67', 'jhnbsj', 'djasi', NULL, '', 'dasdasd'),
(101, 10, 'Kematian', '7371131908050004', 'Muhammad Aksan', 'Makassar', '2005-08-19', 'Laki-laki', 's1', 'Main', 'Islam', 'Belum Menikah', 'WNI', 'jfjaeelfjailf', 0, NULL, '', '', '', '', '', '', '', '', '', '089', '2024-02-21', 'ke', 'kalmka');

-- --------------------------------------------------------

--
-- Struktur dari tabel `role_admin`
--

CREATE TABLE `role_admin` (
  `id` varchar(100) NOT NULL,
  `name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `role_admin`
--

INSERT INTO `role_admin` (`id`, `name`) VALUES
('ROLE000', 'Admin'),
('ROLE001', 'Bendahara'),
('ROLE002', 'Sekretaris');

-- --------------------------------------------------------

--
-- Struktur dari tabel `website_content`
--

CREATE TABLE `website_content` (
  `id` int(11) NOT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `title` text DEFAULT NULL,
  `description` text DEFAULT NULL,
  `address` text DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `website_content`
--

INSERT INTO `website_content` (`id`, `logo`, `title`, `description`, `address`, `email`, `phone`) VALUES
(1, 'kontenwebsite/Logo_Universitas_Muhammadiyah_Makassar_Resmi.jpg', 'Desa Karunrung', 'Kec. Tarowang Kab. Jeneponto', 'Desa Bonto Ujung', 'hellococonut@coconut.or.id', '08999999999');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nikadmin` (`nikadmin`),
  ADD KEY `fk_role_admin` (`role_id`);

--
-- Indeks untuk tabel `datawarga`
--
ALTER TABLE `datawarga`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nik` (`nik`);

--
-- Indeks untuk tabel `pegawai`
--
ALTER TABLE `pegawai`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `requestsuratwarga`
--
ALTER TABLE `requestsuratwarga`
  ADD PRIMARY KEY (`id_pengajuan`),
  ADD KEY `id_warga` (`id_warga`);

--
-- Indeks untuk tabel `role_admin`
--
ALTER TABLE `role_admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indeks untuk tabel `website_content`
--
ALTER TABLE `website_content`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `datawarga`
--
ALTER TABLE `datawarga`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT untuk tabel `pegawai`
--
ALTER TABLE `pegawai`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT untuk tabel `requestsuratwarga`
--
ALTER TABLE `requestsuratwarga`
  MODIFY `id_pengajuan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT untuk tabel `website_content`
--
ALTER TABLE `website_content`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `admin`
--
ALTER TABLE `admin`
  ADD CONSTRAINT `fk_role_admin` FOREIGN KEY (`role_id`) REFERENCES `role_admin` (`id`);

--
-- Ketidakleluasaan untuk tabel `requestsuratwarga`
--
ALTER TABLE `requestsuratwarga`
  ADD CONSTRAINT `requestsuratwarga_ibfk_1` FOREIGN KEY (`id_warga`) REFERENCES `datawarga` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
