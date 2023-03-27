# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 builder 以便后续引用
FROM golang:1.18-alpine AS builder

# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# 安装 Git
#RUN apk --no-cache add git

# 安装gcc
#RUN apk --no-cache add gcc musl-dev

# 设置工作目录
WORKDIR /app

# 将代码拷贝到镜像中
COPY . .

# 先进行依赖下载
RUN go mod tidy && go mod download

# 构建二进制文件，设置了一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o chatgpt-web

# 下面是第二阶段的镜像构建，和之前保持一致
FROM alpine:latest

# 安装相关软件和库
RUN apk update && apk add --no-cache bash supervisor ca-certificates

# 设置工作目录
WORKDIR /app

# 复制资源和静态文件
COPY resources ./resources
COPY static ./static

# 从上一个阶段构建的 builder 容器中复制二进制文件
COPY --from=builder /app/chatgpt-web .

# 添加配置文件和 supervisord 配置文件
COPY supervisord.conf /etc/supervisord.conf
COPY config.dev.json ./config.json

# 通过 Supervisor 管理服务
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
