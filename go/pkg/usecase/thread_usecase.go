package usecase

import (
	"context"
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type ThreadUseCase interface {
	CreateThread(ctx context.Context, title string, userID string, tagIDs []string) (*entity.Thread, error)
	GetThread(ctx context.Context, threadID string) (*entity.Thread, error)
	UpdateThread(ctx context.Context, threadID string, title string) (*entity.Thread, error)
	DeleteThread(ctx context.Context, threadID string) error
	AddTag(ctx context.Context, threadID string, tagID string) error
	RemoveTag(ctx context.Context, threadID string, tagID string) error
	GetThreadsByTag(ctx context.Context, tagID string) ([]*entity.Thread, error)
	GetThreadsByUser(ctx context.Context, userID string) ([]*entity.Thread, error)
	GetLatestThreads(ctx context.Context, limit int) ([]*entity.Thread, error)
	GetPopularThreads(ctx context.Context, limit int) ([]*entity.Thread, error)
	SearchThreads(ctx context.Context, query string) ([]*entity.Thread, error)
}

type threadUseCase struct {
	threadRepo repository.ThreadRepository
	tagRepo    repository.TagRepository
}

func NewThreadUseCase(threadRepo repository.ThreadRepository, tagRepo repository.TagRepository) ThreadUseCase {
	return &threadUseCase{
		threadRepo: threadRepo,
		tagRepo:    tagRepo,
	}
}

func (uc *threadUseCase) CreateThread(ctx context.Context, title string, userID string, tagIDs []string) (*entity.Thread, error) {
	thread := &entity.Thread{
		Title:     title,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := thread.Validate(); err != nil {
		return nil, err
	}

	if err := uc.threadRepo.Create(ctx, thread); err != nil {
		return nil, err
	}

	// タグを追加
	for _, tagID := range tagIDs {
		if err := uc.threadRepo.AddTag(ctx, thread.ID, tagID); err != nil {
			return nil, err
		}
	}

	return thread, nil
}

func (uc *threadUseCase) GetThread(ctx context.Context, threadID string) (*entity.Thread, error) {
	thread, err := uc.threadRepo.FindByID(ctx, threadID)
	if err != nil {
		return nil, err
	}

	// 関連する投稿を取得
	posts, err := uc.threadRepo.FindPosts(ctx, threadID)
	if err != nil {
		return nil, err
	}
	thread.Posts = posts

	// 関連するタグを取得
	tags, err := uc.threadRepo.FindTags(ctx, threadID)
	if err != nil {
		return nil, err
	}
	thread.Tags = tags

	return thread, nil
}

func (uc *threadUseCase) UpdateThread(ctx context.Context, threadID string, title string) (*entity.Thread, error) {
	thread, err := uc.threadRepo.FindByID(ctx, threadID)
	if err != nil {
		return nil, err
	}

	thread.Title = title

	if err := thread.Validate(); err != nil {
		return nil, err
	}

	if err := uc.threadRepo.Update(ctx, thread); err != nil {
		return nil, err
	}

	return thread, nil
}

func (uc *threadUseCase) DeleteThread(ctx context.Context, threadID string) error {
	return uc.threadRepo.Delete(ctx, threadID)
}

func (uc *threadUseCase) AddTag(ctx context.Context, threadID string, tagID string) error {
	return uc.threadRepo.AddTag(ctx, threadID, tagID)
}

func (uc *threadUseCase) RemoveTag(ctx context.Context, threadID string, tagID string) error {
	return uc.threadRepo.RemoveTag(ctx, threadID, tagID)
}

func (uc *threadUseCase) GetThreadsByTag(ctx context.Context, tagID string) ([]*entity.Thread, error) {
	return uc.threadRepo.FindByTag(ctx, tagID)
}

func (uc *threadUseCase) GetThreadsByUser(ctx context.Context, userID string) ([]*entity.Thread, error) {
	return uc.threadRepo.FindByUserID(ctx, userID)
}

func (uc *threadUseCase) GetLatestThreads(ctx context.Context, limit int) ([]*entity.Thread, error) {
	return uc.threadRepo.FindLatest(ctx, limit)
}

func (uc *threadUseCase) GetPopularThreads(ctx context.Context, limit int) ([]*entity.Thread, error) {
	return uc.threadRepo.FindPopular(ctx, limit)
}

func (uc *threadUseCase) SearchThreads(ctx context.Context, query string) ([]*entity.Thread, error) {
	// TODO: 実装する検索ロジック
	return nil, nil
}
