package dto

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

// リクエスト構造体
type CreateThreadRequest struct {
	Title  string   `json:"title" validate:"required"`
	TagIDs []string `json:"tag_ids"`
}

type UpdateThreadRequest struct {
	Title string `json:"title" validate:"required"`
}

type AddTagRequest struct {
	TagID string `json:"tag_id" validate:"required"`
}

// レスポンス構造体
type ThreadResponse struct {
	ID        string         `json:"id"`
	Title     string         `json:"title"`
	UserID    string         `json:"user_id"`
	User      *UserResponse  `json:"user,omitempty"`
	Tags      []TagResponse  `json:"tags"`
	Posts     []PostResponse `json:"posts,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
}

type ThreadListResponse struct {
	Threads []ThreadResponse `json:"threads"`
	Total   int              `json:"total"`
}

// エンティティからDTOへの変換
func NewThreadResponse(thread *entity.Thread) *ThreadResponse {
	response := &ThreadResponse{
		ID:        thread.ID,
		Title:     thread.Title,
		UserID:    thread.UserID,
		CreatedAt: thread.CreatedAt,
		Tags:      make([]TagResponse, 0),
		Posts:     make([]PostResponse, 0),
	}

	// タグの変換
	if thread.Tags != nil {
		for _, tag := range thread.Tags {
			response.Tags = append(response.Tags, *NewTagResponse(tag))
		}
	}

	// 投稿の変換
	if thread.Posts != nil {
		for _, post := range thread.Posts {
			response.Posts = append(response.Posts, *NewPostResponse(post))
		}
	}

	return response
}

func NewThreadListResponse(threads []*entity.Thread, total int) *ThreadListResponse {
	response := &ThreadListResponse{
		Threads: make([]ThreadResponse, len(threads)),
		Total:   total,
	}
	for i, thread := range threads {
		response.Threads[i] = *NewThreadResponse(thread)
	}
	return response
}

// 検索関連
type SearchThreadRequest struct {
	Query  string `json:"query" validate:"required"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type ThreadStatsResponse struct {
	TotalThreads int `json:"total_threads"`
	TotalPosts   int `json:"total_posts"`
	TotalTags    int `json:"total_tags"`
}
