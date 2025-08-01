package repositories

import (
	"sales-app/internal/models"
	"time"

	"gorm.io/gorm"
)

type ClientRepository struct {
	*BaseRepository
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{NewBaseRepository(db)}
}

func (r *ClientRepository) Create(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *ClientRepository) Update(client *models.Client) error {
	return r.db.Save(client).Error
}

func (r *ClientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Client{}, id).Error
}

func (r *ClientRepository) SoftDelete(id uint) error {
	return r.db.Model(&models.Client{}).Where("id = ?", id).Update("is_active", false).Error
}

func (r *ClientRepository) FindByID(id uint) (*models.Client, error) {
	var client models.Client
	err := r.db.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("client_comments.comment_date DESC")
	}).First(&client, id).Error
	return &client, err
}

func (r *ClientRepository) FindAll(activeOnly bool, search string) ([]models.Client, error) {
	var clients []models.Client
	query := r.db.Preload("Comments")
	
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}
	
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	
	err := query.Order("name ASC").Find(&clients).Error
	return clients, err
}

func (r *ClientRepository) AddComment(comment *models.ClientComment) error {
	comment.CommentDate = time.Now()
	return r.db.Create(comment).Error
}

func (r *ClientRepository) GetComments(clientID uint, commentType string) ([]models.ClientComment, error) {
	var comments []models.ClientComment
	query := r.db.Where("client_id = ?", clientID)
	
	if commentType != "" {
		query = query.Where("comment_type = ?", commentType)
	}
	
	err := query.Order("comment_date DESC").Find(&comments).Error
	return comments, err
}

func (r *ClientRepository) GetClientsWithPreferences() ([]models.Client, error) {
	var clients []models.Client
	err := r.db.Joins("JOIN client_comments ON clients.id = client_comments.client_id").
		Where("client_comments.comment_type = ?", "preference").
		Group("clients.id").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Where("comment_type = ?", "preference").Order("client_comments.comment_date DESC")
		}).
		Find(&clients).Error
	return clients, err
}
