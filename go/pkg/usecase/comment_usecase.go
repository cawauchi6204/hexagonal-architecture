package usecase

import (
	"context"
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type CommentUseCase interface {
	CreateComment(ctx context.Context, postID, userID, content string) (*entity.Comment, error)
	GetComment(ctx context.Context, commentID string) (*entity.Comment, error)
	UpdateComment(ctx context.Context, commentID, content string) (*entity.Comment, error)
	DeleteComment(ctx context.Context, commentID string) error
	AddLike(ctx context.Context, commentID, userID string) error
	RemoveLike(ctx context.Context, commentID, userID string) error
	GetLikes(ctx context.Context, commentID string) (int, error)
	GetUserComments(ctx context.Context, userID string) ([]*entity.Comment, error)
	GetPostComments(ctx context.Context, postID string, offset, limit int) ([]*entity.Comment, error)
}

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
	}
}

func (uc *commentUseCase) CreateComment(ctx context.Context, postID, userID, content string) (*entity.Comment, error) {
	comment := &entity.Comment{
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := comment.Validate(); err != nil {
		return nil, err
	}

	if err := uc.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (uc *commentUseCase) GetComment(ctx context.Context, commentID string) (*entity.Comment, error) {
	return uc.commentRepo.FindByID(ctx, commentID)
}

func (uc *commentUseCase) UpdateComment(ctx context.Context, commentID, content string) (*entity.Comment, error) {
	comment, err := uc.commentRepo.FindByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	comment.Content = content

	if err := comment.Validate(); err != nil {
		return nil, err
	}

	if err := uc.commentRepo.Update(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (uc *commentUseCase) DeleteComment(ctx context.Context, commentID string) error {
	return uc.commentRepo.Delete(ctx, commentID)
}

func (uc *commentUseCase) AddLike(ctx context.Context, commentID, userID string) error {
	return uc.commentRepo.AddLike(ctx, commentID, userID)
}

func (uc *commentUseCase) RemoveLike(ctx context.Context, commentID, userID string) error {
	return uc.commentRepo.RemoveLike(ctx, commentID, userID)
}

func (uc *commentUseCase) GetLikes(ctx context.Context, commentID string) (int, error) {
	return uc.commentRepo.CountLikes(ctx, commentID)
}

func (uc *commentUseCase) GetUserComments(ctx context.Context, userID string) ([]*entity.Comment, error) {
	return uc.commentRepo.FindByUserID(ctx, userID)
}

func (uc *commentUseCase) GetPostComments(ctx context.Context, postID string, offset, limit int) ([]*entity.Comment, error) {
	return uc.commentRepo.FindByPostIDWithPagination(ctx, postID, offset, limit)
}
