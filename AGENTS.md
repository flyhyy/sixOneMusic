# AGENTS.md

- Thinking思考过程使用中文表述
- Reply回答用中文回复
  NAS music player: Vue 3 frontend + Go backend, deployed via Docker Compose on a NAS.

## Repo layout (non-obvious)

- The root is a git repo but tracks only `.gitignore`, `README.md`, `docker-compose.yml`.
- `frontend/` and `backend/` are **separate nested git repos** (each has its own `.git`, untracked by root). Commit each independently.
- `music/` is the local media library (gitignored), NOT committed. The backend scans this dir in dev.
- `frontend/AGENTS.md` already documents the frontend in detail — read it before touching frontend code; don't duplicate it here.

## Backend (`backend/`)

Go 1.26, Gin + GORM + glebarez/sqlite. Module name is `fnMusicServe`.

Commands (run from `backend/`):

- `go run ./cmd` — start server on `:8061` (entrypoint `cmd/main.go`).
- `air` — hot-reload dev server (config in `.air.toml`).
- `go build ./...` — type/compile check; no tests or lint exist.

Gotchas:

- Package path is `interal` (typo of "internal") everywhere — imports are `fnMusicServe/interal/...`. Do not "fix" it.
- Scanning requires **ffmpeg/ffprobe** on PATH (`interal/utils/ff_mpeg.go` shells out to `ffprobe`). The Docker image installs it; local dev needs it too.
- SQLite DB is created at `data/nas_music.db` **relative to the process working directory** (`os.Getwd()`), with `AutoMigrate` on startup (`interal/repository/db.go`). Run from `backend/` or the DB lands elsewhere.
- Music paths are also resolved relative to cwd. Scan flow: first write folders via `POST /api/folder/write`, then trigger `GET /api/scan/handler`. Covers are written to `music/covers/`; lyrics fetched from network when missing.
- Response envelope is `{ code, data, msg }` (`code === 200` = success, `interal/utils/response.go`). Exception: `/api/audio/play/:id` and `/api/audio/play/lrc/:id` stream raw files with no envelope.
- Auth: JWT secret hardcoded `fnMusic` (`interal/utils/jwt.go`). `/api/auth/{login,register}` are public; most other routes use `JWTAuth()` (`Authorization: Bearer <token>`). Audio play routes are **not** protected (`AudioAuth` middleware exists but is commented out).
- Layering is `control` -> `service` -> `repository`, wired manually in `interal/router/router.go`.

## Frontend (`frontend/`)

See `frontend/AGENTS.md`. Highlights: `npm run dev` (proxies `/api` -> `localhost:8061`), `npm run build` is the only type-check (`vue-tsc -b && vite build`), no lint/tests, strict TS (`verbatimModuleSyntax`, `noUnusedLocals`), components auto-imported by `unplugin-vue-components`.

## Docker (`docker-compose.yml`)

- `docker compose up -d --build` runs both. Frontend `8080:80` (nginx), backend `8061` (internal only; healthcheck uses `nc`).
- The backend bind-mount source is a **hardcoded machine-specific path** `C:/Users/47211/Desktop/OpenCode/music` -> `/app/music` — update it for other machines.
- nginx (`frontend/nginx.conf`) adds SPA fallback and proxies `/api/` -> `go-backend:8061`.

## Conventions

- Chinese is used for comments and UI strings throughout.
