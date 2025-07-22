# Klinik Backend API

Aplikasi backend klinik modern berbasis Golang (Fiber) dengan MongoDB, Clean Architecture, dan role-based access. Mendukung fitur pasien, screening, antrian, pemeriksaan fisik, konsultasi dokter, pendampingan, pembayaran, master data obat/produk, dan manajemen kuesioner.

---

## 🚀 Fitur Utama
- Registrasi & manajemen pasien (otomatis buat akun jika screening via kasir)
- Screening kuesioner (customizable, role admin)
- Antrian screening & pemeriksaan fisik (paramedis)
- Konsultasi dokter & pendampingan (paramedis/dokter)
- Riwayat pemeriksaan fisik & screening
- Master data obat, batch, harga, produk (admin)
- Pembayaran (kasir, cash/transfer/qris, upload bukti)
- Role-based access: admin, dokter, paramedis, kasir, pasien
- Pagination di semua endpoint list

---

## 🛠️ Setup & Menjalankan
1. **Clone repo & install dependency**
2. **Jalankan MongoDB** (pastikan env `MONGO_URI` sudah benar)
3. **Jalankan server:**
   ```bash
   MONGO_URI="mongodb://localhost:27017" go run cmd/server/main.go
   ```
4. **Seeder data (opsional):**
   ```bash
   MONGO_URI="mongodb://localhost:27017" go run cmd/seeder/role_user_seeder.go
   MONGO_URI="mongodb://localhost:27017" go run cmd/seeder/screening_question_seeder.go
   ```

---

## 📚 API Endpoint List (v1)

### **Auth & User**
- `POST /api/v1/register` — Registrasi user pasien
- `POST /api/v1/login` — Login (JWT Bearer)

### **Pasien**
- `POST /api/v1/patients` — Tambah/update data pasien (kasir)
- `GET /api/v1/doctor/patients` — List pasien (dashboard dokter, pagination)

### **Screening**
- `GET /api/v1/screening/questions` — List pertanyaan screening (public)
- `POST /api/v1/screening/questions` — Tambah pertanyaan (admin only)
- `PATCH /api/v1/screening/questions/:id` — Edit pertanyaan (admin only)
- `POST /api/v1/screening/with-patient` — Screening + data pasien (kasir/pasien)
- `POST /api/v1/screening/answers` — Submit jawaban screening
- `PATCH /api/v1/screening/answers/:id` — Edit jawaban screening (paramedis)
- `GET /api/v1/screening/queue` — List antrian screening (paramedis, pagination)
- `POST /api/v1/screening/queue` — Tambah ke antrian screening

### **Pemeriksaan Fisik & Konsultasi**
- `POST /api/v1/physical-examinations` — Tambah pemeriksaan fisik (paramedis)
- `PATCH /api/v1/physical-examinations/:id` — Edit pemeriksaan fisik (paramedis/dokter)
- `GET /api/v1/physical-examinations/by-patient?patient_id=...` — Riwayat pemeriksaan fisik pasien
- `GET /api/v1/doctor/consultations` — List pasien butuh konsultasi dokter (dashboard dokter, pagination)
- `PATCH /api/v1/physical-examinations/:id/consultation-status` — Update status konsultasi dokter

### **Obat & Produk**
- `POST /api/v1/medicines` — Tambah obat (admin only)
- `PATCH /api/v1/medicines/:id` — Edit obat (admin only)
- `GET /api/v1/medicines` — List obat (pagination)
- (Produk, batch, harga, dsb: endpoint serupa, bisa dikembangkan)

### **Pembayaran**
- `POST /api/v1/payments` — Proses pembayaran (kasir)
- `GET /api/v1/payments/history` — History transaksi pembayaran (pagination)

### **Lainnya**
- `GET /api/v1/ping` — Health check

---

## 🔒 Role-based Access
- **Admin:** CRUD kuesioner, master data obat/produk, harga layanan
- **Paramedis:** Screening, pemeriksaan fisik, edit riwayat
- **Dokter:** Konsultasi, edit hasil pemeriksaan fisik, lihat daftar pasien
- **Kasir:** Input data pasien, proses pembayaran, history transaksi
- **Pasien:** Screening mandiri, lihat riwayat sendiri (bisa dikembangkan)

---

## 📦 Struktur Folder (Feature-based)
- `internal/domain/` — Entity/domain object per fitur
- `internal/repository/` — Interface & implementasi DB per fitur
- `internal/usecase/` — Business logic per fitur
- `internal/delivery/http/` — Handler Fiber per fitur
- `cmd/server/main.go` — Entry point aplikasi

---

## 📝 Catatan
- Semua endpoint list mendukung pagination: `?page=1&limit=10`
- Untuk endpoint yang butuh login, gunakan JWT Bearer di header `Authorization`
- Untuk endpoint admin-only, wajib login sebagai admin
- Untuk upload file (KTP, bukti pembayaran), gunakan `multipart/form-data`

---

## 👨‍💻 Kontribusi & Pengembangan
- Silakan kembangkan fitur produk, batch obat, harga layanan, pembayaran, dsb sesuai kebutuhan klinik Anda! 

## seed

POSTGRES_DSN=postgres://aksabumilangit:sea@localhost:5432/klinik?sslmode=disable go run cmd/seeder/screening_question_seeder_postgres.go

