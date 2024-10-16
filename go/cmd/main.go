package main

import (
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/application/core"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/routers"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/service_locater"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	// 環境変数からポート番号を取得
	port := core.MustGetEnv("PORT")
	if port == "" {
		port = "8080" // デフォルトのポート番号
	}
	env := core.MustGetEnv("ENV")
	s := service_locater.BuildServiceLocater(env)
	e := routers.NewRouter(s)

	e.Logger.Fatal(e.Start(":" + port))

}
