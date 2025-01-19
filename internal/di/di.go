package di

import (
	"github.com/zeal2end/pratibimbh/internal/handlers"
	"github.com/zeal2end/pratibimbh/internal/repository"
	"github.com/zeal2end/pratibimbh/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler *handlers.UserHandler
}

func NewContainer(db *gorm.DB) *Container {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userSvc := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userSvc)

	return &Container{
		UserHandler: userHandler,
	}
}
