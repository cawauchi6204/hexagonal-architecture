package usecase

import (
	"context"
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type TagUseCase interface {
	CreateTag(ctx context.Context, name string) (*entity.Tag, error)
	GetTag(ctx context.Context, tagID string) (*entity.Tag, error)
	UpdateTag(ctx context.Context, tagID, name string) (*entity.Tag, error)
	DeleteTag(ctx context.Context, tagID string) error
	GetAllTags(ctx context.Context) ([]*entity.Tag, error)
	SearchTags(ctx context.Context, query string) ([]*entity.Tag, error)
	GetPopularTags(ctx context.Context, limit int) ([]*entity.Tag, error)
	GetThreadTags(ctx context.Context, threadID string) ([]*entity.Tag, error)
	GetUserTags(ctx context.Context, userID string) ([]*entity.Tag, error)
}

type tagUseCase struct {
	tagRepo repository.TagRepository
}

func NewTagUseCase(tagRepo repository.TagRepository) TagUseCase {
	return &tagUseCase{
		tagRepo: tagRepo,
	}
}

func (uc *tagUseCase) CreateTag(ctx context.Context, name string) (*entity.Tag, error) {
	tag := &entity.Tag{
		Name:      name,
		CreatedAt: time.Now(),
	}

	if err := tag.Validate(); err != nil {
		return nil, err
	}

	if err := uc.tagRepo.Create(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (uc *tagUseCase) GetTag(ctx context.Context, tagID string) (*entity.Tag, error) {
	return uc.tagRepo.FindByID(ctx, tagID)
}

func (uc *tagUseCase) UpdateTag(ctx context.Context, tagID, name string) (*entity.Tag, error) {
	tag, err := uc.tagRepo.FindByID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	if err := tag.UpdateName(name); err != nil {
		return nil, err
	}

	if err := uc.tagRepo.Update(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (uc *tagUseCase) DeleteTag(ctx context.Context, tagID string) error {
	return uc.tagRepo.Delete(ctx, tagID)
}

func (uc *tagUseCase) GetAllTags(ctx context.Context) ([]*entity.Tag, error) {
	return uc.tagRepo.FindAll(ctx)
}

func (uc *tagUseCase) SearchTags(ctx context.Context, query string) ([]*entity.Tag, error) {
	return uc.tagRepo.SearchByName(ctx, query)
}

func (uc *tagUseCase) GetPopularTags(ctx context.Context, limit int) ([]*entity.Tag, error) {
	return uc.tagRepo.FindPopular(ctx, limit)
}

func (uc *tagUseCase) GetThreadTags(ctx context.Context, threadID string) ([]*entity.Tag, error) {
	return uc.tagRepo.FindByThreadID(ctx, threadID)
}

func (uc *tagUseCase) GetUserTags(ctx context.Context, userID string) ([]*entity.Tag, error) {
	return uc.tagRepo.FindByUserID(ctx, userID)
}
