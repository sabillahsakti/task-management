# task-management

## Prerequisites

Pastikan menginstal:

- **Go**: [Unduh Go](https://golang.org/dl/)
- **PostgreSQL**: [Instal PostgreSQL](https://www.postgresql.org/download/)

## Instalasi

1. **Clone Repository**

   ```
   git clone https://github.com/sabillahsakti/task-management.git
   cd task-management
   ```
2. **Setup Database dan ENV**

Buat database baru di PostgreSQL, misalnya task_management.
Sesuaikan konfigurasi database di file .env:
```
dsn := "host=localhost user=yourusername password=yourpassword dbname=task_management port=5432"
```

Tambahkan juga JWT_KEY pada .env. Contoh :
```
JWT_KEY=inikeyuntukjwt
```

3. **Instal Dependensi**

Jalankan perintah berikut untuk menginstal dependensi:
```
go mod tidy
```

## Menjalankan API
1. **Jalankan Aplikasi**
```
go run main.go
```

2. **API akan berjalan di port 8080**

## Daftar Endpoint API

| Endpoint                | Metode HTTP | Deskripsi                                       | Contoh Input / Output                                          |
|------------------------ |-------------|------------------------------------------------|--------------------------------------------------------------|
| `/register`        | POST        | Mendaftar pengguna baru                        | **Input:** `{ "username": "sabil", "email": "sabil1@gmail.com", "password": "yourpassword" }`<br>**Output:** `{ "statusCode": 200, "message": "User registered successfully", "data": { "id": 1, "username": "sabil", "email": "sabil1@gmail.com", "created_at": "2024-10-24T10:47:59.53544+07:00", "updated_at": "2024-10-24T10:47:59.53544+07:00" } }` |
| `/login`           | POST        | Masuk pengguna                                 | **Input:** `{ "email": "sabil1@gmail.com", "password": "yourpassword" }`<br>**Output:** `{ "statusCode": 200, "message": "Login successful", "data": { "token": "your_jwt_token" } }` |
| `/api/task`            | GET         | Mendapatkan daftar tugas pengguna dengan filter dan sorting | **Input URL:** `/api/task?status=in-progress&sort_by=due_date&order=desc`<br>**Output:** `{ "statusCode": 200, "message": "Success", "data": [{ "id": 1, "title": "Task 1", "status": "in-progress", "due_date": "2024-10-30" }, ...] }` |
| `/api/task/{id}`       | GET         | Mendapatkan detail tugas berdasarkan ID        | **Output:** `{ "statusCode": 200, "message": "Success", "data": { "id": 1, "title": "Task 1", "status": "in-progress", "due_date": "2024-10-30" } }` |
| `/api/task`            | POST        | Menambahkan tugas baru                         | **Input:** `{ "title": "Task 1", "status": "pending", "due_date": "2024-10-30", "user_id": 1 }`<br>**Output:** `{ "statusCode": 200, "message": "Task created successfully", "data": { "id": 1, "title": "Task 1", "status": "pending", "due_date": "2024-10-30" } }` |
| `/api/task/{id}`       | PUT         | Memperbarui tugas berdasarkan ID               | **Input:** `{ "title": "Updated Task", "status": "completed", "due_date": "2024-11-01" }`<br>**Output:** `{ "statusCode": 200, "message": "Success", "data": { "id": 1, "title": "Updated Task", "status": "completed", "due_date": "2024-11-01" } }` |
| `/api/task/{id}`       | DELETE      | Menghapus tugas berdasarkan ID                 | **Output:** `{ "statusCode": 200, "message": "Task deleted successfully" }` |

**Catatan :**
1. Lakukan login terlebih dahulu untuk mendapatkan token
2. semua endpoint kecuali register dan login harus mengirimkan token ke header dengan key: Authorization dan Value: Bearer token_value. Contoh:
   ```
   Key : Authorization
   Value : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ1c2VybmFtZSI6InNhYmlsIiwiaXNzIjoiZXZvdGluZyIsImV4cCI6MTcyOTgyNDExNX0.rAOqDi54cPayf63YWmGFqPQcSkfkBaHyyoFQWpTMZf8
   ```
