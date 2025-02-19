package thread

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity/base"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity/post"
)

type Thread struct {
	ID        string
	Title     string
	UserID    string
	CreatedAt time.Time
	Posts     []*post.Post
	Tags      []*base.Tag
}

func NewThread(title string, userID string) *Thread {
	return &Thread{
		Title:     title,
		UserID:    userID,
		CreatedAt: time.Now(),
		Posts:     make([]*post.Post, 0),
		Tags:      make([]*base.Tag, 0),
	}
}

func (t *Thread) Validate() error {
	if t.Title == "" {
		return base.ErrEmptyTitle
	}
	if t.UserID == "" {
		return base.ErrInvalidUserID
	}
	return nil
}

func (t *Thread) AddPost(post *post.Post) {
	t.Posts = append(t.Posts, post)
}

func (t *Thread) AddTag(tag *base.Tag) {
	t.Tags = append(t.Tags, tag)
}

func (t *Thread) RemoveTag(tagID string) {
	for i, tag := range t.Tags {
		if tag.ID == tagID {
			t.Tags = append(t.Tags[:i], t.Tags[i+1:]...)
			return
		}
	}
}
