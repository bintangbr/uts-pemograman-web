
# 🛍️ Aplikasi CRUD Go + Bootstrap (Produk & Kategori)

Aplikasi web CRUD sederhana menggunakan **Golang (net/http)** untuk backend dan **Bootstrap 5 + HTML + JavaScript** untuk frontend. Aplikasi ini memungkinkan manajemen **produk** dan **kategori** yang saling berelasi.

---

## 🚀 Fitur

- ✅ Tambah, tampilkan, edit, dan hapus produk
- ✅ Tambah kategori melalui modal
- ✅ Relasi produk dan kategori (one-to-many)
- ✅ Modal Bootstrap untuk input/edit data
- ✅ Pencarian produk secara langsung (live search)
- ✅ Frontend modern dan responsif dengan Bootstrap 5

---

## 🗂️ Struktur Folder

```
crud-app/
├── main.go
├── go.mod
├── database/
│   └── db.go
├── models/
│   ├── kategori.go
│   └── produk.go
├── handlers/
│   ├── kategori_handler.go
│   └── produk_handler.go
├── public/
│   └── index.html
└── static/
    ├── css/
    └── js/
```

---

## ⚙️ Cara Menjalankan

### 1. Clone repository

```bash
git clone https://github.com/username/crud-app.git
cd crud-app
```

### 2. Inisialisasi Go module

```bash
go mod init crud-app
go get github.com/go-sql-driver/mysql
```

### 3. Buat database dan tabel

```sql
CREATE DATABASE crud_app;

CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    harga FLOAT NOT NULL,
    kategori_id INT,
    FOREIGN KEY (kategori_id) REFERENCES categories(id)
);
```

### 4. Konfigurasi koneksi database

Edit file `database/db.go`:

```go
dsn := "username:password@tcp(127.0.0.1:3306)/crud_app"
```

Ganti `username` dan `password` sesuai pengaturan MySQL kamu.

### 5. Jalankan aplikasi

```bash
go run main.go
```

Buka browser ke:
```
http://localhost:8080
```

---

## 📑 Endpoint API

### Produk

| Method | Endpoint              | Deskripsi         |
|--------|------------------------|-------------------|
| GET    | `/products`            | Tampilkan semua produk |
| POST   | `/products/create`     | Tambah produk baru |
| POST   | `/products/update`     | Update produk     |
| DELETE | `/products/delete?id=` | Hapus produk      |

### Kategori

| Method | Endpoint              | Deskripsi         |
|--------|------------------------|-------------------|
| GET    | `/categories`          | Tampilkan semua kategori |
| POST   | `/categories/create`   | Tambah kategori   |

---

## 📸 Tampilan Aplikasi

> Tambahkan screenshot antarmuka aplikasi di sini jika tersedia.

---

## 🧰 Teknologi

- [Go](https://golang.org/) — Backend REST API
- [MySQL](https://www.mysql.com/) — Database
- [Bootstrap 5](https://getbootstrap.com/) — UI Framework
- HTML + JavaScript — Frontend logic

---

## 📄 Lisensi

MIT License © 2025 by [Nama Anda]
