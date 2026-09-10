# AGENTS.md

- Thinking思考过程使用中文表述
- Reply回答用中文回复、

Vue 3 + TypeScript + Vite frontend for a NAS music player. This directory is its own git repo; the sibling `backend/` (Go/Gin/GORM/SQLite) is a separate repo. `docker-compose.yml` at the `sixOneMusic/` parent runs both together.

## Commands

- `npm run dev` — Vite dev server; proxies `/api` -> `http://localhost:8061` (the Go backend).
- `npm run build` — `vue-tsc -b && vite build` (type-check then build; this is the only type-check).
- `npm run preview` — preview the built `dist/`.

No lint or test scripts exist. Formatting is Prettier + `prettier-plugin-tailwindcss` (`.prettierrc`); no ESLint. Node >= 20.

## Conventions that will trip you up

- Strict TS: `erasableSyntaxOnly` (no `enum`/namespaces) and `verbatimModuleSyntax` (type-only imports must be `import type`). `noUnusedLocals`/`noUnusedParameters` are on, so `npm run build` fails on unused vars.
- Path alias `@` -> `src/`.
- Components in `src/components/` are auto-imported and globally registered by `unplugin-vue-components` (deep scan). Do not import them manually in templates. `src/types/components.d.ts` is generated — never hand-edit.
- Icons are referenced via the `Icons` object in `src/config/icons.ts` using `ri:` / `mdi:` / `material-symbols:` / `system-uicons:` names; those JSON icon sets are registered offline in `main.ts`. Add icons by referencing an existing set in `Icons`, not by installing new packages.
- `__APP_VERSION__` global is injected from `package.json` via vite `define` (declared in `src/types/global.d.ts`).
- Chinese is used throughout for comments/UI strings; many files start with a `/* prettier-ignore */` + `@Author`/`@Date`/`@Description` header.

## HTTP / API

- All requests go through `src/http/index.ts` (axios, `baseURL: "/api"`). Response envelope is `{ code, data, msg }`; `code === 200` = success. The interceptor unwraps and returns `data` directly, so `src/api/*` functions resolve with the payload, not the envelope.
- Auth: `Authorization: Bearer <token>`, token in `localStorage` key `token` (see `src/storage/token.ts`). The 401 handler is a stub (no auto-logout/redirect implemented).
- Router guard in `src/router/index.ts` redirects to `/` (login) when no token; `/music/*` is the authed layout. Many routes are commented out in place.

## Structure notes

- `src/pages/publicSongList/index.vue` exports the shared `MusicBase` interface; stores and API files import it from there (not from `src/types`).
- Pinia stores: `src/store/{audio,music,playList}.ts`.
- `docker-compose.yml` (parent dir) serves built `dist/` via nginx; `nginx.conf` adds SPA fallback and proxies `/api` to `go-backend:8061`. Backend bind-mounts the `music/` dir as its media library.
