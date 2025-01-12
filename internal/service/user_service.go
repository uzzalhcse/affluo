package service

import (
	"affluo/ent"
	"affluo/ent/commissionplan"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
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
	query := s.client.User.Query().WithCommissionPlan()
	for _, opt := range opts {
		opt(query)
	}
	return query.All(ctx)
}

// CreateUser creates a new user with hashed password
func (s *UserService) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
	role string,
	firstName *string,
	lastName *string,
	commissionPlanID *int,
) (*ent.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Start a new transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	// Ensure proper rollback in case of error
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Create user builder
	userBuilder := tx.User.
		Create().
		SetUsername(username).
		SetEmail(email).
		SetPasswordHash(string(hashedPassword)).
		SetRole(user.Role(role))

	// Set optional fields
	if firstName != nil {
		userBuilder.SetFirstName(*firstName)
	}
	if lastName != nil {
		userBuilder.SetLastName(*lastName)
	}

	// Handle commission plan if provided
	if commissionPlanID != nil {
		// Verify commission plan exists and is active
		plan, err := tx.CommissionPlan.
			Query().
			Where(
				commissionplan.ID(*commissionPlanID),
				commissionplan.IsActive(true),
			).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				tx.Rollback()
				return nil, errors.New("commission plan not found or inactive")
			}
			tx.Rollback()
			return nil, fmt.Errorf("querying commission plan: %w", err)
		}

		userBuilder.SetCommissionPlanID(plan.ID)
	} else {
		// If no commission plan provided, try to get the default one
		defaultPlan, err := tx.CommissionPlan.
			Query().
			Where(
				commissionplan.IsDefault(true),
				commissionplan.IsActive(true),
			).
			Only(ctx)

		if err == nil {
			userBuilder.SetCommissionPlanID(defaultPlan.ID)
		}
		// If no default plan, continue without setting a commission plan
	}

	// Create the user
	user, err := userBuilder.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("creating user: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.IDEQ(id)).
		Only(ctx)
}
func (s *UserService) ToggleActive(ctx context.Context, id int64) error {
	return s.client.User.UpdateOneID(id).
		SetIsActive(!s.client.User.Query().Where(user.IDEQ(id)).OnlyX(ctx).IsActive).
		Exec(ctx)
}

// GetUserByEmail retrieves a user by their email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}

// UpdateUser updates user details
func (s *UserService) UpdateUser(ctx context.Context, id int64, updates map[string]interface{}) error {
	// Start a new transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Get the update builder
	updateBuilder := tx.User.UpdateOneID(id)

	// Apply all updates
	for field, value := range updates {
		switch field {
		case "username":
			updateBuilder.SetUsername(value.(string))
		case "email":
			updateBuilder.SetEmail(value.(string))
		case "first_name":
			updateBuilder.SetFirstName(value.(string))
		case "last_name":
			updateBuilder.SetLastName(value.(string))
		case "role":
			updateBuilder.SetRole(user.Role(value.(string)))
		case "is_active":
			updateBuilder.SetIsActive(value.(bool))
		case "commission_plan_id":
			planID := value.(int)
			// Verify commission plan exists and is active
			plan, err := tx.CommissionPlan.
				Query().
				Where(
					commissionplan.ID(planID),
					commissionplan.IsActive(true),
				).
				Only(ctx)

			if err != nil {
				tx.Rollback()
				if ent.IsNotFound(err) {
					return errors.New("commission plan not found or inactive")
				}
				return fmt.Errorf("querying commission plan: %w", err)
			}
			updateBuilder.SetCommissionPlanID(plan.ID)
		}
	}

	// Perform the update
	_, err = updateBuilder.Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("updating user: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
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
