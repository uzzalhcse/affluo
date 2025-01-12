package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateUserRequest struct {
	Username         string  `json:"username" validate:"required,min=3,max=50"`
	Email            string  `json:"email" validate:"required,email"`
	Password         string  `json:"password" validate:"required,min=6"`
	Role             string  `json:"role" validate:"required,oneof=admin affiliate manager"`
	FirstName        *string `json:"first_name,omitempty" validate:"omitempty,max=50"`
	LastName         *string `json:"last_name,omitempty" validate:"omitempty,max=50"`
	CommissionPlanID *int    `json:"commission_plan_id,omitempty" validate:"omitempty,gt=0"`
}

func (r *CreateUserRequest) Validate() error {
	return validate.Struct(r)
}

type UpdateUserRequest struct {
	Username         *string `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Email            *string `json:"email,omitempty" validate:"omitempty,email"`
	FirstName        *string `json:"first_name,omitempty" validate:"omitempty,max=50"`
	LastName         *string `json:"last_name,omitempty" validate:"omitempty,max=50"`
	Role             *string `json:"role,omitempty" validate:"omitempty,oneof=admin affiliate manager"`
	IsActive         *bool   `json:"is_active,omitempty"`
	CommissionPlanID *int    `json:"commission_plan_id,omitempty" validate:"omitempty,gt=0"`
}

func (r *UpdateUserRequest) Validate() error {
	return validate.Struct(r)
}
