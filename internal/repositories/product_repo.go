package repositories

import (
	"ArepasSA/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	*BaseRepository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{NewBaseRepository(db)}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) SoftDelete(id uint) error {
	return r.db.Model(&models.Product{}).Where("id = ?",
		id).Update("is_active", false).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindAll(activeOnly bool) ([]models.Product, error) {
	var products []models.Product
	query := r.db
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}
	err := query.Find(&products).Error
	return products, err
}

func (r *ProductRepository) CheckStock(productID uint, quantity float64) (bool, error) {
	var product models.Product
	if err := r.db.Select("stock").First(&product, productID).Error; err !=
		nil {
		return false, err
	}
	return product.Stock >= quantity, nil
}

func (r *ProductRepository) UpdateStock(productID uint, quantity float64) error {
	return r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

func (r *ProductRepository) GetLowStockProducts(minStock float64) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("stock <= min_stock AND is_active = ?", true).Find(&products).Error
	return products, err
}
