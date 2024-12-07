package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *CreateUserRequest) Validate() error {
	return validate.Struct(r)
}

type UpdateUserRequest struct {
	Username *string `json:"username" validate:"omitempty,min=3,max=50"`
	Email    *string `json:"email" validate:"omitempty,email"`
}

func (r *UpdateUserRequest) Validate() error {
	return validate.Struct(r)
}
