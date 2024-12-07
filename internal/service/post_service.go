package service

import (
	"affluo/ent"
	"context"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	client *ent.Client
}

func NewPostService(client *ent.Client) *PostService {
	return &PostService{client: client}
}

func (s *PostService) CreatePost(ctx context.Context, title, content string, userID int64) (*ent.Post, error) {
	return s.client.Post.Create().
		SetID(uuid.New().String()).
		SetTitle(title).
		SetContent(content).
		SetCreatedAt(time.Now()).
		SetAuthorID(userID).
		Save(ctx)
}

func (s *PostService) GetPostByID(ctx context.Context, id string) (*ent.Post, error) {
	return s.client.Post.Get(ctx, id)
}

func (s *PostService) UpdatePost(ctx context.Context, id, title, content string) (*ent.Post, error) {
	return s.client.Post.UpdateOneID(id).
		SetTitle(title).
		SetContent(content).
		Save(ctx)
}

func (s *PostService) DeletePost(ctx context.Context, id string) error {
	return s.client.Post.DeleteOneID(id).Exec(ctx)
}
