package entity

import "time"

type Thread struct {
	ID        string
	Title     string
	UserID    string
	CreatedAt time.Time
	Posts     []*Post
	Tags      []*Tag
}

func NewThread(title string, userID string) *Thread {
	return &Thread{
		Title:     title,
		UserID:    userID,
		CreatedAt: time.Now(),
		Posts:     make([]*Post, 0),
		Tags:      make([]*Tag, 0),
	}
}

func (t *Thread) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	if t.UserID == "" {
		return ErrInvalidUserID
	}
	return nil
}

func (t *Thread) AddPost(post *Post) {
	t.Posts = append(t.Posts, post)
}

func (t *Thread) AddTag(tag *Tag) {
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
