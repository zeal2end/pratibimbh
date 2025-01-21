package repository

import (
	"gorm.io/gorm"
)

type Comment struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	BlogID  int    `json:"blog_id"`
	UserID  int    `json:"user_id"`
	Blog    Blog   `json:"blog" gorm:"foreignKey:BlogID"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetCommentsByBlogID(blogID int) ([]Comment, error) {
	var comments []Comment
	err := r.db.Where("blog_id = ?", blogID).Preload("User").Find(&comments).Error
	return comments, err
}
