package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"ArepasSA/internal/utils"
	"time"
)

type SaleService struct {
	saleRepo    *repositories.SaleRepository
	productRepo *repositories.ProductRepository
	comboRepo   *repositories.ComboRepository
}

func NewSaleService(
	saleRepo *repositories.SaleRepository,
	productRepo *repositories.ProductRepository,
	comboRepo *repositories.ComboRepository,
) *SaleService {
	return &SaleService{
		saleRepo:    saleRepo,
		productRepo: productRepo,
		comboRepo:   comboRepo,
	}
}

func (s *SaleService) CreateSale(sale *models.Sale) error {
	// Validar stock antes de crear la venta
	for _, item := range sale.Items {
		if item.ProductID != nil {
			hasStock, err := s.productRepo.CheckStock(*item.ProductID, item.Quantity)
			if err != nil || !hasStock {
				return err
			}
		} else if item.ComboID != nil {
			hasStock, err := s.comboRepo.CheckComboStock(*item.ComboID, item.IsPartial)
			if err != nil || !hasStock {
				return err
			}
		}
	}

	// Actualizar inventario
	for _, item := range sale.Items {
		if item.ProductID != nil {
			if err := s.productRepo.UpdateStock(*item.ProductID, -item.Quantity); err != nil {
				return err
			}
		} else if item.ComboID != nil {
			if err := s.comboRepo.UpdateComboStock(*item.ComboID, item.IsPartial); err != nil {
				return err
			}
		}
	}

	return s.saleRepo.Create(sale)
}

func (s *SaleService) GetSale(id uint) (*models.Sale, error) {
	return s.saleRepo.FindByID(id)
}

func (s *SaleService) GetAllSales(startDate, endDate time.Time) ([]models.Sale, error) {
	return s.saleRepo.FindAll(startDate, endDate)
}

func (s *SaleService) GetPeakHoursReport() (map[int]int, error) {
	return s.saleRepo.GetSalesByHour()
}

func (s *SaleService) GetTopProducts(limit int) ([]models.ProductSales, error) {
	products, err := s.saleRepo.GetTopProducts(limit)
	if err != nil {
		return nil, err
	}

	// Obtener nombres de productos
	for i := range products {
		product, err := s.productRepo.FindByID(products[i].ProductID)
		if err == nil {
			products[i].ProductName = product.Name
		}
	}

	return products, nil
}

func (s *SaleService) GetLeastSoldProducts(limit int) ([]models.ProductSales, error) {
	products, err := s.saleRepo.GetLeastSoldProducts(limit)
	if err != nil {
		return nil, err
	}

	// Obtener nombres de productos
	for i := range products {
		product, err := s.productRepo.FindByID(products[i].ProductID)
		if err == nil {
			products[i].ProductName = product.Name
		}
	}

	return products, nil
}

func (s *SaleService) GetSalesByPriceRange() ([]models.PriceRangeSales, error) {
	return s.saleRepo.GetSalesByPriceRange()
}

func (s *SaleService) AddComment(saleID uint, comment *models.SaleComment) error {
	if err := s.validator.Struct(comment); err != nil {
		return utils.TranslateValidationErrors(err)
	}

	// Verificar que la venta existe
	if _, err := s.saleRepo.FindByID(saleID); err != nil {
		return err
	}

	return s.saleRepo.AddComment(saleID, comment)
}

func (s *SaleService) GetComments(saleID uint, includeInternal bool) ([]models.SaleComment, error) {
	return s.saleRepo.GetComments(saleID, includeInternal)
}
