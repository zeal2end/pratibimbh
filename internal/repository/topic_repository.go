package repository

import (
	"gorm.io/gorm"
)

type Topic struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Blogs []Blog `json:"blogs" gorm:"many2many:blog_topics;"`
}

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) CreateTopic(topic *Topic) error {
	return r.db.Create(topic).Error
}

func (r *TopicRepository) GetTopicByID(id int) (*Topic, error) {
	var topic Topic
	err := r.db.Preload("Blogs").First(&topic, id).Error
	return &topic, err
}

func (r *TopicRepository) ListTopics() ([]Topic, error) {
	var topics []Topic
	err := r.db.Find(&topics).Error
	return topics, err
}
