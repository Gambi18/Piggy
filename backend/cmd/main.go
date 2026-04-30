package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/handlers"
	"piggy.com/internal/piggyservice"
)

func main() {
	route := gin.Default()

	// Configure Cors
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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
	dbUrl := "postgres://admin:2323@localhost:5433/piggy?sslmode=disable"
	dbConn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established!")
	repostory := repo.NewRepository(dbConn)
	if err := repo.MigrateUp(dbUrl, "./internal/db/migrations", zerolog.Nop().With().Logger()); err != nil {
		panic(err)
	}

	// Initialize service
	appService := piggyservice.NewService(repostory)
	handlers := handlers.NewHandler(appService)

	// Define application endpoints
	route.POST("/api/v1/transactions", handlers.CreateTransaction)
	route.GET("/api/v1/transactions", handlers.GetTransactions) // Run application
	route.POST("/api/v1/signup", handlers.SignUp)
	route.POST("/api/v1/signin", handlers.Login)
	fmt.Println("Server running on port 8081")
	route.Run(":8081")
}