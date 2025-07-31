package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"ArepasSA/internal/utils"

	"github.com/go-playground/validator/v10"
)

type ClientService struct {
	repo      *repositories.ClientRepository
	validator *validator.Validate
}

func NewClientService(repo *repositories.ClientRepository) *ClientService {
	return &ClientService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	if err := s.validator.Struct(client); err != nil {
		return utils.TranslateValidationErrors(err)
	}
	return s.repo.Create(client)
}

func (s *ClientService) UpdateClient(client *models.Client) error {
	if err := s.validator.Struct(client); err != nil {
		return utils.TranslateValidationErrors(err)
	}
	return s.repo.Update(client)
}

func (s *ClientService) DeleteClient(id uint) error {
	return s.repo.Delete(id)
}

func (s *ClientService) SoftDeleteClient(id uint) error {
	return s.repo.SoftDelete(id)
}

func (s *ClientService) GetClient(id uint) (*models.Client, error) {
	return s.repo.FindByID(id)
}

func (s *ClientService) GetAllClients(activeOnly bool) ([]models.Client, error) {
	return s.repo.FindAll(activeOnly)
}

func (s *ClientService) AddComment(clientID uint, comment *models.ClientComment) error {
	if err := s.validator.Struct(comment); err != nil {
		return utils.TranslateValidationErrors(err)
	}

	// Verificar que el cliente existe
	if _, err := s.repo.FindByID(clientID); err != nil {
		return err
	}

	return s.repo.AddComment(clientID, comment)
}

func (s *ClientService) GetPreferences(clientID uint) ([]models.ClientComment, error) {
	return s.repo.GetComments(clientID, true)
}

func (s *ClientService) GetAllComments(clientID uint) ([]models.ClientComment, error) {
	return s.repo.GetComments(clientID, false)
}
