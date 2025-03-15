# INOVASI DESA COCONUT 2025 MAIN APP

API untuk Log-in Log-out Admin desa untuk mengakses Apk Bendahara dan Sekretaris.

## 🛠 Tech Stack

**Framework & Libraries:**
- Go Validator v10 - Library untuk validasi request secara otomatis.

- Godotenv v1.5.1 - Mengelola environment variables dari file .env.

- MySQL Driver v1.8.1 - Driver database MySQL untuk Go.

- Google UUID v1.6.0 - Library untuk menghasilkan UUID.

- JWT (dgrijalva/jwt-go) v3.2.0+incompatible - Library untuk menangani autentikasi menggunakan JSON Web Token (JWT).

- httprouter v1.3.0 - Router HTTP yang ringan dan cepat.

- golang.org/x/crypto v0.32.0 - Paket kriptografi untuk keamanan aplikasi.

- filippo.io/edwards25519 v1.1.0 - Library untuk implementasi algoritma Edwards25519.

DB_HOST=localhost   
# Alamat host database, dalam hal ini database berjalan di mesin lokal.  

DB_PORT=3306     
# Port yang digunakan oleh database MySQL (default: 3306).  

DB_USER=root  
# Nama pengguna yang digunakan untuk mengakses database.  

DB_PASSWORD=haha123  
# Kata sandi untuk autentikasi pengguna database.  

DB_NAME=  
# Nama database yang digunakan oleh aplikasi.  

APP_PORT=
# Port tempat aplikasi web berbasis Go berjalan.  

SMTP_EMAIL= 
# Alamat email yang digunakan untuk mengirim email melalui SMTP.  

SMTP_PASSWORD=  
# Kata sandi atau kode autentikasi SMTP untuk mengirim email.  

SMTP_HOST=smtp.gmail.com  
# Alamat server SMTP untuk mengirim email melalui Gmail.  

SMTP_PORT=587  
# Port yang digunakan untuk komunikasi dengan server SMTP Gmail (587 untuk TLS).  

ALLOWED_ORIGIN=http://localhost:3000, * untuk mengizinkan tanpa terkecuali
# Domain yang diizinkan untuk mengakses API (digunakan dalam konfigurasi CORS).  

BASE_URL=localhost:3000  
# URL dasar aplikasi frontend yang digunakan untuk membentuk link dalam email.  


**Database:**
- MySQL

**Tools:**
- go1.23.2
- Postman (testing)

## 🚀 Fitur
- Log-in, Log-out, Register admin desa
- Manajemen respons terstruktur
- Validasi request
- Transactional database operations
- Error handling terstruktur

## 📁 Project Structure
.
├── config/
│   └── connection.go        # Konfigurasi database
├── controller/
│   ├── user_controller.go          # Controller untuk user
│   ├── user_controller_impl.go     # Implementasi controller user
│   ├── warga_controller.go         # Controller untuk warga
│   └── warga_controller_impl.go    # Implementasi controller warga
├── dto/
│   ├── response_list.go    # Daftar response DTO
│   ├── user_request.go     # DTO untuk request user
│   ├── user_response.go    # DTO untuk response user
│   ├── warga_request.go    # DTO untuk request warga
│   └── warga_response.go   # DTO untuk response warga
├── filewarga/
│   ├── logo.pdf                         # File terkait warga (misal: dokumen atau logo)
│   └── Screenshot-2025-02-23.png         # Screenshot terkait proyek
├── model/
│   └── register.go        # Model database untuk registrasi
├── repository/
│   ├── user_repository.go          # Interface repository user
│   ├── user_repository_impl.go     # Implementasi repository user
│   ├── warga_repository.go         # Interface repository warga
│   └── warga_repository_impl.go    # Implementasi repository warga
├── service/
│   ├── user_service.go          # Interface service user
│   ├── user_service_impl.go     # Implementasi service user
│   ├── warga_service.go         # Interface service warga
│   └── warga_service_impl.go    # Implementasi service warga
├── util/
│   ├── error.go        # Handler error
│   ├── json.go         # Helper JSON
│   ├── mailer.go       # Utilitas pengiriman email
│   ├── model.go        # Utilitas model
│   └── transaction.go  # Utilitas transaksi database
├── .env               # Konfigurasi environment
├── .gitignore         # File untuk mengabaikan file tertentu di Git
├── go.mod             # Module Go
├── go.sum             # Dependency checksum
├── main.go            # Entry point aplikasi
└── README.md          # Dokumentasi proyek


## 🏃 Menjalankan Server
go run main.go
<!-- Server akan berjalan di http://localhost:8080. -->

## 📚 API Endpoints

# Method	    # Endpoint	                      # Deskripsi
<!-- Admin -->
POST	          /api/user/sign-up               Buat Admin baru yang hanya dapat diakses jika sudah log-in
POST	          /api/user/login       	        Untuk Login Sebagai admin
GET             /api/user/me                    Untuk Mengakses data Admin yang Login saat ini
POST            /api/user/forgot-password       Untuk Req Reset Password saat Admin lupa sandi
POST            /api/user/reset-password        Untuk Reset password
GET             /api/user/dashboard-bendahara   bagi hak akses apk didalamnya dengan role bendahara / ROLE001
GET             /api/user/dashboard-sekretaris  bagi hak akses apk didalamnya dengan role sekretaris / ROLE002

<!-- Warga -->
POST            /api/warga/register           Untuk Form warga untuk req ke sekretaris
