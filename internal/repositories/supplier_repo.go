package repositories

import (
	"ArepasSA/internal/models"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	*BaseRepository
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{NewBaseRepository(db)}
}

func (r *SupplierRepository) Create(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *SupplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *SupplierRepository) Delete(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}

func (r *SupplierRepository) SoftDelete(id uint) error {
	return r.db.Model(&models.Supplier{}).Where("id = ?", id).Update("is_active", false).Error
}

func (r *SupplierRepository) FindByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *SupplierRepository) FindAll(activeOnly bool) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	query := r.db
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}
	err := query.Find(&suppliers).Error
	return suppliers, err
}

func (r *SupplierRepository) FindByProductID(productID uint) (*models.Supplier, error) {
	var product models.Product
	if err := r.db.Preload("Supplier").First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product.Supplier, nil
}
