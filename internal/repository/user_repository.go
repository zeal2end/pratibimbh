package repository

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id int) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetUserByName(name string) (*User, error) {
	var user User
	err := r.db.Where("name = ?", name).First(&user).Error
	return &user, err
}
