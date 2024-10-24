# Bookstore API with Echo and Gorm

## Deskripsi Proyek

Proyek ini merupakan aplikasi API sederhana untuk mengelola data buku menggunakan framework Echo dan Gorm di Go. API ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data buku yang tersimpan di dalam database SQLite.

## Struktur Direktori

```plaintext
├── cmd
│   └── main.go              # File utama untuk menjalankan aplikasi
├── pkg
│   ├── config
│   │   └── app.go           # Inisialisasi database SQLite menggunakan Gorm
│   ├── controllers
│   │   └── book-controllers.go # Berisi handler untuk operasi CRUD pada buku
│   ├── models
│   │   └── book.go          # Model Book untuk mendefinisikan struktur data buku
│   ├── routes
│   │   └── bookstore-routes.go # Mengatur routing aplikasi
│   └── utils
│       └── utils.go         # Fungsi utilitas untuk penanganan error
├── bookstore.db              # File database SQLite
├── go.mod                    # File Go module untuk dependency
├── go.sum                    # Checksum dependency Go
```

## Instalasi dan Penggunaan

### Prasyarat
- Go 1.16 atau versi yang lebih baru harus terpasang di sistem Anda.
- Pastikan Go module sudah aktif (`go mod init` jika belum aktif).

### Cara Menginstal

1. Clone repository ini:
    ```bash
    git clone https://github.com/adtbch/Weekly-Task3.git
    ```

2. Masuk ke direktori proyek:
    ```bash
    cd repo-name
    ```

3. Unduh dependensi yang diperlukan:
    ```bash
    go mod tidy
    ```

4. Jalankan aplikasi:
    ```bash
    go run cmd/main.go
    ```

5. Aplikasi akan berjalan di `localhost:8080`.

### Endpoints API

| Metode  | Endpoint        | Deskripsi                       |
|---------|-----------------|----------------------------------|
| `GET`   | `/books`        | Mengambil semua buku             |
| `POST`  | `/books`        | Menambahkan buku baru            |
| `GET`   | `/books/:id`    | Mengambil detail buku berdasarkan ID |
| `PUT`   | `/books/:id`    | Memperbarui buku berdasarkan ID  |
| `DELETE`| `/books/:id`    | Menghapus buku berdasarkan ID    |

### Contoh Request

- **GET /books**:
    ```bash
    curl -X GET http://localhost:8080/books
    ```

- **POST /books**:
    ```bash
    curl -X POST http://localhost:8080/books -d '{"title":"Golang Basics","author":"John Doe","price":50000}' -H "Content-Type: application/json"
    ```

### Struktur Data Buku (Model)

```go
type Book struct {
    gorm.Model
    Title  string `json:"title"`
    Author string `json:"author"`
    Price  int    `json:"price"`
}
```

### Penanganan Error

Untuk penanganan error, aplikasi menggunakan fungsi `ErrorResponse` yang terletak di `pkg/utils/utils.go`. Fungsi ini memformat response error dalam bentuk JSON.

### Inisialisasi Database

Database diinisialisasi di `pkg/config/app.go`. Aplikasi menggunakan database SQLite, dan model `Book` di-migrate secara otomatis ketika aplikasi dijalankan.

## Dependencies

- [Echo](https://echo.labstack.com/) - Framework web yang digunakan untuk membuat API.
- [Gorm](https://gorm.io/) - ORM yang digunakan untuk berinteraksi dengan database.
- [SQLite](https://www.sqlite.org/) - Database yang digunakan untuk menyimpan data buku.

## Contact
Jika Anda memiliki pertanyaan atau masukan, jangan ragu untuk menghubungi kami:

- Nama: Aditya Bachtiar
- Email: adit.bachtiar091@gmail.com