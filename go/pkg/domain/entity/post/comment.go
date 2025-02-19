package post

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity/base"
)

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
	Likes     int
}

func NewComment(postID, userID, content string) *Comment {
	return &Comment{
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
		Likes:     0,
	}
}

func (c *Comment) Validate() error {
	if c.PostID == "" {
		return base.ErrInvalidPostID
	}
	if c.UserID == "" {
		return base.ErrInvalidUserID
	}
	if c.Content == "" {
		return base.ErrEmptyCommentContent
	}
	return nil
}

func (c *Comment) IncrementLikes() {
	c.Likes++
}

func (c *Comment) DecrementLikes() {
	if c.Likes > 0 {
		c.Likes--
	}
}
