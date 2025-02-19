package base

import "time"

type Tag struct {
	ID           string
	Name         string
	CreatedAt    time.Time
	ThreadsCount int
}

func NewTag(name string) *Tag {
	return &Tag{
		Name:         name,
		CreatedAt:    time.Now(),
		ThreadsCount: 0,
	}
}

func (t *Tag) Validate() error {
	if t.Name == "" {
		return ErrEmptyTagName
	}
	return nil
}

func (t *Tag) IncrementThreadsCount() {
	t.ThreadsCount++
}

func (t *Tag) DecrementThreadsCount() {
	if t.ThreadsCount > 0 {
		t.ThreadsCount--
	}
}

func (t *Tag) UpdateName(newName string) error {
	if newName == "" {
		return ErrEmptyTagName
	}
	t.Name = newName
	return nil
}
