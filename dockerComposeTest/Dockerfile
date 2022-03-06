FROM golang
# 设置 go mod
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app
# 拷贝 go.mod 和 go.sum，如果缺少 go.sum 会导致 tidy 失败
COPY go.mod go.sum ./
# 下载依赖
RUN go mod tidy
RUN ls
# 将当前目录的内容全部复制到容器中
COPY . .
# 编译出二进制文件
RUN go build -o main .
# 暴露对外端口
EXPOSE 6666
# 运行二进制文件
CMD [ "./main" ]


