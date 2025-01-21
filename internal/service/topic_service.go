package service

import "github.com/zeal2end/pratibimbh/internal/repository"

type TopicService struct {
	repo *repository.TopicRepository
}

func NewTopicService(repo *repository.TopicRepository) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) CreateTopic(topic *repository.Topic) error {
	return s.repo.CreateTopic(topic)
}

func (s *TopicService) GetTopicByID(id int) (*repository.Topic, error) {
	return s.repo.GetTopicByID(id)
}

func (s *TopicService) ListTopics() ([]repository.Topic, error) {
	return s.repo.ListTopics()
}
