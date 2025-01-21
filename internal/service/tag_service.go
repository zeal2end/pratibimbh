package service

import "github.com/zeal2end/pratibimbh/internal/repository"

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) CreateTag(tag *repository.Tag) error {
	return s.repo.CreateTag(tag)
}

func (s *TagService) GetTagByID(id int) (*repository.Tag, error) {
	return s.repo.GetTagByID(id)
}

func (s *TagService) ListTags() ([]repository.Tag, error) {
	return s.repo.ListTags()
}
