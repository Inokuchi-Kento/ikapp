package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/Inokuchi-Kento/ikapp/config"
	"github.com/Inokuchi-Kento/ikapp/ent"
	_ "github.com/lib/pq"
)

func OpenDB(ctx context.Context, cfg *config.Config) *ent.Client {
	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	))
	if err != nil {
		log.Fatalf("failed connecting to Postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
