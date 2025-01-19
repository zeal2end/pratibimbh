package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/zeal2end/pratibimbh/internal/config"
	"github.com/zeal2end/pratibimbh/internal/database"
	"github.com/zeal2end/pratibimbh/internal/di"
	"github.com/zeal2end/pratibimbh/internal/middleware"
)

func main() {
	cfg := config.Load()

	database.InitDB(cfg.DatabaseURL)

	container := di.NewContainer(database.GetDB())

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api")
	setupRoutes(api, container)

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func setupRoutes(api *gin.RouterGroup, container *di.Container) {
	userGroup := api.Group("/user")
	{
		userGroup.GET("/:id", container.UserHandler.GetUserByID)
		userGroup.GET("", container.UserHandler.GetUserByName)
		userGroup.POST("", container.UserHandler.CreateUser)
	}

	// Add more route groups here if needed
}
