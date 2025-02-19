package handlers

import (
	"net/http"
	"strconv"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/dto"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postUseCase usecase.PostUseCase
}

func NewPostHandler(postUseCase usecase.PostUseCase) *PostHandler {
	return &PostHandler{
		postUseCase: postUseCase,
	}
}

// @Summary 投稿作成
// @Description スレッドに新規投稿を作成する
// @Tags posts
// @Accept json
// @Produce json
// @Param thread_id path string true "スレッドID"
// @Param request body dto.CreatePostRequest true "投稿作成情報"
// @Success 201 {object} dto.PostResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /threads/{thread_id}/posts [post]
func (h *PostHandler) Create(c echo.Context) error {
	var req dto.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	threadID := c.Param("thread_id")
	userID := c.Get("user_id").(string) // JWTから取得

	post, err := h.postUseCase.CreatePost(c.Request().Context(), threadID, userID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.NewPostResponse(post))
}

// @Summary 投稿取得
// @Description 指定された投稿を取得する
// @Tags posts
// @Produce json
// @Param id path string true "投稿ID"
// @Success 200 {object} dto.PostResponse
// @Failure 404 {object} ErrorResponse
// @Router /posts/{id} [get]
func (h *PostHandler) Get(c echo.Context) error {
	postID := c.Param("id")
	post, err := h.postUseCase.GetPost(c.Request().Context(), postID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	return c.JSON(http.StatusOK, dto.NewPostResponse(post))
}

// @Summary 投稿更新
// @Description 投稿を更新する
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "投稿ID"
// @Param request body dto.UpdatePostRequest true "更新情報"
// @Success 200 {object} dto.PostResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /posts/{id} [put]
func (h *PostHandler) Update(c echo.Context) error {
	postID := c.Param("id")
	var req dto.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	post, err := h.postUseCase.UpdatePost(c.Request().Context(), postID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewPostResponse(post))
}

// @Summary 投稿削除
// @Description 投稿を削除する
// @Tags posts
// @Param id path string true "投稿ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /posts/{id} [delete]
func (h *PostHandler) Delete(c echo.Context) error {
	postID := c.Param("id")
	if err := h.postUseCase.DeletePost(c.Request().Context(), postID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary コメント追加
// @Description 投稿にコメントを追加する
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "投稿ID"
// @Param request body dto.CreateCommentRequest true "コメント情報"
// @Success 201 {object} dto.CommentResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /posts/{id}/comments [post]
func (h *PostHandler) AddComment(c echo.Context) error {
	postID := c.Param("id")
	var req dto.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := c.Get("user_id").(string) // JWTから取得

	comment, err := h.postUseCase.AddComment(c.Request().Context(), postID, userID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.NewCommentResponse(comment))
}

// @Summary スレッドの投稿一覧取得
// @Description スレッドの投稿一覧を取得する
// @Tags posts
// @Produce json
// @Param thread_id path string true "スレッドID"
// @Param limit query int false "取得件数" default(10)
// @Param offset query int false "開始位置" default(0)
// @Success 200 {object} dto.PostListResponse
// @Router /threads/{thread_id}/posts [get]
func (h *PostHandler) ListByThread(c echo.Context) error {
	threadID := c.Param("thread_id")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if offset < 0 {
		offset = 0
	}

	posts, err := h.postUseCase.GetThreadPosts(c.Request().Context(), threadID, offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	total := len(posts) // TODO: 実際の総件数を取得する処理を実装

	return c.JSON(http.StatusOK, dto.NewPostListResponse(posts, total))
}
