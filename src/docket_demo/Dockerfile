#FROM goland:alpine AS builder
#
## 为我们的镜像设置必要的环境变量
#ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
## 移动到工作目录: /build
#WORKDIR /build
#
## 将代码复制到容器中
#COPY . .
#
## 将我们的代码编译成二进制可执行文件 app
#RUN go build -o app .
#
#FROM scratch
#
## 从builder镜像中把 /dist/app 拷贝到当前目录
#COPY --from=builder /build/app /
#
## 需要运行的命令
#ENTRYPOINT ["/app"]
#
#
##LABEL authors="Eutop1a"
##
##ENTRYPOINT ["top", "-b"]
#
#FROM golang:alpine AS builder
#
## 为我们的镜像设置必要的环境变量
#ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
## 移动到工作目录：/build
#WORKDIR /build
#
## 将代码复制到容器中
#COPY . .
#
## 将我们的代码编译成二进制可执行文件 app
#RUN go build -o app .
#
####################
## 接下来创建一个小镜像
####################
#FROM scratch
#
## 从builder镜像中把/dist/app 拷贝到当前目录
#COPY --from=builder /build/app /
#
## 需要运行的命令
#ENTRYPOINT ["/app"]

# 使用官方的Golang镜像作为基础镜像
FROM golang:1.16-alpine

# 设置工作目录
WORKDIR /app

# 将本地文件复制到容器内的工作目录
COPY main.go .

# 构建Go应用程序
RUN go build -o main main.go

# 暴露应用程序使用的端口
EXPOSE 8888

# 启动应用程序
CMD ["./main"]
