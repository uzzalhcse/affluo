// internal/dto/auth_request.go
package dto

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required,max=50"`
	LastName  string `json:"last_name" validate:"required,max=50"`
	Username  string `json:"username" validate:"required,max=50"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required"`
}

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordResetConfirmRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
