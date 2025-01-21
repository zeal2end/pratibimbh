package service

import (
	"errors"

	"github.com/zeal2end/pratibimbh/internal/repository"
)

type BlogService struct {
	repo *repository.BlogRepository
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreateBlog(blog *repository.Blog) error {
	return s.repo.CreateBlog(blog)
}

func (s *BlogService) GetBlogByID(id int) (*repository.Blog, error) {
	return s.repo.GetBlogByID(id)
}

func (s *BlogService) ListBlogs() ([]repository.Blog, error) {
	return s.repo.ListBlogs()
}

func (s *BlogService) GetBlogsByUserID(userID int) ([]repository.Blog, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetBlogsByUserID(userID)
}

func (s *BlogService) UpdateBlog(blog *repository.Blog, userID int) error {
	existingBlog, err := s.repo.GetBlogByID(blog.ID)
	if err != nil {
		return err
	}

	// Check if user has permission to update
	if existingBlog.UserID != userID {
		return errors.New("unauthorized to update this blog")
	}

	return s.repo.UpdateBlog(blog)
}

func (s *BlogService) DeleteBlog(blogID, userID int) error {
	existingBlog, err := s.repo.GetBlogByID(blogID)
	if err != nil {
		return err
	}

	// Check if user has permission to delete
	if existingBlog.UserID != userID {
		return errors.New("unauthorized to delete this blog")
	}

	return s.repo.DeleteBlog(blogID)
}

func (s *BlogService) GetBlogsByTopic(topicID int) ([]repository.Blog, error) {
	if topicID <= 0 {
		return nil, errors.New("invalid topic ID")
	}
	return s.repo.GetBlogsByTopic(topicID)
}

func (s *BlogService) GetBlogsByTag(tagID int) ([]repository.Blog, error) {
	if tagID <= 0 {
		return nil, errors.New("invalid tag ID")
	}
	return s.repo.GetBlogsByTag(tagID)
}

func (s *BlogService) AddTopicToBlog(blogID, topicID, userID int) error {
	blog, err := s.repo.GetBlogByID(blogID)
	if err != nil {
		return err
	}

	if blog.UserID != userID {
		return errors.New("unauthorized to modify this blog")
	}

	return s.repo.AddTopicToBlog(blogID, topicID)
}

func (s *BlogService) AddTagToBlog(blogID, tagID, userID int) error {
	blog, err := s.repo.GetBlogByID(blogID)
	if err != nil {
		return err
	}

	if blog.UserID != userID {
		return errors.New("unauthorized to modify this blog")
	}

	return s.repo.AddTagToBlog(blogID, tagID)
}
