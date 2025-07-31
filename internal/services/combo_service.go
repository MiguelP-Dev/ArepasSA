package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"ArepasSA/internal/utils"

	"github.com/go-playground/validator/v10"
)

type ComboService struct {
	comboRepo   *repositories.ComboRepository
	productRepo *repositories.ProductRepository
	validator   *validator.Validate
}

func NewComboService(comboRepo *repositories.ComboRepository, productRepo *repositories.ProductRepository) *ComboService {
	return &ComboService{
		comboRepo:   comboRepo,
		productRepo: productRepo,
		validator:   validator.New(),
	}
}

func (s *ComboService) CreateCombo(combo *models.Combo) error {
	if err := s.validator.Struct(combo); err != nil {
		return utils.TranslateValidationErrors(err)
	}

	// Validar que todos los productos existan
	for _, item := range combo.Items {
		if _, err := s.productRepo.FindByID(item.ProductID); err != nil {
			return err
		}
	}

	return s.comboRepo.Create(combo)
}

func (s *ComboService) UpdateCombo(combo *models.Combo) error {
	if err := s.validator.Struct(combo); err != nil {
		return utils.TranslateValidationErrors(err)
	}

	// Validar que todos los productos existan
	for _, item := range combo.Items {
		if _, err := s.productRepo.FindByID(item.ProductID); err != nil {
			return err
		}
	}

	return s.comboRepo.Update(combo)
}

func (s *ComboService) DeleteCombo(id uint) error {
	return s.comboRepo.Delete(id)
}

func (s *ComboService) GetCombo(id uint) (*models.Combo, error) {
	return s.comboRepo.FindByID(id)
}

func (s *ComboService) GetAllCombos(activeOnly bool) ([]models.Combo, error) {
	return s.comboRepo.FindAll(activeOnly)
}

func (s *ComboService) SellCombo(comboID uint, quantity float64) error {
	// Verificar stock
	hasStock, err := s.comboRepo.CheckComboStock(comboID, false)
	if err != nil || !hasStock {
		return err
	}

	// Actualizar stock
	for i := 0; i < int(quantity); i++ {
		if err := s.comboRepo.UpdateComboStock(comboID, false); err != nil {
			return err
		}
	}

	return nil
}

func (s *ComboService) SellPartialCombo(comboID uint) error {
	// Verificar stock
	hasStock, err := s.comboRepo.CheckComboStock(comboID, true)
	if err != nil || !hasStock {
		return err
	}

	// Actualizar stock
	return s.comboRepo.UpdateComboStock(comboID, true)
}