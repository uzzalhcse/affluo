package service

import (
	"affluo/ent"
	"affluo/ent/user"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}

// ListUsers retrieves all users with optional filtering
func (s *UserService) ListUsers(ctx context.Context, opts ...func(*ent.UserQuery)) ([]*ent.User, error) {
	query := s.client.User.Query()
	for _, opt := range opts {
		opt(query)
	}
	return query.All(ctx)
}

// CreateUser creates a new user with hashed password
func (s *UserService) CreateUser(ctx context.Context, username, email, password, role string, firstName, lastName *string) (*ent.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user with optional first and last name
	userCreate := s.client.User.Create().
		SetUsername(username).
		SetEmail(email).
		SetPasswordHash(string(hashedPassword)).
		SetCreatedAt(time.Now()).
		SetRole(user.Role(role)).
		SetIsActive(true)

	if firstName != nil {
		userCreate.SetFirstName(*firstName)
	}

	if lastName != nil {
		userCreate.SetLastName(*lastName)
	}

	return userCreate.Save(ctx)
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.IDEQ(id)).
		Only(ctx)
}

// GetUserByEmail retrieves a user by their email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}

// UpdateUser updates user details
func (s *UserService) UpdateUser(ctx context.Context, id int64, updates map[string]interface{}) (*ent.User, error) {
	update := s.client.User.UpdateOneID(id)

	if username, ok := updates["username"].(string); ok {
		update.SetUsername(username)
	}

	if email, ok := updates["email"].(string); ok {
		update.SetEmail(email)
	}

	if firstName, ok := updates["first_name"].(string); ok {
		update.SetFirstName(firstName)
	}

	if lastName, ok := updates["last_name"].(string); ok {
		update.SetLastName(lastName)
	}

	if role, ok := updates["role"].(string); ok {
		update.SetRole(user.Role(role))
	}

	if isActive, ok := updates["is_active"].(bool); ok {
		update.SetIsActive(isActive)
	}

	return update.Save(ctx)
}

// UpdateUserPassword updates a user's password
func (s *UserService) UpdateUserPassword(ctx context.Context, id int64, newPassword string) error {
	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.client.User.UpdateOneID(id).
		SetPasswordHash(string(hashedPassword)).
		Save(ctx)

	return err
}

// DeleteUser soft deletes or hard deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int64, hardDelete bool) error {
	if hardDelete {
		return s.client.User.DeleteOneID(id).Exec(ctx)
	}

	// Soft delete by setting is_active to false
	_, err := s.client.User.UpdateOneID(id).
		SetIsActive(false).
		Save(ctx)
	return err
}

// AuthenticateUser validates user credentials
func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*ent.User, error) {
	// Retrieve user by email
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
