package usecase

import (
	"context"
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, threadID, userID, content string) (*entity.Post, error)
	GetPost(ctx context.Context, postID string) (*entity.Post, error)
	UpdatePost(ctx context.Context, postID, content string) (*entity.Post, error)
	DeletePost(ctx context.Context, postID string) error
	GetThreadPosts(ctx context.Context, threadID string, offset, limit int) ([]*entity.Post, error)
	GetUserPosts(ctx context.Context, userID string, offset, limit int) ([]*entity.Post, error)
	AddComment(ctx context.Context, postID, userID, content string) (*entity.Comment, error)
	GetComments(ctx context.Context, postID string, offset, limit int) ([]*entity.Comment, error)
	SearchPosts(ctx context.Context, query string) ([]*entity.Post, error)
}

type postUseCase struct {
	postRepo    repository.PostRepository
	commentRepo repository.CommentRepository
}

func NewPostUseCase(postRepo repository.PostRepository, commentRepo repository.CommentRepository) PostUseCase {
	return &postUseCase{
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (uc *postUseCase) CreatePost(ctx context.Context, threadID, userID, content string) (*entity.Post, error) {
	post := &entity.Post{
		ThreadID:  threadID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	if err := uc.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUseCase) GetPost(ctx context.Context, postID string) (*entity.Post, error) {
	post, err := uc.postRepo.FindByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// コメントを取得
	comments, err := uc.postRepo.FindComments(ctx, postID)
	if err != nil {
		return nil, err
	}
	post.Comments = comments

	return post, nil
}

func (uc *postUseCase) UpdatePost(ctx context.Context, postID, content string) (*entity.Post, error) {
	post, err := uc.postRepo.FindByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	post.Content = content

	if err := post.Validate(); err != nil {
		return nil, err
	}

	if err := uc.postRepo.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUseCase) DeletePost(ctx context.Context, postID string) error {
	return uc.postRepo.Delete(ctx, postID)
}

func (uc *postUseCase) GetThreadPosts(ctx context.Context, threadID string, offset, limit int) ([]*entity.Post, error) {
	return uc.postRepo.FindByThreadIDWithPagination(ctx, threadID, offset, limit)
}

func (uc *postUseCase) GetUserPosts(ctx context.Context, userID string, offset, limit int) ([]*entity.Post, error) {
	return uc.postRepo.FindByUserIDWithPagination(ctx, userID, offset, limit)
}

func (uc *postUseCase) AddComment(ctx context.Context, postID, userID, content string) (*entity.Comment, error) {
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

func (uc *postUseCase) GetComments(ctx context.Context, postID string, offset, limit int) ([]*entity.Comment, error) {
	return uc.commentRepo.FindByPostIDWithPagination(ctx, postID, offset, limit)
}

func (uc *postUseCase) SearchPosts(ctx context.Context, query string) ([]*entity.Post, error) {
	// TODO: 実装する検索ロジック
	return nil, nil
}
