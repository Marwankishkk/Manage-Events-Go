# Event Management REST API

A small REST API built with Go and Gin for listing events, managing event details, and registering authenticated users for events.

## Stack

- Go
- Gin
- SQLite (`api.db`)
- JWT (`github.com/golang-jwt/jwt/v5`)
- bcrypt (`golang.org/x/crypto/bcrypt`)

## Prerequisites

- Go 1.25 or compatible toolchain (see `go.mod`)

## Run

From the project root:

```bash
go run .
```

The server listens on `http://localhost:8080`.

The database file is created automatically on first run if it does not exist.

## API

### Public

| Method | Path | Description |
|--------|------|-------------|
| GET | `/events` | List all events |
| GET | `/events/:id` | Get one event |
| POST | `/signup` | Create a user (JSON: `email`, `password`) |
| POST | `/login` | Login (JSON: `email`, `password`); returns a JWT |

### Authenticated

Send `Authorization: <token>` with the JWT returned from login (no `Bearer` prefix in the current implementation).

| Method | Path | Description |
|--------|------|-------------|
| POST | `/events` | Create an event |
| PUT | `/events/:id` | Update an event (owner only) |
| DELETE | `/events/:id` | Delete an event (owner only) |
| POST | `/events/:id/register` | Register the current user for the event |
| DELETE | `/events/:id/register` | Cancel registration |

Event payloads use fields such as `name`, `description`, `location`, and `DataTIme` (time) as defined in the models.

## Project layout

- `main.go` – application entrypoint
- `db/` – SQLite connection and schema
- `models/` – domain types and persistence
- `routes/` – HTTP handlers and route registration
- `middlewares/` – JWT authentication middleware
- `utils/` – password hashing and JWT helpers

## Security note

JWT signing uses a hardcoded secret in `utils/jwt.go`. Replace with environment-based configuration before any production deployment.
