FROM --platform=linux/amd64 golang:1.22.4
ARG PROTOBUF_VERSION=21.12

# protocをインストールするためにunzipを用意
RUN apt update && apt install unzip && \
    apt install -y nodejs \
    npm
# RUN npm i -g ts-proto@${TS_PROTO_VERSION}
# protocのダウンロード
WORKDIR /tmp/protoc
# protocのインストール
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip && \
    # 解凍
    unzip protoc.zip && \
    # バイナリをパスに通す
    mv bin/* /usr/local/bin/ && \
    # ヘッダーファイルをパスに通す
    mv include/* /usr/local/include/

WORKDIR /usr/src/app

# プラグインの導入
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 && \
    go install github.com/x-motemen/ghq@v1.4.0 && \
    ghq get -u https://github.com/googleapis/googleapis.git
