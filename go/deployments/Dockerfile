FROM golang:1.22.4-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED=0
WORKDIR ${ROOT}

RUN apk update --no-cache && \
    apk add --no-cache --update git \
    mysql-client

RUN go install github.com/volatiletech/sqlboiler/v4@v4.7.0
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.7.0

COPY . .
# RUN go mod download
# RUN CGO_ENABLED=0 go build -o .app ./cmd/app
# ビルドコンテキスト内のファイルをコピー
RUN ls -la
# COPY /go.mod ../go.sum ./

# RUN go mod download

# COPY . .

# EXPOSE 8080
# CMD ["go", "run", "main.go"]