
# ============================
# 第一阶段：构建阶段 (Builder)
# ============================


# 使用较新的 Node.js 长期支持版 (Alpine 体积更小)
# node:20：表示需要 Node.js 20 版本的环境（因为编译 Vue 3 需要运行 npm）。
# -alpine：Alpine 是一个极简版的 Linux 操作系统。普通的 Linux 镜像动辄几百兆，而 Alpine 只有 5MB 左右。带 -alpine 后缀意味着这个 Node 环境是装在这个迷你系统里的，能让你的下载构建速度飞快。
# AS builder —— "给这个阶段起个代号" AS 是“作为”的意思。 builder 是你给这个阶段起的一个别名（你完全可以叫它 AS step1 或者 AS 打包机器，但业界约定俗成叫 builder）。

FROM node:20-alpine AS builder 
#  导航定位：告诉 Docker 接下来的操作都在容器的 /app 目录下进行
WORKDIR /app

# 把本机的 package.json 复制到当前目录（也就是 /app 里），用 ./ 即可
COPY package*.json ./
# 在 /app 目录下执行安装
RUN  npm install 
# COPY —— 动作：复制
#  第一个 . (点) —— 来源地：你的电脑当前目录,这个点代表的是宿主机（你的电脑）的当前路径（通常也就是你存放 Dockerfile 和 Vue3 源码的那个项目根目录）。它包含了你的 src 文件夹、public 文件夹、package.json、vite.config.ts 等等所有源码。
# 第二个 . (点) —— 目的地：容器内的当前目录,这个点代表的是容器内部的当前路径。这正好呼应了你上一个问题：因为你在此之前写了一句 WORKDIR /app（把容器当前目录设置为了 /app），所以这里的第二个点，实际上指代的就是容器里的 /app 目录。
COPY . .

RUN npm run build 

FROM nginx:alpine 

COPY --from=builder /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf 

EXPOSE 80

CMD ["nginx","-g","daemon off;"]


# 命令 
# 构建镜像
# docker build -t my-vue-app .
# docker build：告诉 Docker 开始构建镜像。
# -t my-vue-app：-t (tag) 是给这个镜像起个名字。你可以随便起，比如 frontend-v1，这里我们叫它 my-vue-app。
# . (最后的小圆点)：非常重要！它表示上下文路径是“当前目录”，Docker 会在当前目录下寻找 Dockerfile 和源码。

# 运行容器
# docker run -d -p 8080:80 --name vue-web-container my-vue-app

# -d：后台运行 (Detached mode)。让容器在后台默默工作，这样你的命令行终端就不会被卡住，可以继续输入其他命令
# -p 8080:80：端口映射 (Port)。
# 前面的 8080 是你电脑（宿主机）上的端口，你可以改成 3000、9090 等任何未被占用的端口。
# 后面的 80 是容器内部 Nginx 默认暴露的端口。
# 连起来的意思是：当用户访问你电脑的 8080 端口时，Docker 会自动把请求转发给容器里的 Nginx 处理。
# --name vue-web-container：给运行起来的容器起个响亮的名字（比如叫 vue-web-container），方便以后管理（停止、删除等），如果不写，Docker 会随机分配一个毫无意义的英文名字。
# my-vue-app：指定你要运行的镜像名字（就是你在第一步里起的名字）。