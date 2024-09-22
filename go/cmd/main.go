package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// 環境変数からポート番号を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" // デフォルトのポート番号
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
