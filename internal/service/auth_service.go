// internal/service/auth_service.go
package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"affluo/ent"
	"affluo/ent/user"
	"affluo/internal/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	client       *ent.Client
	redisClient  *redis.Client
	jwtSecretKey []byte
}

func NewAuthService(client *ent.Client, redisClient *redis.Client, jwtSecretKey string) *AuthService {
	return &AuthService{
		client:       client,
		redisClient:  redisClient,
		jwtSecretKey: []byte(jwtSecretKey),
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*ent.User, error) {
	// Check if user already exists
	exists, err := s.client.User.Query().
		Where(user.EmailEQ(req.Email)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser, err := s.client.User.Create().
		SetEmail(req.Email).
		SetPasswordHash(string(hashedPassword)).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetUsername(req.Username).
		Save(ctx)

	return newUser, err
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	// Find user by email
	existingUser, err := s.client.User.Query().
		Where(user.UsernameEQ(req.Username)).
		First(ctx)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, expiresAt, err := s.generateJWT(existingUser)
	if err != nil {
		return nil, err
	}

	// Update last login
	_, err = existingUser.Update().
		SetLastLogin(time.Now()).
		Save(ctx)

	return &dto.AuthResponse{
		Token:     token,
		Type:      "Bearer",
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GetUserFromContext(ctx context.Context, userID int64) (*ent.User, error) {

	userInfo, err := s.client.User.Query().Where(user.IDEQ(userID)).First(ctx)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return userInfo, nil

}
func (s *AuthService) generateJWT(u *ent.User) (string, int64, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"user_id": u.ID,
		"email":   u.Email,
		"role":    u.Role,
		"exp":     expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecretKey)

	return tokenString, expiresAt, err
}

func (s *AuthService) ValidateJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return s.jwtSecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, claims, nil
}

func (s *AuthService) InitiatePasswordReset(ctx context.Context, email string) error {
	// Find user by email
	existingUser, err := s.client.User.Query().
		Where(user.EmailEQ(email)).
		First(ctx)
	if err != nil {
		return errors.New("user not found")
	}

	// Generate reset token
	resetToken := uuid.New().String()

	// Store reset token with expiration in Redis
	err = s.redisClient.Set(ctx, "reset_token:"+resetToken, existingUser.ID, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	// Update user with reset token (optional, depending on your design)
	_, err = existingUser.Update().
		SetResetToken(resetToken).
		SetResetTokenExpiresAt(time.Now().Add(1 * time.Hour)).
		Save(ctx)

	return err
}

func (s *AuthService) ResetPassword(ctx context.Context, token, newPassword string) error {
	// Validate reset token from Redis
	userID, err := s.redisClient.Get(ctx, "reset_token:"+token).Result()
	if err != nil {
		return errors.New("invalid or expired reset token")
	}

	// Convert userID from string to int64
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return errors.New("invalid user ID")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update user password
	_, err = s.client.User.Update().
		Where(user.IDEQ(userIDInt)).
		SetPasswordHash(string(hashedPassword)).
		SetResetToken("").
		SetResetTokenExpiresAt(time.Time{}).
		Save(ctx)

	// Remove reset token from Redis
	s.redisClient.Del(ctx, "reset_token:"+token)

	return err
}
