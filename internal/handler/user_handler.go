package handler

import (
	"affluo/ent"
	"affluo/ent/user"
	"affluo/internal/dto"
	"affluo/internal/service"
	"net/mail"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate request
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Optional email validation
	if err := h.validateEmail(req.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email format",
		})
	}

	user, err := h.userService.CreateUser(c.Context(), req.Username, req.Email, req.Password, req.Role, req.FirstName, req.LastName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.NewUserResponse(user))
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	// Optional filtering
	role := c.Query("role")
	isActive := c.Query("is_active")

	var queryOpts []func(*ent.UserQuery)

	// Add role filter if provided
	if role != "" {
		queryOpts = append(queryOpts, func(q *ent.UserQuery) {
			q.Where(user.RoleEQ(user.Role(role)))
		})
	}

	// Add is_active filter if provided
	if isActive != "" {
		active, err := strconv.ParseBool(isActive)
		if err == nil {
			queryOpts = append(queryOpts, func(q *ent.UserQuery) {
				q.Where(user.IsActiveEQ(active))
			})
		}
	}

	users, err := h.userService.ListUsers(c.Context(), queryOpts...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users: " + err.Error(),
		})
	}

	// Convert users to response DTOs
	return c.JSON(dto.NewUsersResponse(users))
}
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetUserByID(c.Context(), idInt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found: " + err.Error(),
		})
	}

	return c.JSON(dto.NewUserResponse(user))
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate request
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Prepare updates map
	updates := make(map[string]interface{})

	// Optional email validation and update
	if req.Email != nil {
		if err := h.validateEmail(*req.Email); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid email format",
			})
		}
		updates["email"] = *req.Email
	}

	// Add other optional updates
	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}
	if req.Role != nil {
		updates["role"] = *req.Role
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	// Perform update if there are changes
	if len(updates) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No update fields provided",
		})
	}

	user, err := h.userService.UpdateUser(c.Context(), idInt, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user: " + err.Error(),
		})
	}

	return c.JSON(dto.NewUserResponse(user))
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Check if hard delete is requested via query parameter
	hardDelete := c.Query("hard_delete") == "true"

	if err := h.userService.DeleteUser(c.Context(), idInt, hardDelete); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user: " + err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Authentication handler
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate email format
	if err := h.validateEmail(loginReq.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email format",
		})
	}

	// Authenticate user
	user, err := h.userService.AuthenticateUser(c.Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication failed",
		})
	}

	// Return user response (without sensitive data)
	return c.JSON(dto.NewUserResponse(user))
}
