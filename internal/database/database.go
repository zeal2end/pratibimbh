package database

import (
	"log"

	"github.com/zeal2end/pratibimbh/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	migrateEntities()
}

func migrateEntities() {
	// Add all entities that need to be migrated here
	// Order matters for foreign key relationships
	DB.AutoMigrate(
		&repository.User{},
		&repository.Topic{},
		&repository.Tag{},
		&repository.Blog{},
		&repository.Comment{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
