package repository

import (
	"gorm.io/gorm"
)

type Blog struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	UserID  int     `json:"user_id"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Topics  []Topic `json:"topics" gorm:"many2many:blog_topics;"`
	Tags    []Tag   `json:"tags" gorm:"many2many:blog_tags;"`
}

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) CreateBlog(blog *Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) GetBlogByID(id int) (*Blog, error) {
	var blog Blog
	err := r.db.Preload("User").Preload("Topics").Preload("Tags").First(&blog, id).Error
	return &blog, err
}

func (r *BlogRepository) ListBlogs() ([]Blog, error) {
	var blogs []Blog
	err := r.db.Preload("User").Preload("Topics").Preload("Tags").Find(&blogs).Error
	return blogs, err
}
