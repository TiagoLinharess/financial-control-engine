package main

import (
	"context"
	"financialcontrol/internal/constants"
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
		os.Getenv(constants.EnvDBUser),
		os.Getenv(constants.EnvDBPassword),
		os.Getenv(constants.EnvDBHost),
		os.Getenv(constants.EnvDBPort),
		os.Getenv(constants.EnvDBName),
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

	port := os.Getenv(constants.EnvAppPort)
	if port == "" {
		port = constants.DefaultAppPort
	}

	fmt.Printf("Starting Server on port :%s\n", port)
	if err := api.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
