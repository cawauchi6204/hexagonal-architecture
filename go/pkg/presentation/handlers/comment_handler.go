package handlers

import (
	"net/http"
	"strconv"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/dto"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUseCase usecase.CommentUseCase
}

func NewCommentHandler(commentUseCase usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{
		commentUseCase: commentUseCase,
	}
}

// @Summary コメント取得
// @Description 指定されたコメントを取得する
// @Tags comments
// @Produce json
// @Param id path string true "コメントID"
// @Success 200 {object} dto.CommentResponse
// @Failure 404 {object} ErrorResponse
// @Router /comments/{id} [get]
func (h *CommentHandler) Get(c echo.Context) error {
	commentID := c.Param("id")
	comment, err := h.commentUseCase.GetComment(c.Request().Context(), commentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Comment not found")
	}

	return c.JSON(http.StatusOK, dto.NewCommentResponse(comment))
}

// @Summary コメント更新
// @Description コメントを更新する
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "コメントID"
// @Param request body dto.CreateCommentRequest true "更新情報"
// @Success 200 {object} dto.CommentResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /comments/{id} [put]
func (h *CommentHandler) Update(c echo.Context) error {
	commentID := c.Param("id")
	var req dto.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := h.commentUseCase.UpdateComment(c.Request().Context(), commentID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewCommentResponse(comment))
}

// @Summary コメント削除
// @Description コメントを削除する
// @Tags comments
// @Param id path string true "コメントID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /comments/{id} [delete]
func (h *CommentHandler) Delete(c echo.Context) error {
	commentID := c.Param("id")
	if err := h.commentUseCase.DeleteComment(c.Request().Context(), commentID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary いいね追加
// @Description コメントにいいねを追加する
// @Tags comments
// @Produce json
// @Param id path string true "コメントID"
// @Success 200 {object} dto.CommentResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /comments/{id}/like [post]
func (h *CommentHandler) AddLike(c echo.Context) error {
	commentID := c.Param("id")
	userID := c.Get("user_id").(string) // JWTから取得

	if err := h.commentUseCase.AddLike(c.Request().Context(), commentID, userID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := h.commentUseCase.GetComment(c.Request().Context(), commentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Comment not found")
	}

	return c.JSON(http.StatusOK, dto.NewCommentResponse(comment))
}

// @Summary いいね削除
// @Description コメントのいいねを削除する
// @Tags comments
// @Produce json
// @Param id path string true "コメントID"
// @Success 200 {object} dto.CommentResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /comments/{id}/unlike [post]
func (h *CommentHandler) RemoveLike(c echo.Context) error {
	commentID := c.Param("id")
	userID := c.Get("user_id").(string) // JWTから取得

	if err := h.commentUseCase.RemoveLike(c.Request().Context(), commentID, userID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := h.commentUseCase.GetComment(c.Request().Context(), commentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Comment not found")
	}

	return c.JSON(http.StatusOK, dto.NewCommentResponse(comment))
}

// @Summary 投稿のコメント一覧取得
// @Description 投稿のコメント一覧を取得する
// @Tags comments
// @Produce json
// @Param post_id path string true "投稿ID"
// @Param limit query int false "取得件数" default(10)
// @Param offset query int false "開始位置" default(0)
// @Success 200 {array} dto.CommentResponse
// @Router /posts/{post_id}/comments [get]
func (h *CommentHandler) ListByPost(c echo.Context) error {
	postID := c.Param("post_id")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if offset < 0 {
		offset = 0
	}

	comments, err := h.commentUseCase.GetPostComments(c.Request().Context(), postID, offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}
