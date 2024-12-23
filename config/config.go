package config

import (
	"affluo/ent"
	"context"
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func InitDatabase() (*ent.Client, error) {
	// Database connection parameters
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create connection string
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, password, host, port, dbname)

	// Create Ent client
	client, err := ent.Open(dialect.Postgres, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Run migration
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, nil
}
