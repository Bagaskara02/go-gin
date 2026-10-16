# 🚀 Go Todo API with Gin & PostgreSQL

REST API sederhana untuk manajemen todo (create, read, update, delete) yang dibangun menggunakan **Golang**, **Gin Web Framework**, dan **PostgreSQL** dengan arsitektur layered yang rapi dan siap dikembangkan.

---

## 📋 Table of Contents

- [🚀 Go Todo API with Gin & PostgreSQL](#-go-todo-api-with-gin--postgresql)
- [📋 Table of Contents](#-table-of-contents)
- [✨ Features](#-features)
- [🛠️ Tech Stack](#️-tech-stack)
- [📁 Project Structure](#-project-structure)
- [⚙️ Installation](#️-installation)
- [🗄️ Database Setup](#️-database-setup)
- [🏃‍♂️ Running the Application](#️-running-the-application)
- [📡 API Endpoints](#-api-endpoints)
- [🧪 Testing](#-testing)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---

## ✨ Features

| Feature | Status | Description |
|---------|--------|-------------|
| ✅ Create Todo | Done | POST - Menambahkan todo baru |
| ✅ Read All Todo | Done | GET - Mengambil semua todo dengan opsi filter |
| ✅ Read One Todo | Done | GET by ID - Mengambil todo tertentu |
| ✅ Update Todo | Done | PUT - Memperbarui todo berdasarkan ID |
| ✅ Delete Todo | Done | DELETE - Menghapus todo berdasarkan ID |
| ✅ Filter by Status | Done | Query parameter `?done=true/false` |
| ✅ CORS Enabled | Done | Bisa diakses dari frontend HTML/JS |
| ✅ Validation | Done | Validasi input dengan `go-playground/validator` |
| ✅ Environment Config | Done | Konfigurasi DB melalui `.env` |

---

## 🛠️ Tech Stack

| Layer | Technology | Version |
|-------|-----------|---------|
| Language | **Go** | 1.23+ |
| Web Framework | **Gin** | v1.12.0 |
| Database Driver | **pgx / sql** | v5.11.0 |
| Validation | **go-playground/validator** | v10.30.5 |
| Environment Loader | **godotenv** | v1.5.1 |
| Database | **PostgreSQL** | Supabase / Local |

---

## 📁 Project Structure

```
golangframework/
├── .env                  # Konfigurasi environment (DATABASE_URL)
├── .gitignore            # File yang diabaikan Git
├── go.mod                # Go module dependencies
├── go.sum                # Checksum dependencies
├── main.go               # Entry point aplikasi
├── README.md             # Dokumentasi project
├── database/
│   └── database.go       # Koneksi ke PostgreSQL
├── router/
│   └── router.go         # Setup router Gin dan CORS
└── todo/
    ├── handler.go        # Logic handler (CRUD todo)
    ├── model.go          # Struct model & request body
    └── routes.go         # Definisi route endpoint todo
```

---

## ⚙️ Installation

### 1. Clone repository

```bash
git clone https://github.com/username/golangframework.git
cd golangframework
```

### 2. Install dependencies

```bash
go mod tidy
go mod download
```

### 3. Setup environment variable

Buat file `.env` di root project:

```env
DATABASE_URL=postgresql://username:password@host:port/database
```

**Contoh (Supabase):**
```env
DATABASE_URL=postgresql://postgres.xxx:password@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres
```

---

## 🗄️ Database Setup

Pastikan Anda sudah membuat tabel `todos` di PostgreSQL. Jalankan SQL berikut:

```sql
-- Buat tabel todos
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tambahkan index untuk pencarian yang lebih cepat
CREATE INDEX IF NOT EXISTS idx_todos_is_done ON todos(is_done);
CREATE INDEX IF NOT EXISTS idx_todos_created_at ON todos(created_at DESC);
```

---

## 🏃‍♂️ Running the Application

### Menjalankan server:

```bash
go run main.go
```

Server akan berjalan di: `http://localhost:8080`

### Output yang diharapkan:

```
✅ Terhubung ke database PostgreSQL
🚀 Server Gin berjalan di http://localhost:8080
```

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1/todos`

| Method | Endpoint | Deskripsi | Body / Query |
|--------|----------|-----------|--------------|
| `GET` | `/todos` | Dapatkan semua todo | `?done=true` / `?done=false` |
| `GET` | `/todos/:id` | Dapatkan todo by ID | - |
| `POST` | `/todos` | Buat todo baru | `{ "title": "...", "is_done": false }` |
| `PUT` | `/todos/:id` | Update todo by ID | `{ "title": "...", "is_done": true }` |
| `DELETE` | `/todos/:id` | Hapus todo by ID | - |

### Contoh Request & Response

**Create Todo**
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go", "is_done": false}'
```

Response:
```json
{
  "id": 1,
  "title": "Belajar Go",
  "is_done": false
}
```

**Get All Todo**
```bash
curl http://localhost:8080/api/v1/todos?done=false
```

**Update Todo**
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go Advanced", "is_done": true}'
```

**Delete Todo**
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

---

## 🧪 Testing

### Manual testing dengan Postman / Insomnia

Import collection atau coba endpoint secara langsung melalui browser atau aplikasi API client.

### Unit Testing (opsional)

Tambahkan file `*_test.go` di folder `todo/` untuk pengujian otomatis.

```bash
go test ./...
```

---

## 🤝 Contributing

Kontribusi sangat diterima! Silakan ikuti langkah berikut:

1. Fork repository ini
2. Buat branch baru: `git checkout -b feature/nama-fitur`
3. Commit perubahan: `git commit -m "feat: tambahkan fitur X"`
4. Push ke branch: `git push origin feature/nama-fitur`
5. Buat Pull Request

---

## 📄 License

Dibuat untuk tujuan pembelajaran dan latihan Go Framework. Bebaas digunakan dan dikembangkan.

---

<div align="center">

**💡 Dibuat dengan ❤️ menggunakan Golang & Gin**

</div>