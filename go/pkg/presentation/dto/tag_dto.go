package dto

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

// リクエスト構造体
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

// レスポンス構造体
type TagResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ThreadsCount int       `json:"threads_count"`
	CreatedAt    time.Time `json:"created_at"`
}

type TagListResponse struct {
	Tags  []TagResponse `json:"tags"`
	Total int           `json:"total"`
}

// エンティティからDTOへの変換
func NewTagResponse(tag *entity.Tag) *TagResponse {
	return &TagResponse{
		ID:           tag.ID,
		Name:         tag.Name,
		ThreadsCount: tag.ThreadsCount,
		CreatedAt:    tag.CreatedAt,
	}
}

func NewTagListResponse(tags []*entity.Tag, total int) *TagListResponse {
	response := &TagListResponse{
		Tags:  make([]TagResponse, len(tags)),
		Total: total,
	}
	for i, tag := range tags {
		response.Tags[i] = *NewTagResponse(tag)
	}
	return response
}

// 検索関連
type SearchTagRequest struct {
	Query  string `json:"query" validate:"required"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type TagStatsResponse struct {
	TotalTags         int `json:"total_tags"`
	TotalThreads      int `json:"total_threads"`
	AverageThreadsTag int `json:"average_threads_tag"`
}
