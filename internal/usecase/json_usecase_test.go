package usecase

import (
	"testing"

	"github.com/marcusbiava/json-post-box/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJSONRepository struct {
	mock.Mock
}

func (m *MockJSONRepository) Store(data *domain.JSON) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockJSONRepository) FindByID(id string) (*domain.JSON, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.JSON), args.Error(1)
}

func TestJSONUsecase_Store(t *testing.T) {
	t.Run("should store json successfully", func(t *testing.T) {
		mockRepo := new(MockJSONRepository)
		uc := NewJSONUsecase(mockRepo)

		data := map[string]interface{}{"test": "value"}
		mockRepo.On("Store", mock.AnythingOfType("*domain.JSON")).Return(nil)

		result, err := uc.Store(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result.Data)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when data is nil", func(t *testing.T) {
		mockRepo := new(MockJSONRepository)
		uc := NewJSONUsecase(mockRepo)

		result, err := uc.Store(nil)
		assert.Equal(t, domain.ErrInvalidJSON, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := new(MockJSONRepository)
		uc := NewJSONUsecase(mockRepo)

		data := map[string]interface{}{"test": "value"}
		mockRepo.On("Store", mock.AnythingOfType("*domain.JSON")).Return(domain.ErrStorageFailure)

		result, err := uc.Store(data)
		assert.Equal(t, domain.ErrStorageFailure, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestJSONUsecase_GetByID(t *testing.T) {
	t.Run("should get json by id successfully", func(t *testing.T) {
		mockRepo := new(MockJSONRepository)
		uc := NewJSONUsecase(mockRepo)

		expectedJSON := &domain.JSON{
			ID:   "1",
			Data: map[string]interface{}{"test": "value"},
		}
		mockRepo.On("FindByID", "1").Return(expectedJSON, nil)

		result, err := uc.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, expectedJSON, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when json not found", func(t *testing.T) {
		mockRepo := new(MockJSONRepository)
		uc := NewJSONUsecase(mockRepo)

		mockRepo.On("FindByID", "999").Return(nil, domain.ErrJSONNotFound)

		result, err := uc.GetByID("999")
		assert.Equal(t, domain.ErrJSONNotFound, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
