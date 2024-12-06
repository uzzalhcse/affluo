package service

import (
	"affluo/ent"
	"affluo/ent/user"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}
func (s *UserService) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return s.client.User.Query().All(ctx)
}
func (s *UserService) CreateUser(ctx context.Context, username, email, password string) (*ent.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	return s.client.User.Create().
		SetID(uuid.New().String()).
		SetUsername(username).
		SetEmail(email).
		SetPasswordHash(string(hashedPassword)).
		SetCreatedAt(time.Now()).
		Save(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.IDEQ(id)).
		Only(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, username, email *string) (*ent.User, error) {
	update := s.client.User.UpdateOneID(id)

	if username != nil {
		update.SetUsername(*username)
	}

	if email != nil {
		update.SetEmail(*email)
	}

	return update.Save(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.client.User.DeleteOneID(id).Exec(ctx)
}
