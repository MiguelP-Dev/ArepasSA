package repositories

import (
	"ArepasSA/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{NewBaseRepository(db)}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) SoftDelete(id uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("is_active", false).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindAll(activeOnly bool) ([]models.User, error) {
	var users []models.User
	query := r.db
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}
	err := query.Find(&users).Error
	return users, err
}
