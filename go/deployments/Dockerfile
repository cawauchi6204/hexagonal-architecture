FROM golang:1.22.7-alpine as builder

WORKDIR /go/src/app

# 必要なツールのインストール
RUN go install github.com/x-motemen/ghq@v1.4.2
RUN go install github.com/volatiletech/sqlboiler/v4@v4.7.0
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.7.0
RUN go install github.com/cosmtrek/air@v1.43.0

COPY . .

# 依存関係のダウンロードとビルド
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./tmp/app ./cmd/main.go

# 実行用の最小イメージ
FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/src/app/ /

# ポートの公開
EXPOSE 8080

# 実行コマンド
CMD ["/tmp/app"]
