package entity

import "time"

type Post struct {
	ID        string
	ThreadID  string
	UserID    string
	Content   string
	CreatedAt time.Time
	Comments  []*Comment
}

func NewPost(threadID, userID, content string) *Post {
	return &Post{
		ThreadID:  threadID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
		Comments:  make([]*Comment, 0),
	}
}

func (p *Post) Validate() error {
	if p.ThreadID == "" {
		return ErrInvalidThreadID
	}
	if p.UserID == "" {
		return ErrInvalidUserID
	}
	if p.Content == "" {
		return ErrEmptyContent
	}
	return nil
}

func (p *Post) AddComment(comment *Comment) {
	p.Comments = append(p.Comments, comment)
}

func (p *Post) RemoveComment(commentID string) {
	for i, comment := range p.Comments {
		if comment.ID == commentID {
			p.Comments = append(p.Comments[:i], p.Comments[i+1:]...)
			return
		}
	}
}
