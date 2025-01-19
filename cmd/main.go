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
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  if err := db.AutoMigrate(&handlers.User{}); err != nil {
    log.Fatalf("AutoMigrate failed: %v", err)
  }

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/user/:id", func(c *gin.Context) {
    	handlers.GetUser(c, db)
    })
    api.POST("/user", func(c *gin.Context) {
    	handlers.CreateUser(c, db)
    })
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
