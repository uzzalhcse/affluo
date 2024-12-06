package routes

import (
	"affluo/ent"
	"affluo/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router, client *ent.Client) {
	userHandler := handler.NewUserHandler(client)

	router.Post("/users", userHandler.CreateUser)
	router.Get("/users", userHandler.ListUsers)
	router.Get("/users/:id", userHandler.GetUser)
	router.Put("/users/:id", userHandler.UpdateUser)
	router.Delete("/users/:id", userHandler.DeleteUser)
}
