package repositories

import (
	"ArepasSA/internal/models"

	"gorm.io/gorm"
)

type ComboRepository struct {
	*BaseRepository
}

func NewComboRepository(db *gorm.DB) *ComboRepository {
	return &ComboRepository{NewBaseRepository(db)}
}

func (r *ComboRepository) Create(combo *models.Combo) error {
	return r.db.Create(combo).Error
}

func (r *ComboRepository) Update(combo *models.Combo) error {
	return r.db.Save(combo).Error
}

func (r *ComboRepository) Delete(id uint) error {
	return r.db.Delete(&models.Combo{}, id).Error
}

func (r *ComboRepository) FindByID(id uint) (*models.Combo, error) {
	var combo models.Combo
	err := r.db.Preload("Items").Preload("Items.Product").First(&combo, id).Error
	return &combo, err
}

func (r *ComboRepository) FindAll(activeOnly bool) ([]models.Combo, error) {
	var combos []models.Combo
	query := r.db.Preload("Items").Preload("Items.Product")

	if activeOnly {
		query = query.Where("is_active = ?", true)
	}

	err := query.Find(&combos).Error
	return combos, err
}

func (r *ComboRepository) CheckComboStock(comboID uint, partial bool) (bool, error) {
	var items []models.ComboItem
	if err := r.db.Where("combo_id = ?", comboID).Find(&items).Error; err != nil {
		return false, err
	}

	productRepo := NewProductRepository(r.db)
	for _, item := range items {
		requiredQuantity := item.Quantity
		if partial {
			requiredQuantity = item.Quantity / 2
		}

		hasStock, err := productRepo.CheckStock(item.ProductID, requiredQuantity)
		if err != nil || !hasStock {
			return false, err
		}
	}

	return true, nil
}

func (r *ComboRepository) UpdateComboStock(comboID uint, partial bool) error {
	var items []models.ComboItem
	if err := r.db.Where("combo_id = ?", comboID).Find(&items).Error; err != nil {
		return err
	}

	productRepo := NewProductRepository(r.db)
	for _, item := range items {
		quantity := item.Quantity
		if partial {
			quantity = item.Quantity / 2
		}

		if err := productRepo.UpdateStock(item.ProductID, -quantity); err != nil {
			return err
		}
	}

	return nil
}
