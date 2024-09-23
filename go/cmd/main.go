package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra/repository_impl"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	e := echo.New()

	// 環境変数からポート番号を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" // デフォルトのポート番号
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.GET("users", func(c echo.Context) error {
		db := infra.InitDB()
		userRepositoryImpl := repository_impl.NewUserRepositoryImpl(db)
		rows, err := userRepositoryImpl.FindAll(context.Background())
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, rows)
	})

	e.Logger.Fatal(e.Start(":" + port))

}
