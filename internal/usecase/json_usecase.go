package usecase

import (
	"github.com/marcusbiava/json-post-box/internal/domain"
)

// JSONUsecase defines the interface for JSON operations
type JSONUsecase interface {
	Store(data interface{}) (*domain.JSON, error)
	GetByID(id string) (*domain.JSON, error)
}

type jsonUsecase struct {
	repo domain.JSONRepository
}

// NewJSONUsecase creates a new JSONUsecase instance
func NewJSONUsecase(repo domain.JSONRepository) JSONUsecase {
	return &jsonUsecase{repo: repo}
}

func (uc *jsonUsecase) Store(data interface{}) (*domain.JSON, error) {
	if data == nil {
		return nil, domain.ErrInvalidJSON
	}

	jsonData := &domain.JSON{
		Data: data,
	}

	if err := jsonData.Validate(); err != nil {
		return nil, err
	}

	if err := uc.repo.Store(jsonData); err != nil {
		return nil, domain.ErrStorageFailure
	}

	return jsonData, nil
}

func (uc *jsonUsecase) GetByID(id string) (*domain.JSON, error) {
	jsonData, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if jsonData == nil {
		return nil, domain.ErrJSONNotFound
	}
	return jsonData, nil
}
