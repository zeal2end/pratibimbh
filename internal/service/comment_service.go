package service

import "github.com/zeal2end/pratibimbh/internal/repository"

type CommentService struct {
	repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *repository.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetCommentsByBlogID(blogID int) ([]repository.Comment, error) {
	return s.repo.GetCommentsByBlogID(blogID)
}
