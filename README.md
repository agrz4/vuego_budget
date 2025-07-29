# 💰 BudgetApp: Aplikasi Manajemen Budget dan Pengeluaran

<div align="center">

![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=for-the-badge&logo=vue.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-1.x-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-3.x-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)

**Aplikasi web full-stack sederhana untuk membantu pengguna mengelola budget dan melacak pengeluaran mereka.**

</div>

---

## 🎯 Fitur Utama

### 🔐 **Autentikasi Pengguna**
- **Registrasi**: Buat akun pengguna baru
- **Login**: Masuk ke akun yang ada menggunakan username dan password
- **JWT Authentication**: Menggunakan JSON Web Token untuk autentikasi sesi

### 📊 **Dashboard Pengguna**
- Menampilkan ringkasan total budget, total pengeluaran, dan budget tersisa
- Visualisasi pengeluaran per kategori dalam bentuk pie chart
- Tampilan responsif untuk desktop dan mobile

### 💸 **Manajemen Budget**
- ✅ Tambah budget baru dengan kategori, jumlah, dan deskripsi
- 📋 Lihat daftar semua budget yang dibuat
- ✏️ Edit detail budget yang ada
- 🗑️ Hapus budget

### 💳 **Manajemen Pengeluaran**
- ✅ Tambah pengeluaran baru dengan kategori, jumlah, deskripsi, dan tanggal
- 📋 Lihat daftar semua pengeluaran
- 🔍 **Filter & Pencarian**: Filter pengeluaran berdasarkan kategori, rentang tanggal, dan pencarian berdasarkan deskripsi atau kategori
- ✏️ Edit detail pengeluaran yang ada
- 🗑️ Hapus pengeluaran

### ✅ **Validasi & Keamanan**
- Validasi input di frontend dan backend untuk memastikan integritas data
- Password hashing menggunakan bcrypt
- JWT token untuk autentikasi yang aman

---

## 🛠️ Stack Teknologi

### **Frontend**
| Teknologi | Deskripsi |
|-----------|-----------|
| ![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=flat&logo=vue.js&logoColor=white) | Framework JavaScript progresif untuk membangun antarmuka pengguna |
| ![Composition API](https://img.shields.io/badge/Composition_API-Vue_3-4FC08D?style=flat) | Untuk logika komponen yang lebih terorganisir dan dapat digunakan kembali |
| ![Vue Router](https://img.shields.io/badge/Vue_Router-4.x-4FC08D?style=flat) | Untuk routing di sisi klien |
| ![Axios](https://img.shields.io/badge/Axios-1.x-5A29E4?style=flat&logo=axios&logoColor=white) | HTTP client berbasis Promise untuk membuat request ke API backend |
| ![TailwindCSS](https://img.shields.io/badge/TailwindCSS-3.x-38B2AC?style=flat&logo=tailwind-css&logoColor=white) | Framework CSS utility-first untuk styling yang cepat dan responsif |
| ![Chart.js](https://img.shields.io/badge/Chart.js-4.x-FF6384?style=flat&logo=chart.js&logoColor=white) | Untuk visualisasi data (pie chart) |
| ![Lodash](https://img.shields.io/badge/Lodash-4.x-3492FF?style=flat&logo=lodash&logoColor=white) | Untuk utilitas seperti debounce pada input pencarian |

### **Backend**
| Teknologi | Deskripsi |
|-----------|-----------|
| ![Go](https://img.shields.io/badge/Go-1.x-00ADD8?style=flat&logo=go&logoColor=white) | Bahasa pemrograman untuk backend |
| ![Gin](https://img.shields.io/badge/Gin-1.x-00ADD8?style=flat) | Web framework untuk REST API |
| ![GORM](https://img.shields.io/badge/GORM-1.x-00ADD8?style=flat) | ORM untuk database |
| ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791?style=flat&logo=postgresql&logoColor=white) | Database utama |
| ![JWT](https://img.shields.io/badge/JWT-Authentication-000000?style=flat&logo=json-web-tokens&logoColor=white) | Autentikasi menggunakan JSON Web Token |
| ![bcrypt](https://img.shields.io/badge/bcrypt-Password_Hashing-00ADD8?style=flat) | Hashing password untuk keamanan |

---

## 📁 Struktur Folder Proyek

```
vuego_budget/
├── 📁 backend/                    # Proyek Golang
│   ├── 📄 main.go                 # Entry point aplikasi backend
│   ├── 📁 models/                 # Definisi model database (User, Budget, Expense)
│   ├── 📁 controllers/            # Logika handler untuk API (Auth, Budget, Expense)
│   ├── 📁 middleware/             # Middleware kustom (JWT authentication)
│   ├── 📁 database/               # Konfigurasi koneksi database
│   ├── 📁 utils/                  # Utilitas (JWT generation/parsing)
│   ├── 📄 go.mod                  # Go modules dan dependencies
│   └── 📄 go.sum
├── 📁 frontend/                   # Proyek Vue.js
│   ├── 📁 public/                 # Asset statis
│   ├── 📁 src/                    # Kode sumber Vue.js
│   │   ├── 📁 assets/             # Asset statis lainnya (e.g., gambar)
│   │   ├── 📁 components/         # Komponen Vue yang dapat digunakan kembali (e.g., Chart.vue)
│   │   ├── 📁 router/             # Konfigurasi Vue Router
│   │   ├── 📁 services/           # Layanan API untuk interaksi backend
│   │   ├── 📁 views/              # View/halaman aplikasi
│   │   │   ├── 📁 Auth/           # Halaman Login dan Register
│   │   │   ├── 📁 Budget/         # Halaman daftar dan form budget
│   │   │   ├── 📁 Expense/        # Halaman daftar dan form pengeluaran
│   │   │   └── 📄 Dashboard.vue   # Halaman dashboard
│   │   ├── 📄 App.vue             # Komponen aplikasi utama
│   │   └── 📄 main.js             # Entry point aplikasi Vue.js
│   ├── 📄 index.html              # File HTML utama
│   ├── 📄 package.json            # Dependencies Node.js
│   ├── 📄 postcss.config.js       # Konfigurasi PostCSS (untuk TailwindCSS)
│   ├── 📄 tailwind.config.js      # Konfigurasi TailwindCSS
│   └── 📄 vite.config.js          # Konfigurasi Vite
└── 📄 .env                        # Variabel environment untuk konfigurasi
```

---

## 🚀 Cara Menjalankan Aplikasi

Untuk menjalankan aplikasi, Anda akan menjalankan komponen frontend dan backend secara manual.

### 1️⃣ **Prerequisites Manual**

#### **Install Node.js dan npm/Yarn:**
- Download dari https://nodejs.org/ (versi LTS direkomendasikan)
- Verifikasi instalasi: `node -v` dan `npm -v` (atau `yarn -v`)

#### **Install Golang:**
- Download dari https://golang.org/dl/
- Ikuti instruksi instalasi untuk sistem operasi Anda
- Verifikasi instalasi: `go version`

#### **Instalasi dan Konfigurasi PostgreSQL:**
- Install PostgreSQL di sistem operasi Anda (Windows, macOS, Linux)
- Panduan instalasi dapat ditemukan di situs resmi PostgreSQL: https://www.postgresql.org/download/
- Setelah instalasi, buat user database baru dan database. Contoh perintah (untuk Linux/macOS, atau gunakan pgAdmin di Windows):

```sql
sudo -u postgres psql
CREATE USER user_budget WITH PASSWORD 'password_budget';
CREATE DATABASE budget_db OWNER user_budget;
GRANT ALL PRIVILEGES ON DATABASE budget_db TO user_budget;
\q
```

Pastikan kredensial (`user_budget`, `password_budget`, `budget_db`) sesuai dengan yang akan Anda gunakan di file `.env`.

### 2️⃣ **Konfigurasi File .env**

Update file `.env` di folder root proyek Anda.

```env
# PostgreSQL Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=user_budget
DB_PASSWORD=password_budget
DB_NAME=budget_db

# JWT Configuration
JWT_SECRET=supersecretkeyyangsangatpanjangdanaman

# Golang Backend Configuration
PORT=8080
GIN_MODE=debug
FRONTEND_DOMAIN=http://localhost:5173

# Vue.js Frontend Configuration (accessed by Vite)
VITE_API_BASE_URL=http://localhost:8080/api
```

### 3️⃣ **Menjalankan Backend Golang**

**Navigasi ke Direktori Backend:**
```bash
cd backend
```

**Download Dependencies Golang:**
```bash
go mod tidy
```

**Jalankan Aplikasi Backend:**
```bash
go run main.go
```

Backend server akan berjalan di `http://localhost:8080`.

### 4️⃣ **Menjalankan Frontend Vue.js**

**Navigasi ke Direktori Frontend:**
Buka terminal baru dan navigasi ke folder `frontend/`:

```bash
cd frontend
```

**Install Dependencies Frontend:**
```bash
npm install
npm install chart.js lodash # Pastikan ini juga terinstall
```

atau

```bash
yarn
yarn add chart.js lodash
```

**Jalankan Development Server Frontend:**
```bash
npm run dev
```

atau

```bash
yarn dev
```

Aplikasi frontend akan berjalan di `http://localhost:5173`.

---

## ⚙️ API Endpoints (Backend Golang)

Semua endpoint API dimulai dengan `/api`.

### 🔐 **Authentication**

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `POST` | `/api/register` | Registrasi user baru |
| `POST` | `/api/login` | Login user dan dapatkan JWT token |

**Request Body untuk Register/Login:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response untuk Login:**
```json
{
  "message": "Login successful",
  "token": "jwt_token_string"
}
```

### 💸 **Budgets** (Memerlukan JWT Authentication)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `POST` | `/api/budgets` | Buat budget baru |
| `GET` | `/api/budgets` | Dapatkan semua budget user |
| `GET` | `/api/budgets/:id` | Dapatkan budget berdasarkan ID |
| `PUT` | `/api/budgets/:id` | Update budget berdasarkan ID |
| `DELETE` | `/api/budgets/:id` | Hapus budget berdasarkan ID |

**Request Body untuk Create/Update Budget:**
```json
{
  "category": "string",
  "amount": 1000.00,
  "description": "string"
}
```

### 💳 **Expenses** (Memerlukan JWT Authentication)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `POST` | `/api/expenses` | Buat pengeluaran baru |
| `GET` | `/api/expenses` | Dapatkan semua pengeluaran user |
| `GET` | `/api/expenses/:id` | Dapatkan pengeluaran berdasarkan ID |
| `PUT` | `/api/expenses/:id` | Update pengeluaran berdasarkan ID |
| `DELETE` | `/api/expenses/:id` | Hapus pengeluaran berdasarkan ID |

**Request Body untuk Create/Update Expense:**
```json
{
  "category": "string",
  "amount": 500.00,
  "description": "string",
  "date": "2024-01-15"
}
```

**Query Parameters untuk GET Expenses:**
- `category`: Filter berdasarkan kategori
- `start_date`: Filter dari tanggal (YYYY-MM-DD)
- `end_date`: Filter sampai tanggal (YYYY-MM-DD)
- `search`: Pencarian berdasarkan deskripsi atau kategori

### 📊 **Dashboard** (Memerlukan JWT Authentication)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `GET` | `/api/dashboard` | Dapatkan ringkasan total budget, total pengeluaran, budget tersisa, dan pengeluaran berdasarkan kategori |

---

## ⚠️ Catatan Penting

### 🔒 **Keamanan JWT**
- Pastikan `JWT_SECRET` di file `.env` Anda adalah string yang sangat kuat, panjang, dan acak
- Jangan pernah mengeksposnya di repository publik

### 🗄️ **Database Migration**
- Backend Golang akan secara otomatis melakukan migrasi database (membuat tabel) saat pertama kali dijalankan jika tabel belum ada

### 🌐 **CORS**
- Konfigurasi CORS di backend saat ini hanya mengizinkan `http://localhost:5173`
- Sesuaikan jika Anda menjalankan frontend di domain/port yang berbeda

### ✅ **Validasi**
- Validasi dasar telah diimplementasikan
- Untuk aplikasi produksi, pertimbangkan validasi yang lebih advanced dan error handling yang lebih detail

### 🐛 **Error Handling**
- Aplikasi ini memiliki error handling dasar yang akan menampilkan pesan error kepada pengguna
- Untuk produksi, Anda mungkin ingin mengimplementasikan sistem logging yang lebih komprehensif

---

## 🤝 Kontribusi

Kontribusi selalu diterima! Silakan buat pull request atau buka issue untuk melaporkan bug atau menyarankan fitur baru.

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT - lihat file [LICENSE](LICENSE) untuk detail.

---

<div align="center">

**Selamat menggunakan BudgetApp! 🎉**

*Dibuat dengan ❤️ menggunakan Vue.js dan Golang*

</div>