# Backend Article Management API (Sharing Vision Test)

RESTful API backend untuk Article Management System yang dibangun menggunakan **Golang**, **Gin Web Framework**, dan **GORM** dengan database **MySQL**.

---

## Tech Stack & Dependencies

- **Language:** Go (v1.20+)
- **HTTP Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database ORM:** [GORM](https://gorm.io/) with MySQL Driver
- **Validation:** [go-playground/validator/v10](https://github.com/go-playground/validator)
- **Environment Management:** [godotenv](https://github.com/joho/godotenv)
- **CORS Support:** [gin-contrib/cors](https://github.com/gin-contrib/cors)

---

## Database Schema

Aplikasi secara otomatis melakukan migrasi skema tabel (`AutoMigrate`) ke database saat dijalankan.

**Nama Tabel:** `posts`

| Field | Tipe Data | Deskripsi |
| :--- | :--- | :--- |
| `id` | `INT AUTO_INCREMENT` | Primary Key |
| `title` | `VARCHAR(200)` | Judul artikel |
| `content` | `TEXT` | Isi konten artikel |
| `category` | `VARCHAR(100)` | Kategori artikel |
| `created_date` | `TIMESTAMP` | Waktu pembuatan artikel |
| `updated_date` | `TIMESTAMP` | Waktu pembaruan artikel |
| `status` | `VARCHAR(100)` | Status artikel (`publish`, `draft`, atau `thrash`) |

---

## Prasyarat (Prerequisites)

Sebelum menjalankan aplikasi, pastikan Anda telah menginstal:
1. [Go](https://go.dev/dl/) (versi 1.20 atau yang lebih baru)
2. Database [MySQL](https://www.mysql.com/) lokal atau cloud (misalnya Aiven MySQL / Docker)
3. Database `article` sudah dibuat di MySQL Anda:
   ```sql
   CREATE DATABASE article;
   ```

---

## Konfigurasi Environment (`.env`)

Duplikasi file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Sesuaikan nilai variabel lingkungan dengan konfigurasi database MySQL Anda:

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=article
SERVER_PORT=8080
```

---

## Cara Menjalankan Aplikasi (How to Run)

### 1. Download Dependencies
Pastikan semua dependensi Go terunduh:
```bash
go mod tidy
```

### 2. Jalankan Server (Development Mode)
Jalankan server langsung menggunakan perintah:
```bash
go run cmd/api/main.go
```

Setelah dijalankan, Anda akan melihat log:
```
Database connection successfully established.
Database migration completed.
[GIN-debug] Listening and serving HTTP on :8080
```

### 3. Build Binary (Production Mode - Opsional)
Jika ingin meng-compile binary aplikasi:
```bash
# Windows
go build -o api.exe cmd/api/main.go
./api.exe

# Linux / macOS
go build -o api cmd/api/main.go
./api
```

Server akan berjalan secara default pada port `http://localhost:8080`.

---

## API Endpoints

Semua endpoint dikelompokkan dalam route group `/article`:

| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| **POST** | `/article/` | Membuat artikel baru |
| **GET** | `/article/:limit/:offset` | Mengambil daftar artikel dengan pagination |
| **GET** | `/article/:id` | Mengambil detail artikel berdasarkan ID |
| **PUT** | `/article/:id` | Memperbarui artikel berdasarkan ID |
| **DELETE**| `/article/:id` | Soft delete artikel (mengubah status menjadi `thrash`) |

---

### Request & Response Examples

#### 1. Create Article (`POST /article/`)

**Request Body:**
```json
{
  "title": "Judul Artikel Minimal 20 Karakter Di Sini",
  "content": "Konten artikel ini harus memiliki panjang minimal 200 karakter agar dapat melewati validasi strict yang telah ditentukan. Pastikan teks ini cukup panjang dan representatif sesuai dengan kebutuhan pengujian backend sharing vision.",
  "category": "Technology",
  "status": "publish"
}
```

**Aturan Validasi Input:**
- `title`: Wajib diisi (`required`), minimal 20 karakter (`min=20`).
- `content`: Wajib diisi (`required`), minimal 200 karakter (`min=200`).
- `category`: Wajib diisi (`required`), minimal 3 karakter (`min=3`).
- `status`: Wajib diisi (`required`), salah satu dari: `publish`, `draft`, atau `thrash`.

**Response Sukses (`200 OK`):**
```json
{}
```

**Response Validasi Gagal (`400 Bad Request`):**
```json
{
  "errors": [
    {
      "field": "title",
      "message": "Title must be at least 20 characters"
    },
    {
      "field": "content",
      "message": "Content must be at least 200 characters"
    }
  ]
}
```

---

#### 2. Get Articles with Pagination (`GET /article/:limit/:offset`)
Contoh: `GET /article/10/0` (limit 10 data, mulai offset 0)

**Response (`200 OK`):**
```json
[
  {
    "id": 1,
    "title": "Judul Artikel Minimal 20 Karakter Di Sini",
    "content": "Konten artikel...",
    "category": "Technology",
    "created_date": "2026-10-15T10:00:00Z",
    "updated_date": "2026-10-15T10:00:00Z",
    "status": "publish"
  }
]
```

---

#### 3. Get Article by ID (`GET /article/:id`)
Contoh: `GET /article/1`

**Response Sukses (`200 OK`):**
```json
{
  "id": 1,
  "title": "Judul Artikel Minimal 20 Karakter Di Sini",
  "content": "Konten artikel...",
  "category": "Technology",
  "created_date": "2026-10-15T10:00:00Z",
  "updated_date": "2026-10-15T10:00:00Z",
  "status": "publish"
}
```

**Response Tidak Ditemukan (`404 Not Found`):**
```json
{
  "errors": [
    {
      "field": "id",
      "message": "Article not found"
    }
  ]
}
```

---

#### 4. Update Article (`PUT /article/:id`)
Contoh: `PUT /article/1`

**Request Body:** Sama dengan request body `POST /article/`.

**Response Sukses (`200 OK`):**
```json
{}
```

---

#### 5. Delete Article (`DELETE /article/:id`)
Contoh: `DELETE /article/1`

Menghapus artikel secara soft delete (status artikel diubah menjadi `thrash`).

**Response Sukses (`200 OK`):**
```json
{
  "message": "Article moved to thrash"
}
```

---
