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

func (r *BlogRepository) GetBlogsByUserID(userID int) ([]Blog, error) {
	var blogs []Blog
	err := r.db.Where("user_id = ?", userID).
		Preload("User").
		Preload("Topics").
		Preload("Tags").
		Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) UpdateBlog(blog *Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) DeleteBlog(blogID int) error {
	return r.db.Delete(&Blog{}, blogID).Error
}

func (r *BlogRepository) GetBlogsByTopic(topicID int) ([]Blog, error) {
	var blogs []Blog
	err := r.db.Joins("JOIN blog_topics ON blogs.id = blog_topics.blog_id").
		Where("blog_topics.topic_id = ?", topicID).
		Preload("User").
		Preload("Topics").
		Preload("Tags").
		Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) GetBlogsByTag(tagID int) ([]Blog, error) {
	var blogs []Blog
	err := r.db.Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagID).
		Preload("User").
		Preload("Topics").
		Preload("Tags").
		Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) AddTopicToBlog(blogID, topicID int) error {
	return r.db.Exec("INSERT INTO blog_topics (blog_id, topic_id) VALUES (?, ?)", blogID, topicID).Error
}

func (r *BlogRepository) AddTagToBlog(blogID, tagID int) error {
	return r.db.Exec("INSERT INTO blog_tags (blog_id, tag_id) VALUES (?, ?)", blogID, tagID).Error
}
