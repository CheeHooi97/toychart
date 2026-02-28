# toychart

## Air hot reload

This repo's Go module lives in [`backend/`](/Users/hwanchinyang/Documents/personal/toychart/backend), so the `air` config is defined at the repo root and rebuilds that module into [`tmp/`](/Users/hwanchinyang/Documents/personal/toychart/tmp).

Install `air` once:

```bash
go install github.com/air-verse/air@latest
```

Run the backend with hot reload from the repo root:

```bash
air
```

Notes:

- Keep your env vars in `.env` at the repo root or `backend/.env`.
- `air` rebuilds `backend/main.go` into `tmp/toychart-api` and restarts it on Go file changes.

## Local Postgres

The backend expects PostgreSQL on `127.0.0.1:5432` with the credentials currently defined in `backend/.env`.

Start the local database from the repo root:

```bash
docker compose up -d postgres
```

Stop it later with:

```bash
docker compose down
```

Notes:

- The database uses a named Docker volume, so data persists across restarts.
- The app now creates the `toychart` schema automatically before running migrations.
