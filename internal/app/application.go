package app

import (
	"affluo/config"
	"affluo/ent"
	"affluo/internal/handler"
	"affluo/internal/middleware"
	"affluo/internal/service"
	"context"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	Config   *config.Config
	DB       *ent.Client
	Redis    *redis.Client
	App      *fiber.App
	Services *service.Container
	Handlers *handler.Container
}

func NewApplication() (*Application, error) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize dependencies
	dbClient, err := setupDatabase(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	redisClient := setupRedis(cfg.RedisURL)

	// Initialize services and handlers
	services := service.NewContainer(dbClient, redisClient)
	handlers := handler.NewContainer(services)

	// Initialize Fiber app
	fiberApp := fiber.New()
	setupMiddleware(fiberApp)

	application := &Application{
		Config:   cfg,
		DB:       dbClient,
		Redis:    redisClient,
		App:      fiberApp,
		Services: services,
		Handlers: handlers,
	}

	application.SetupRoutes()
	return application, nil
}

func (a *Application) Run() {
	go a.StartBackgroundJobs()

	serverAddress := fmt.Sprintf(":%d", a.Config.Port)
	if err := a.App.Listen(serverAddress); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}
}

func (a *Application) Cleanup() {
	if err := a.DB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
	if err := a.Redis.Close(); err != nil {
		log.Printf("Error closing Redis: %v", err)
	}
}

func (a *Application) StartBackgroundJobs() {
	ctx := context.Background()
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		if err := a.Services.Reporting.GenerateScheduledReports(ctx); err != nil {
			log.Printf("Error generating reports: %v", err)
		}
	}
}

func setupMiddleware(app *fiber.App) {
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
}

func setupDatabase(databaseURL string) (*ent.Client, error) {
	client, err := ent.Open(dialect.Postgres, databaseURL)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema: %v", err)
	}
	return client, nil
}

func setupRedis(redisURL string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: redisURL})
}