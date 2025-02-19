package handlers

import (
	"net/http"
	"strconv"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/dto"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type ThreadHandler struct {
	threadUseCase usecase.ThreadUseCase
}

func NewThreadHandler(threadUseCase usecase.ThreadUseCase) *ThreadHandler {
	return &ThreadHandler{
		threadUseCase: threadUseCase,
	}
}

// @Summary スレッド作成
// @Description 新規スレッドを作成する
// @Tags threads
// @Accept json
// @Produce json
// @Param request body dto.CreateThreadRequest true "スレッド作成情報"
// @Success 201 {object} dto.ThreadResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /threads [post]
func (h *ThreadHandler) Create(c echo.Context) error {
	var req dto.CreateThreadRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := c.Get("user_id").(string) // JWTから取得
	thread, err := h.threadUseCase.CreateThread(c.Request().Context(), req.Title, userID, req.TagIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.NewThreadResponse(thread))
}

// @Summary スレッド取得
// @Description 指定されたスレッドを取得する
// @Tags threads
// @Produce json
// @Param id path string true "スレッドID"
// @Success 200 {object} dto.ThreadResponse
// @Failure 404 {object} ErrorResponse
// @Router /threads/{id} [get]
func (h *ThreadHandler) Get(c echo.Context) error {
	threadID := c.Param("id")
	thread, err := h.threadUseCase.GetThread(c.Request().Context(), threadID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Thread not found")
	}

	return c.JSON(http.StatusOK, dto.NewThreadResponse(thread))
}

// @Summary スレッド更新
// @Description スレッドを更新する
// @Tags threads
// @Accept json
// @Produce json
// @Param id path string true "スレッドID"
// @Param request body dto.UpdateThreadRequest true "更新情報"
// @Success 200 {object} dto.ThreadResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /threads/{id} [put]
func (h *ThreadHandler) Update(c echo.Context) error {
	threadID := c.Param("id")
	var req dto.UpdateThreadRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	thread, err := h.threadUseCase.UpdateThread(c.Request().Context(), threadID, req.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewThreadResponse(thread))
}

// @Summary スレッド削除
// @Description スレッドを削除する
// @Tags threads
// @Param id path string true "スレッドID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /threads/{id} [delete]
func (h *ThreadHandler) Delete(c echo.Context) error {
	threadID := c.Param("id")
	if err := h.threadUseCase.DeleteThread(c.Request().Context(), threadID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary タグ追加
// @Description スレッドにタグを追加する
// @Tags threads
// @Accept json
// @Produce json
// @Param id path string true "スレッドID"
// @Param request body dto.AddTagRequest true "タグ情報"
// @Success 200 {object} dto.ThreadResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /threads/{id}/tags [post]
func (h *ThreadHandler) AddTag(c echo.Context) error {
	threadID := c.Param("id")
	var req dto.AddTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.threadUseCase.AddTag(c.Request().Context(), threadID, req.TagID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	thread, err := h.threadUseCase.GetThread(c.Request().Context(), threadID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Thread not found")
	}

	return c.JSON(http.StatusOK, dto.NewThreadResponse(thread))
}

// @Summary スレッド一覧取得
// @Description スレッドの一覧を取得する
// @Tags threads
// @Produce json
// @Param limit query int false "取得件数" default(10)
// @Param offset query int false "開始位置" default(0)
// @Success 200 {object} dto.ThreadListResponse
// @Router /threads [get]
func (h *ThreadHandler) List(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if offset < 0 {
		offset = 0
	}

	threads, err := h.threadUseCase.GetLatestThreads(c.Request().Context(), limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	total := len(threads) // TODO: 実際の総件数を取得する処理を実装

	return c.JSON(http.StatusOK, dto.NewThreadListResponse(threads, total))
}
