# 第一阶段：编译代码 (无需任何 C 语言依赖)
FROM golang:1.26-alpine AS builder

WORKDIR /app

# 设置代理
ENV GOPROXY=https://goproxy.cn,direct

# 缓存依赖
COPY go.mod go.sum ./

RUN go mod download

# 复制代码并编译
COPY . .

# 核心改变：关闭 CGO (CGO_ENABLED=0)，编译出纯静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

# 第二阶段：运行环境 (极小体积)
FROM alpine:3.21

WORKDIR /app

# 建议：带上时区数据，防止 Go 代码处理时间时报错
RUN apk add --no-cache tzdata ffmpeg

# 从编译阶段拷贝生成的可执行文件
COPY --from=builder /app/myapp .

EXPOSE  8061

# 创建数据目录
RUN mkdir -p /app/data

# 运行程序
CMD ["./myapp"]
