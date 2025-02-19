package handlers

import (
	"net/http"
	"strconv"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/presentation/dto"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	tagUseCase usecase.TagUseCase
}

func NewTagHandler(tagUseCase usecase.TagUseCase) *TagHandler {
	return &TagHandler{
		tagUseCase: tagUseCase,
	}
}

// @Summary タグ作成
// @Description 新規タグを作成する
// @Tags tags
// @Accept json
// @Produce json
// @Param request body dto.CreateTagRequest true "タグ作成情報"
// @Success 201 {object} dto.TagResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /tags [post]
func (h *TagHandler) Create(c echo.Context) error {
	var req dto.CreateTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tag, err := h.tagUseCase.CreateTag(c.Request().Context(), req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.NewTagResponse(tag))
}

// @Summary タグ取得
// @Description 指定されたタグを取得する
// @Tags tags
// @Produce json
// @Param id path string true "タグID"
// @Success 200 {object} dto.TagResponse
// @Failure 404 {object} ErrorResponse
// @Router /tags/{id} [get]
func (h *TagHandler) Get(c echo.Context) error {
	tagID := c.Param("id")
	tag, err := h.tagUseCase.GetTag(c.Request().Context(), tagID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Tag not found")
	}

	return c.JSON(http.StatusOK, dto.NewTagResponse(tag))
}

// @Summary タグ更新
// @Description タグを更新する
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string true "タグID"
// @Param request body dto.UpdateTagRequest true "更新情報"
// @Success 200 {object} dto.TagResponse
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /tags/{id} [put]
func (h *TagHandler) Update(c echo.Context) error {
	tagID := c.Param("id")
	var req dto.UpdateTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tag, err := h.tagUseCase.UpdateTag(c.Request().Context(), tagID, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewTagResponse(tag))
}

// @Summary タグ削除
// @Description タグを削除する
// @Tags tags
// @Param id path string true "タグID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Security BearerAuth
// @Router /tags/{id} [delete]
func (h *TagHandler) Delete(c echo.Context) error {
	tagID := c.Param("id")
	if err := h.tagUseCase.DeleteTag(c.Request().Context(), tagID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary タグ一覧取得
// @Description タグの一覧を取得する
// @Tags tags
// @Produce json
// @Success 200 {object} dto.TagListResponse
// @Router /tags [get]
func (h *TagHandler) List(c echo.Context) error {
	tags, err := h.tagUseCase.GetAllTags(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewTagListResponse(tags, len(tags)))
}

// @Summary 人気タグ取得
// @Description 人気のタグを取得する
// @Tags tags
// @Produce json
// @Param limit query int false "取得件数" default(10)
// @Success 200 {object} dto.TagListResponse
// @Router /tags/popular [get]
func (h *TagHandler) ListPopular(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	tags, err := h.tagUseCase.GetPopularTags(c.Request().Context(), limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewTagListResponse(tags, len(tags)))
}

// @Summary スレッドのタグ一覧取得
// @Description スレッドに関連付けられたタグの一覧を取得する
// @Tags tags
// @Produce json
// @Param thread_id path string true "スレッドID"
// @Success 200 {object} dto.TagListResponse
// @Router /threads/{thread_id}/tags [get]
func (h *TagHandler) ListByThread(c echo.Context) error {
	threadID := c.Param("thread_id")
	tags, err := h.tagUseCase.GetThreadTags(c.Request().Context(), threadID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dto.NewTagListResponse(tags, len(tags)))
}
