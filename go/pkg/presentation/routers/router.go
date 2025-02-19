package routers

import (
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/handlers"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	userUseCase usecase.UserUseCase,
	threadUseCase usecase.ThreadUseCase,
	postUseCase usecase.PostUseCase,
	commentUseCase usecase.CommentUseCase,
	tagUseCase usecase.TagUseCase,
) *echo.Echo {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userUseCase)
	threadHandler := handlers.NewThreadHandler(threadUseCase)
	postHandler := handlers.NewPostHandler(postUseCase)
	commentHandler := handlers.NewCommentHandler(commentUseCase)
	tagHandler := handlers.NewTagHandler(tagUseCase)

	// ヘルスチェック
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// ユーザー関連のルート
	users := e.Group("/users")
	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
		users.GET("", userHandler.List)
		users.GET("/:id", userHandler.GetProfile)
		users.PUT("/:id", userHandler.UpdateProfile)
		users.POST("/:id/follow", userHandler.Follow)
		users.POST("/:id/unfollow", userHandler.Unfollow)
	}

	// スレッド関連のルート
	threads := e.Group("/threads")
	{
		threads.POST("", threadHandler.Create)
		threads.GET("", threadHandler.List)
		threads.GET("/:id", threadHandler.Get)
		threads.PUT("/:id", threadHandler.Update)
		threads.DELETE("/:id", threadHandler.Delete)

		// スレッドの投稿関連
		threads.POST("/:thread_id/posts", postHandler.Create)
		threads.GET("/:thread_id/posts", postHandler.ListByThread)

		// スレッドのタグ関連
		threads.POST("/:id/tags", threadHandler.AddTag)
		threads.GET("/:thread_id/tags", tagHandler.ListByThread)
	}

	// 投稿関連のルート
	posts := e.Group("/posts")
	{
		posts.GET("/:id", postHandler.Get)
		posts.PUT("/:id", postHandler.Update)
		posts.DELETE("/:id", postHandler.Delete)

		// 投稿のコメント関連
		posts.POST("/:id/comments", postHandler.AddComment)
		posts.GET("/:post_id/comments", commentHandler.ListByPost)
	}

	// コメント関連のルート
	comments := e.Group("/comments")
	{
		comments.GET("/:id", commentHandler.Get)
		comments.PUT("/:id", commentHandler.Update)
		comments.DELETE("/:id", commentHandler.Delete)
		comments.POST("/:id/like", commentHandler.AddLike)
		comments.POST("/:id/unlike", commentHandler.RemoveLike)
	}

	// タグ関連のルート
	tags := e.Group("/tags")
	{
		tags.POST("", tagHandler.Create)
		tags.GET("", tagHandler.List)
		tags.GET("/popular", tagHandler.ListPopular)
		tags.GET("/:id", tagHandler.Get)
		tags.PUT("/:id", tagHandler.Update)
		tags.DELETE("/:id", tagHandler.Delete)
	}

	return e
}
