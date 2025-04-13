package mocks

import (
	"github.com/marcusbiava/json-post-box/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockJSONUsecase struct {
	mock.Mock
}

func (m *MockJSONUsecase) Store(data map[string]interface{}) (*domain.JSON, error) {
	args := m.Called(data)
	return args.Get(0).(*domain.JSON), args.Error(1)
}

func (m *MockJSONUsecase) GetByID(id string) (*domain.JSON, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.JSON), args.Error(1)
}
