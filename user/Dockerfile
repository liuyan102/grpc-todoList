# 设置基础镜像
FROM golang:latest AS builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录
WORKDIR /go_project/grpc/user

# 将代码复制到容器
COPY . .

# 下载依赖信息
RUN go mod download

# 将代码编译成二进制可执行文件
RUN go build -o user cmd/main.go


# 创建一个小镜像
FROM alpine:latest AS user_server

# 设置工作目录
WORKDIR /app

# 复制builder镜像中编译好的user可执行文件和config.yml文件
COPY --from=builder /go_project/grpc/user/user .
COPY --from=builder /go_project/grpc/user/config/config.yml ./config/config.yml

# 设置服务暴漏端口
EXPOSE 10001

# 设置启动命令
CMD ["/app/user"]








