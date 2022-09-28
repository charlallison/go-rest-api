package comment

import (
	"github.com/jinzhu/gorm"
)

// Service - comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - interface for comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, comment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - returns a comment by an ID
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// GetCommentsBySlug - retrieves comments by slug
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

// PostComment - saves a comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// UpdateComment - updates a comment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, nil
	}

	if result := s.DB.Model(&comment).Update(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// DeleteComment - deletes a comment from the database
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - retrieves all comments from the database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}
