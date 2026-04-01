# auth-service

> REST сервис аутентификации на Go — регистрация и авторизация пользователей с хешированием паролей.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-1.9-blue?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-compose-2496ED?style=flat&logo=docker)

---

## О проекте

Сервис реализует базовую аутентификацию пользователей: регистрацию с хешированием пароля через bcrypt и вход с проверкой credentials. Написан на Go с использованием Gin, подключается к PostgreSQL через pgxpool, миграции управляются через golang-migrate.

---

## Стек

| Технология | Назначение |
|---|---|
| Go + Gin | HTTP сервер и роутинг |
| PostgreSQL | Хранение пользователей |
| pgxpool | Пул соединений с БД |
| bcrypt | Хеширование паролей |
| golang-migrate | Миграции базы данных |
| godotenv | Переменные окружения |
| Docker Compose | Запуск окружения |

---

## API

### POST `/register`

Регистрация нового пользователя.

**Тело запроса:**
```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```

**Ответ `201`:**
```json
{
  "message": "user created"
}
```

---

### POST `/login`

Вход в систему.

**Тело запроса:**
```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```

**Ответ `200`:**
```json
{
  "message": "ok"
}
```

---

## Запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/rrwwmq/auth-service
cd auth-service
```

### 2. Создать `.env` файл

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=auth
SERVER_PORT=8080
```

### 3. Запустить PostgreSQL через Docker

```bash
docker compose up -d
```

### 4. Запустить сервер

```bash
go run ./cmd/main.go
```

---

## Пример через curl

```bash
# Регистрация
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'

# Вход
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

---

## Структура проекта

```
auth-service/
├── cmd/
│   └── main.go
├── internal/
│   ├── domain/
│   ├── handler/
│   ├── repository/
│   └── service/
├── migrations/
├── docker-compose.yml
├── .env.example
└── go.mod
```

---

## Безопасность

Пароли никогда не хранятся в открытом виде — только bcrypt хеш с cost factor 10. Даже при утечке базы данных восстановить оригинальный пароль вычислительно нецелесообразно.
