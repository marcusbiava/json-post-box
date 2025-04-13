package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/marcusbiava/json-post-box/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJSONUsecase struct {
	mock.Mock
}

func (m *MockJSONUsecase) Store(data interface{}) (*domain.JSON, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.JSON), args.Error(1)
}

func (m *MockJSONUsecase) GetByID(id string) (*domain.JSON, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.JSON), args.Error(1)
}

func TestJSONHandler_Store(t *testing.T) {
	e := echo.New()
	mockUsecase := new(MockJSONUsecase)
	handler := NewJSONHandler(mockUsecase)

	t.Run("should store json successfully", func(t *testing.T) {
		jsonData := map[string]interface{}{"test": "value"}
		jsonBytes, _ := json.Marshal(jsonData)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonBytes)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		storedJSON := &domain.JSON{
			ID:   "1",
			Data: jsonData,
		}
		mockUsecase.On("Store", jsonData).Return(storedJSON, nil)

		err := handler.Store(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("should return error when invalid json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Store(c)
		assert.Error(t, err)
	})
}

func TestJSONHandler_Get(t *testing.T) {
	e := echo.New()
	mockUsecase := new(MockJSONUsecase)
	handler := NewJSONHandler(mockUsecase)

	t.Run("should get json successfully", func(t *testing.T) {
		jsonData := map[string]interface{}{"test": "value"}
		storedJSON := &domain.JSON{
			ID:   "1",
			Data: jsonData,
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/json/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockUsecase.On("GetByID", "1").Return(storedJSON, nil)

		err := handler.Get(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, jsonData, response)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("should return error when json not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/json/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		mockUsecase.On("GetByID", "999").Return(nil, domain.ErrJSONNotFound)

		err := handler.Get(c)
		assert.Error(t, err)
		mockUsecase.AssertExpectations(t)
	})
}
