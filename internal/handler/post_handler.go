package handler

import (
	"affluo/ent"
	"affluo/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(client *ent.Client) *PostHandler {
	return &PostHandler{
		postService: service.NewPostService(client),
	}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	type CreatePostRequest struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var req CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	post, err := h.postService.CreatePost(c.Context(), req.UserID, req.Title, req.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func (h *PostHandler) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := h.postService.GetPostByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.JSON(post)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	type UpdatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var req UpdatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	id := c.Params("id")
	post, err := h.postService.UpdatePost(c.Context(), id, req.Title, req.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(post)
}

func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.postService.DeletePost(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
