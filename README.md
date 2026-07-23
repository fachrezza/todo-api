# Todo API

To-do list API biasa, dibikin pakai Go + Gin + PostgreSQL + GORM. Ada JWT auth juga.

## Fitur

- CRUD task (create, get all, get by id, update, delete)
- Pagination + filter status + search title/description
- JWT auth (register & login)
- Validasi input
- Error handling
- Ambil data task & total count-nya jalan bareng pakai goroutine + `sync.WaitGroup` 

## Struktur folder

```
.
├── cmd/server/main.go
├── config
├── internal
│   ├── dto
│   ├── handler
│   ├── middleware
│   ├── model
│   ├── repository
│   ├── routes
│   └── service
├── .env
├── go.mod
└── go.sum
```

Pola : layered architecture (handler → service → repository)

## Cara Menjalankan

1. Clone repo

```bash
git clone https://github.com/fachrezza/todo-api.git
cd todo-api
```

2. Install dependency

```bash
go mod tidy
```

3. Siapin database PostgreSQL, contoh:

```
Database : todo_db
Username : postgres
Password : postgres
```

4. Bikin file `.env` di root project:

```env
SERVER_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD="Password Postgres"
DB_NAME=todo_db

JWT_SECRET=my-secret-key
JWT_EXPIRED_HOUR=24
```

5. Jalanin

```bash
go run ./cmd/server
```

## Auth

Login dulu sebelum akses endpoint `/tasks`. Tokennya dikirim di header:

```
Authorization: Bearer <your_token>
```

## Endpoint

**Register** — `POST /register`

```json
{
  "name": "Andi Muhamad Fachreza",
  "email": "Andi@gmail.com",
  "password": "password123"
}
```

**Login** — `POST /login`

```json
{
  "email": "Andi@gmail.com",
  "password": "password123"
}
```

Hasil:

```json
{
  "message": "Login success",
  "token": "jwt-token"
}
```

**Create task** — `POST /tasks`

```json
{
  "title": "Backend Test",
  "description": "Pembuatan Todo-list",
  "status": "pending",
  "due_date": "2026-08-01"
}
```

**Get all tasks** — `GET /tasks`


Contoh: `GET /tasks?page=1&limit=10&status=pending&search=backend`

**Get task by id** — `GET /tasks/{id}`

**Update task** — `PUT /tasks/{id}`

```json
{
  "title": "Tahapan backend test update",
  "description": "Update Task",
  "status": "completed",
  "due_date": "2026-08-02"
}
```

**Delete task** — `DELETE /tasks/{id}`

## Response

Hasil:

```json
{
  "message": "Task created successfully",
  "task": {}
}
```

error:

```json
{ "message": "validation error" }
```

atau

```json
{ "message": "task not found" }
```

## Testing API

Test menggunakan Postman / Insomnia / Thunder Client, dll
## Catatan

Dibuat buat technical test backend — fokusnya di REST API pakai Go + PostgreSQL, dengan JWT auth, validasi, pagination, filter, search, error handling, dan sedikit concurrency di bagian ambil data.