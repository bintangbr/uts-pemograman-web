
# 🛍️ Aplikasi CRUD Go + Bootstrap (Produk & Kategori) (UTS Pemograman Web)

Aplikasi web CRUD Toko sederhana menggunakan **Golang (net/http)** untuk backend dan **Bootstrap 5 + HTML + JavaScript** untuk frontend. Aplikasi ini memungkinkan manajemen **produk** dan **kategori** yang saling berelasi.

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
    └── index.html
```

---

## ⚙️ Cara Menjalankan

### 1. Clone repository

```bash
git clone https://github.com/bintangbr/uts-pemograman-web.git
cd uts-pemograman-web
```

### 2. Inisialisasi Go module

```bash
go mod init uts-pemograman-web
go get github.com/go-sql-driver/mysql
```

### 3. Buat database dan tabel

```sql
CREATE DATABASE uts-pemograman-web;

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
dsn := "root@tcp(127.0.0.1:3306)/uts-pemograman-web" (usyng passwd no)
```

Ganti `username` dan `password` sesuai pengaturan MySQL.

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

MIT License © 2025 by bntng
