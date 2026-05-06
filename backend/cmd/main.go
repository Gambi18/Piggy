package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/handlers"
	"piggy.com/internal/piggyservice"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: error loading .env file: %v\n", err)
	}

	route := gin.Default()

	// Configure Cors
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000" // fallback for local development
	}
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL, "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Healthcheck
	route.GET("/api/v1/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Healthy!",
		})
	})

	// Initialize repo and apply migrations
	ctx := context.Background()

	// Check for Render's DATABASE_URL first, then fall back to individual env vars
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	dbConn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Database connection established to: %s\n", dbUrl)
	repository := repo.NewRepository(dbConn)
	if err := repo.MigrateUp(dbUrl, "./internal/db/migrations", zerolog.Nop().With().Logger()); err != nil {
		panic(err)
	}

	// Initialize service
	appService := piggyservice.NewService(repository)
	handlers := handlers.NewHandler(appService)

	// Define application endpoints
	route.POST("/api/v1/transactions", handlers.CreateTransaction)
	route.GET("/api/v1/transactions", handlers.GetTransactions)
	route.POST("/api/v1/signup", handlers.SignUp)
	route.POST("/api/v1/login", handlers.Login)
	route.GET("/api/v1/balance", handlers.GetBalance)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // fallback for local development
	}
	fmt.Printf("Server running on port %s\n", port)
	route.Run(":" + port)
}
