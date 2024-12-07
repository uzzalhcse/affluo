package main

import (
	"affluo/internal/app"
	"log"
)

func main() {
	application, err := app.NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer application.Cleanup()

	application.Run()
}
