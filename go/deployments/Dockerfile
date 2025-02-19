FROM golang:1.22.4-alpine AS builder

WORKDIR /app
COPY . .

# 本番用バイナリのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# 実行用の最小イメージ
FROM alpine:3.19

# Cloud SQL Proxyのインストール
RUN apk add --no-cache curl \
    && curl -o /cloud-sql-proxy https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.8.1/cloud-sql-proxy.linux.amd64 \
    && chmod +x /cloud-sql-proxy

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/.env .

# Cloud SQL Proxyを起動し、その後アプリケーションを実行
CMD /cloud-sql-proxy --unix-socket /cloudsql ${INSTANCE_CONNECTION_NAME} & \
    ./server
