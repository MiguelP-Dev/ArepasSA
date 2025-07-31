package repositories

import (
	"ArepasSA/internal/models"
	"time"

	"gorm.io/gorm"
)

type ClientRepository struct {
	*BaseRepository
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{NewBaseRepository(db)}
}

// ... (métodos CRUD básicos existentes)

func (r *ClientRepository) AddComment(clientID uint, comment *models.ClientComment) error {
	comment.ClientID = clientID
	comment.CommentAt = time.Now()
	return r.db.Create(comment).Error
}

func (r *ClientRepository) GetComments(clientID uint, onlyPreferences bool) ([]models.ClientComment, error) {
	var comments []models.ClientComment
	query := r.db.Where("client_id = ?", clientID)

	if onlyPreferences {
		query = query.Where("is_preference = ?", true)
	}

	err := query.Order("comment_at DESC").Find(&comments).Error
	return comments, err
}
