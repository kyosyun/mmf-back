# ビルドステージ
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピーとビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# 実行ステージ
FROM alpine:latest

WORKDIR /app

# Cloud SQL Auth Proxyと必要なパッケージをインストール
RUN apk add --no-cache \
    ca-certificates \
    wget \
    curl \
    tzdata && \
    wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O /usr/local/bin/cloud_sql_proxy && \
    chmod +x /usr/local/bin/cloud_sql_proxy

# タイムゾーンの設定
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

# セキュリティ対策: 非rootユーザーの作成
RUN adduser -D appuser

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/db ./db

# 所有権の変更
RUN chown -R appuser:appuser /app

# 非rootユーザーに切り替え
USER appuser

# ヘルスチェック用のエンドポイントを指定
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

EXPOSE 8080

CMD ["./main", "app"] 