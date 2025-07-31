package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"ArepasSA/internal/utils"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	repo      *repositories.ProductRepository
	validator *validator.Validate
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo:      repo,
		validator: validator.New(),
	}
}
func (s *ProductService) CreateProduct(product *models.Product) error {
	if err := s.validator.Struct(product); err != nil {
		return utils.TranslateValidationErrors(err)
	}
	return s.repo.Create(product)
}
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if err := s.validator.Struct(product); err != nil {
		return utils.TranslateValidationErrors(err)
	}
	return s.repo.Update(product)
}
func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
func (s *ProductService) SoftDeleteProduct(id uint) error {
	return s.repo.SoftDelete(id)
}
func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProducts(activeOnly bool) ([]models.Product,
	error) {
	return s.repo.FindAll(activeOnly)
}
func (s *ProductService) CheckStock(productID uint, quantity float64) (bool,
	error) {
	return s.repo.CheckStock(productID, quantity)
}
func (s *ProductService) UpdateStock(productID uint, quantity float64) error {
	return s.repo.UpdateStock(productID, quantity)
}
func (s *ProductService) GetLowStockAlerts() ([]models.Product, error) {
	return s.repo.GetLowStockProducts(0) // 0 para obtener todos los productos por debajo de su stock mínimo
}
