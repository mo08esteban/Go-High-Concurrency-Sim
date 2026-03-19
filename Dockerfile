# 阶段 1：在虚拟空间中建造火箭（编译）
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY binance_tracker.go .
# 编译出极速运行的二进制文件，命名为 tracker
RUN go build -o tracker binance_tracker.go

# 阶段 2：把造好的火箭装进最轻的货舱（运行）
FROM alpine:latest
WORKDIR /root/
# 只把编译好的程序拿过来，丢掉所有笨重的开发工具
COPY --from=builder /app/tracker .
# 赋予执行权限
RUN chmod +x ./tracker

# 设定点火指令
CMD ["./tracker"]
