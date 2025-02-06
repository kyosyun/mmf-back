# ビルドステージ
FROM golang:1.21-alpine AS builder

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

# セキュリティ対策: 非rootユーザーの作成
RUN adduser -D appuser

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/db ./db

# タイムゾーンの設定
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    apk del tzdata

# 所有権の変更
RUN chown -R appuser:appuser /app

# 非rootユーザーに切り替え
USER appuser

# ヘルスチェック用のエンドポイントを指定
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

EXPOSE 8080

CMD ["./main", "app"] 