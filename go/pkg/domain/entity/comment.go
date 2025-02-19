package entity

import "time"

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
		return ErrInvalidPostID
	}
	if c.UserID == "" {
		return ErrInvalidUserID
	}
	if c.Content == "" {
		return ErrEmptyCommentContent
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
