package repository

import (
	"gorm.io/gorm"
)

type Tag struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Blogs []Blog `json:"blogs" gorm:"many2many:blog_tags;"`
}

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) CreateTag(tag *Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) GetTagByID(id int) (*Tag, error) {
	var tag Tag
	err := r.db.Preload("Blogs").First(&tag, id).Error
	return &tag, err
}

func (r *TagRepository) ListTags() ([]Tag, error) {
	var tags []Tag
	err := r.db.Find(&tags).Error
	return tags, err
}
