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
	"github.com/zeal2end/pratibimbh/internal/handlers"
	"github.com/zeal2end/pratibimbh/internal/middleware"
)

func main() {
	cfg := config.Load()

	handlers.InitDB(cfg.DatabaseURL)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/user/:id", handlers.GetUser)
		api.GET("/user", handlers.GetUserByUsername)
		api.POST("/user", handlers.CreateUser)
	}

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
