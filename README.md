## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
Repository adalah hasil percobaan penulis dalam belajar Golang yang mana menggunakan Clean Architecture, yang memisahkan setiap layer aplikasi. Pada umumnya dibedakan menjadi layer :

1. Repository
2. Service
3. Handler

![clean](images/clean.jpg "clean")

## Technologies
Project ini menggunakan beberapa teknologi yaitu :
* Golang Versi 1.15
* Echo Labstack 
* GORM
* Godotenv
* Logrush
* MySQL
## Setup
Sebelum itu pastikan MySQL sudah running dan buat sebuah database dengan nama **golang_crud** dan sesuaikan dengan yang ada di **file .env**

### Running Locally
Berikut langkah-langkah untuk running repo ini secara lokal:

```bash
$ git clone https://github.com/teten-nugraha/golang-crud.git
$ go mod download
$ go run main.go
```

### Tugas 1
### Running with Docker
Aplikasi ini juga dapat dijalankan menggunakan Docker dengan Dockerfile yang sudah dioptimasi menggunakan multi-stage build.

#### Build Docker Image
```bash
$ docker build -f devops-challenge/Dockerfile -t alfin/interview:latest .
```

#### Run Docker Container
```bash
$ docker run -d \
  --name golang-app \
  -p 9999:9999 \
  --env-file .env \
  alfin/interview:latest
```

**Catatan**: 
- Pastikan file `.env` sudah dikonfigurasi dengan benar seperti yang ada di `.env.example`
- Aplikasi akan berjalan di port `9999`
- Docker image menggunakan distroless base untuk keamanan dan ukuran minimal
- Aplikasi berjalan sebagai non-root user di dalam container

### Tugas 2
#### Running with Docker Compose
Jalankan aplikasi beserta dependensinya menggunakan Docker Compose.

```bash
$ docker compose -f devops-challenge/docker-compose.yaml up -d --build
```

Untuk menghentikan dan menghapus container:

```bash
$ docker compose -f devops-challenge/docker-compose.yaml down
```

**Catatan**:
- Pastikan file `.env` sudah dikonfigurasi sesuai `.env.example`
- Port aplikasi tetap `9999`

### Tugas 3
#### CI/CD Pipeline (GitHub Actions)
Workflow CI/CD ada di `.github/workflows/ci-cd.yml` dengan alur:
- Menjalankan `go test ./...` pada setiap pull request dan push ke `main`.
- Build & push Docker image ke GitHub Container Registry (GHCR) **hanya jika** tahap test berhasil.
- Push image hanya berjalan pada push ke branch `main`.

Image akan dipublish ke `ghcr.io/<owner>/<repo>` dengan tag `latest` dan tag `sha`.