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
	v := validator.New()
	_ = v.RegisterValidation("phone", utils.ValidatePhone)
	
	return &ClientService{
		repo:      repo,
		validator: v,
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

func (s *ClientService) GetAllClients(activeOnly bool, search string) ([]models.Client, error) {
	return s.repo.FindAll(activeOnly, search)
}

func (s *ClientService) AddComment(comment *models.ClientComment) error {
	// Validar que el cliente existe
	if _, err := s.repo.FindByID(comment.ClientID); err != nil {
		return err
	}
	
	if err := s.validator.Struct(comment); err != nil {
		return utils.TranslateValidationErrors(err)
	}
	
	return s.repo.AddComment(comment)
}

func (s *ClientService) GetClientPreferences(clientID uint) ([]models.ClientComment, error) {
	return s.repo.GetComments(clientID, "preference")
}

func (s *ClientService) GetAllClientComments(clientID uint) ([]models.ClientComment, error) {
	return s.repo.GetComments(clientID, "")
}

func (s *ClientService) GetClientsWithPreferences() ([]models.Client, error) {
	return s.repo.GetClientsWithPreferences()
}
