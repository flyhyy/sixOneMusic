# SixOneMusic · NAS 音乐服务

一个自托管的 NAS 音乐播放器：Go 后端扫描本地音乐库（自动提取元数据、封面、歌词），Vue 3 前端提供在线播放、歌单、收藏、歌手 / 专辑 / 曲风浏览等能力。前后端通过 Docker Compose 一键部署。

- 前端：Vue 3 + TypeScript + Vite + Pinia + Tailwind CSS
- 后端：Go + Gin + GORM + SQLite（glebarez/sqlite，纯 Go，无需 CGO）
- 部署：Docker Compose（nginx 托管前端，Go 服务仅内网暴露）

## 功能特性

- 音乐库扫描：基于 `ffprobe` 解析音频元数据（标题、歌手、专辑、时长、码率等）
- 封面与歌词：扫描时提取内嵌封面写入 `music/covers/`，缺失歌词时联网抓取
- 播放：支持音频流式播放与 LRC 歌词加载
- 组织浏览：按曲风、歌手、专辑分类查看歌曲
- 歌单：创建、重命名、删除歌单，添加 / 移除歌曲
- 收藏：歌曲收藏与"我的收藏"列表
- 账号：注册 / 登录，JWT 鉴权

## 目录结构

```
sixOneMusic/
├── backend/            # Go 后端（独立 git 仓库）
│   ├── cmd/main.go     # 入口，监听 :8061
│   ├── interal/        # 业务代码（注意：包名拼写为 interal）
│   │   ├── control/    # HTTP 控制器
│   │   ├── service/    # 业务逻辑
│   │   ├── repository/ # 数据访问（SQLite）
│   │   ├── router/     # 路由注册
│   │   ├── model/      # 数据模型 / 请求响应结构
│   │   └── utils/      # 工具（JWT、响应封装、ffprobe 调用）
│   └── Dockerfile
├── frontend/           # Vue 3 前端（独立 git 仓库）
│   ├── src/
│   ├── nginx.conf      # SPA 回退 + /api 反向代理
│   └── Dockerfile
├── music/              # 本地音乐库（gitignore，不提交）
└── docker-compose.yml  # 一键部署前后端
```

---

## 部署

### 方式一：Docker Compose（推荐）

前置要求：已安装 Docker 与 Docker Compose，宿主机有可用的音乐目录。

1. **准备音乐目录**

   在宿主机准备一个存放音乐的文件夹，例如 `C:/xxxx/music`（Windows）或 `/volume1/music`（NAS）。

2. **修改 `docker-compose.yml` 中的音乐目录挂载**

   ```yaml
   services:
     go-backend:
       volumes:
         - type: bind
           source: C:/xxxx/music # 改为你的音乐目录 ,绝对路径
           target: /app/music
   ```

   > `source` 是宿主机路径，`target: /app/music` 是容器内路径。扫描时填写的路径必须是**容器内路径**（即 `/app/music`），而不是宿主机路径。

3. **构建并启动**

   ```bash
   docker compose up -d --build
   ```

4. **访问**

   浏览器打开 `http://<NAS_IP>:8080`。后端 `8061` 仅在 Docker 内部网络暴露，无需对外开放。

**常用运维命令**

```bash
docker compose ps               # 查看容器状态
docker compose logs -f          # 实时查看全部日志
docker compose logs -f go-backend   # 只看后端日志
docker compose down             # 停止并移除容器
docker compose down -v          # 🚨 同时删除数据卷（数据库会丢失）
docker compose up -d --build    # 重新构建并启动
```

数据持久化：SQLite 数据库位于命名卷 `backend-data`，挂载到容器 `/app/data`。

### 方式二：本地开发

**后端**（需 Go 1.26+，且 `ffmpeg` / `ffprobe` 已加入 PATH）

```bash
cd backend
go run ./cmd        # 监听 :8061
# 或热重载（需安装 air）
air
```

> 数据库路径为 `data/nas_music.db`，相对于**进程工作目录**；音乐路径同理。请务必在 `backend/` 目录下启动，否则数据会落到别处。

**前端**（需 Node.js >= 20）

```bash
cd frontend
npm install
npm run dev         # 开发服务器，/api 代理到 localhost:8061
npm run build       # 类型检查 + 生产构建
```

---

## 使用

### 1. 注册 / 登录

打开站点后进入登录页，可切换「注册」创建账号，再登录。登录成功后 token 存入浏览器 `localStorage`，后续请求自动携带 `Authorization: Bearer <token>`。

### 2. 配置音乐路径

进入左侧「设置」→「文件路径」，添加音乐所在的目录：

- Docker 部署：填写容器内路径，通常为 `/app/music`（可填其子目录）。
- 本地开发：填写相对于后端启动目录的路径，例如 `../music`。

### 3. 扫描音乐库

路径配置完成后触发扫描（前端设置页 / 调用扫描接口）。扫描过程会：

- 遍历目录下的音频文件，调用 `ffprobe` 解析元数据并写入数据库；
- 提取内嵌封面保存到 `music/covers/`；
- 对缺失歌词的文件联网抓取 `.lrc`。

可通过扫描状态接口查看进度。

### 4. 浏览与播放

- **发现音乐**：按曲风分类浏览。
- **全部音乐**：分页查看所有已扫描歌曲。
- **歌单列表 / 播放队列 / 我的收藏**：管理歌单与收藏，点击歌曲加入播放。
- 播放时自动请求音频流与对应歌词。

### 5. 退出登录

「设置」页底部点击「退出登录」。

---

## API 一览

所有接口以 `/api` 为前缀，响应统一为 `{ code, data, msg }`，`code === 200` 表示成功。除标注「公开」外均需携带 `Authorization: Bearer <token>`。

### 认证

| 方法 | 路径                 | 说明                                             |
| ---- | -------------------- | ------------------------------------------------ |
| POST | `/api/auth/register` | 注册（公开），body：`{ "userName", "password" }` |
| POST | `/api/auth/login`    | 登录（公开），返回 `{ "token" }`                 |

### 文件路径

| 方法   | 路径                      | 说明                                                     |
| ------ | ------------------------- | -------------------------------------------------------- |
| GET    | `/api/folder/query`       | 查询已配置路径                                           |
| POST   | `/api/folder/write`       | 写入路径，body：`[{ "id": null, "path": "/app/music" }]` |
| DELETE | `/api/folder/del?id=<id>` | 删除路径                                                 |

### 扫描

| 方法 | 路径                | 说明                           |
| ---- | ------------------- | ------------------------------ |
| GET  | `/api/scan/handler` | 开始扫描（扫描所有已配置路径） |
| GET  | `/api/scan/status`  | 查询扫描状态                   |

### 音乐 / 收藏

| 方法 | 路径                                | 说明                                               |
| ---- | ----------------------------------- | -------------------------------------------------- |
| GET  | `/api/music/all?page=1&pageSize=20` | 分页获取全部歌曲                                   |
| PUT  | `/api/music/collect`                | 收藏 / 取消收藏，body：`{ "songId", "isCollect" }` |
| GET  | `/api/music/collectSongs`           | 获取当前用户收藏                                   |

### 歌单

| 方法   | 路径                         | 说明                                 |
| ------ | ---------------------------- | ------------------------------------ |
| POST   | `/api/playList/add`          | 新建歌单，body：`{ "name" }`         |
| GET    | `/api/playList/queryList`    | 歌单列表                             |
| DELETE | `/api/playList/del/:id`      | 删除歌单                             |
| PUT    | `/api/playList/update`       | 重命名歌单，body：`{ "id", "name" }` |
| POST   | `/api/playList/addSong`      | 添加歌曲到歌单                       |
| GET    | `/api/playList/getSongs/:id` | 获取歌单内歌曲                       |

### 分类浏览

| 方法 | 路径                           | 说明           |
| ---- | ------------------------------ | -------------- |
| GET  | `/api/style/list`              | 曲风列表       |
| GET  | `/api/style/song/:styleId`     | 某曲风下的歌曲 |
| GET  | `/api/singer/list`             | 歌手列表       |
| GET  | `/api/singer/song/:singerName` | 某歌手的歌曲   |
| GET  | `/api/album/list`              | 专辑列表       |
| GET  | `/api/album/song/:albumName`   | 某专辑的歌曲   |

### 音频流（无需鉴权，直接返回原始文件）

| 方法 | 路径                      | 说明     |
| ---- | ------------------------- | -------- |
| GET  | `/api/audio/play/:id`     | 播放音频 |
| GET  | `/api/audio/play/lrc/:id` | 获取歌词 |

---

## 注意事项

- **包名拼写**：后端包路径为 `interal`（"internal" 的拼写错误），导入形如 `fnMusicServe/interal/...`，请勿"修正"。
- **JWT 密钥**：当前硬编码为 `fnMusic`（`backend/interal/utils/jwt.go`），生产环境建议改为可配置项。
- **音频接口无鉴权**：`/api/audio/play/:id` 与 `/api/audio/play/lrc/:id` 直接流式返回文件，未受 `JWTAuth` 保护。
- **ffmpeg 依赖**：后端扫描依赖系统 `ffmpeg` / `ffprobe`，Docker 镜像已内置，本地开发需自行安装。
- **Docker 音乐目录**：`docker-compose.yml` 中的 `source` 为机器相关路径，换机器部署时必须修改。
