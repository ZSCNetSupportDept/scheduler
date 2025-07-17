# 构建阶段
FROM golang:1.23 AS builder

WORKDIR /app

# 复制源码
COPY . .

# 构建二进制文件
RUN go build -o /app/scheduler

# 运行阶段
FROM scratch

WORKDIR /app

# 复制构建好的二进制文件
COPY --from=builder ./scheduler ./scheduler

RUN mkdir -p config

EXPOSE 25005

USER nonroot:nonroot

CMD ["./scheduler","--config","config/config.yaml"]
