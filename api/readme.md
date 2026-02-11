# API Documentation

## Endpoints

### POST `/user`
Register a new user.

**Request Body:**
```json
{
    "username": "username",
    "password": "securepassword"
}
```

**Response:**
```json
{
    "id": "bbbb1733-85ec-4eb0-b599-b5c36179185e",
    "username": "test-1234",
    "createdAt": "2026-02-11T14:07:53.521759+07:00",
    "updatedAt": "2026-02-11T14:07:53.521759+07:00"
}
```

---
### POST `/auth/login`
Authenticate a user.

**Request Body:**
```json
{
    "username": "username",
    "password": "securepassword"
}
```

**Response:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expiredAt": "2026-02-11T20:38:26.5877974+07:00"
}
```

---

### GET `/user/me`
Get Logged-in user (requires authentication).

**Header:**
```json
{
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
}
```

**Response:**
```json
{
    "id": "bbbb1733-85ec-4eb0-b599-b5c36179185e",
    "username": "test-1234",
    "createdAt": "2026-02-11T14:07:53.521759+07:00",
    "updatedAt": "2026-02-11T14:07:53.521759+07:00"
}
```

---

## Database Tables

### Users Table (SQLite)

| Column     | Type         | Constraints                |
|------------|--------------|----------------------------|
| id         | VARCHAR(36)  | PRIMARY KEY                |
| username   | VARCHAR(255) | NOT NULL, UNIQUE           |
| password   | VARCHAR(255) | NOT NULL                   |
| created_at | DATETIME     | DEFAULT CURRENT_TIMESTAMP  |
| updated_at | DATETIME     | DEFAULT CURRENT_TIMESTAMP  |