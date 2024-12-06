package routes

import (
	"affluo/ent"
	"affluo/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(router fiber.Router, client *ent.Client) {
	postHandler := handler.NewPostHandler(client)

	router.Post("/posts", postHandler.CreatePost)
	router.Get("/posts/:id", postHandler.GetPost)
	router.Put("/posts/:id", postHandler.UpdatePost)
	router.Delete("/posts/:id", postHandler.DeletePost)
}
