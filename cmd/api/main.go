package main

import (
	"context"
	"financialcontrol/internal/v1/api"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("FINANCIAL_CONTROL_DATABASE_USER"),
		os.Getenv("FINANCIAL_CONTROL_DATABASE_PASSWORD"),
		os.Getenv("FINANCIAL_CONTROL_DATABASE_HOST"),
		os.Getenv("FINANCIAL_CONTROL_DATABASE_PORT"),
		os.Getenv("FINANCIAL_CONTROL_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.NewApi(gin.Default(), pool)

	api.RegisterRoutes()

	port := os.Getenv("FINANCIAL_CONTROL_APP_PORT")
	if port == "" {
		port = "3080"
	}

	fmt.Printf("Starting Server on port :%s\n", port)
	if err := api.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
