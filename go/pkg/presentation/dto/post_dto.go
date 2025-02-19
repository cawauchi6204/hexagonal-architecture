package dto

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

// リクエスト構造体
type CreatePostRequest struct {
	Content string `json:"content" validate:"required"`
}

type UpdatePostRequest struct {
	Content string `json:"content" validate:"required"`
}

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}

// レスポンス構造体
type PostResponse struct {
	ID        string            `json:"id"`
	ThreadID  string            `json:"thread_id"`
	UserID    string            `json:"user_id"`
	User      *UserResponse     `json:"user,omitempty"`
	Content   string            `json:"content"`
	Comments  []CommentResponse `json:"comments,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}

type CommentResponse struct {
	ID        string        `json:"id"`
	PostID    string        `json:"post_id"`
	UserID    string        `json:"user_id"`
	User      *UserResponse `json:"user,omitempty"`
	Content   string        `json:"content"`
	Likes     int           `json:"likes"`
	CreatedAt time.Time     `json:"created_at"`
}

type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int            `json:"total"`
}

// エンティティからDTOへの変換
func NewPostResponse(post *entity.Post) *PostResponse {
	response := &PostResponse{
		ID:        post.ID,
		ThreadID:  post.ThreadID,
		UserID:    post.UserID,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		Comments:  make([]CommentResponse, 0),
	}

	if post.Comments != nil {
		for _, comment := range post.Comments {
			response.Comments = append(response.Comments, *NewCommentResponse(comment))
		}
	}

	return response
}

func NewCommentResponse(comment *entity.Comment) *CommentResponse {
	return &CommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Content:   comment.Content,
		Likes:     comment.Likes,
		CreatedAt: comment.CreatedAt,
	}
}

func NewPostListResponse(posts []*entity.Post, total int) *PostListResponse {
	response := &PostListResponse{
		Posts: make([]PostResponse, len(posts)),
		Total: total,
	}
	for i, post := range posts {
		response.Posts[i] = *NewPostResponse(post)
	}
	return response
}
