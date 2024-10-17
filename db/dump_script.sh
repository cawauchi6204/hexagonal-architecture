#!/bin/bash

# 環境変数の設定
export MYSQL_USER=user
export MYSQL_PASSWORD=password
export MYSQL_DATABASE=db
export MYSQL_HOST=127.0.0.1
export MYSQL_PORT=3333

# MySQLサーバーが起動するまで待機
while ! nc -z $MYSQL_HOST $MYSQL_PORT; do
  sleep 1
  echo "MySQLサーバーの起動を待っています..."
done

echo "MySQLサーバーが起動しました。データベースに接続しています..."

# テーブルをdumpする
mysqldump --no-data -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -h "$MYSQL_HOST" -P $MYSQL_PORT "$MYSQL_DATABASE" > ./schema_dump.sql

echo "ダンプが完了しました。ファイルは /dump.sql として保存されました。"