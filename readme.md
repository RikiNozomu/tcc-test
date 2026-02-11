# TCC-Test

A full-stack application with an Angular frontend and a Go backend.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Project Structure

```
tcc-test/
├── app/           # Angular frontend
├── api/           # Go backend
└── docker-compose.yaml
```

## Running the Application

1. Clone the repository:
    ```bash
    git clone https://github.com/RikiNozomu/tcc-test.git
    cd tcc-test
    ```

2. Start the application using Docker Compose:
    ```bash
    docker-compose up --build -d
    ```

3. Access the application:
    - Frontend: [http://localhost:4200](http://localhost:4200)
    - Backend API: [http://localhost:8080](http://localhost:8080)

## Stopping the Application

```bash
docker-compose down
```

## Environment Variables

### Backend

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_NAME` | `tcc-test.db` | Database file name |
| `PORT` | `8080` | Server port |
| `REQUEST_DELAY_IN_SECOND` | `10` | Request delay in seconds |
| `MAXIMUM_BRUST` | `100` | Maximum burst limit |
| `JWT_SECRET` | `test-tcc-jwt` | JWT secret key |
| `JWT_TIME_EXPIRED_SECOND` | `3600` | JWT expiration time in seconds |

### Frontend

> **Note:** Users must update the `NG_APP_API_URL` build argument in `docker-compose.yaml` to match their deployment environment.

| Variable | Default | Description |
|----------|---------|-------------|
| `NG_APP_API_URL` | `http://localhost:8080` | Backend API URL |