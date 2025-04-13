package memory

import (
	"testing"

	"github.com/marcusbiava/json-post-box/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMemoryJSONRepository_Store(t *testing.T) {
	repo := NewJSONRepository()

	t.Run("should store json successfully", func(t *testing.T) {
		json := &domain.JSON{
			Data: map[string]interface{}{"test": "value"},
		}

		err := repo.Store(json)
		assert.NoError(t, err)
		assert.Equal(t, "1", json.ID)
	})

	t.Run("should return error when json is nil", func(t *testing.T) {
		err := repo.Store(nil)
		assert.Equal(t, domain.ErrInvalidJSON, err)
	})

	t.Run("should increment id for each stored json", func(t *testing.T) {
		json1 := &domain.JSON{Data: "test1"}
		json2 := &domain.JSON{Data: "test2"}

		_ = repo.Store(json1)
		_ = repo.Store(json2)

		assert.Equal(t, "2", json1.ID)
		assert.Equal(t, "3", json2.ID)
	})
}

func TestMemoryJSONRepository_FindByID(t *testing.T) {
	repo := NewJSONRepository()

	t.Run("should find json by id", func(t *testing.T) {
		expectedJSON := &domain.JSON{
			Data: map[string]interface{}{"test": "value"},
		}
		_ = repo.Store(expectedJSON)

		result, err := repo.FindByID(expectedJSON.ID)
		assert.NoError(t, err)
		assert.Equal(t, expectedJSON, result)
	})

	t.Run("should return error when id is empty", func(t *testing.T) {
		result, err := repo.FindByID("")
		assert.Equal(t, domain.ErrInvalidJSON, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when json not found", func(t *testing.T) {
		result, err := repo.FindByID("999")
		assert.Equal(t, domain.ErrJSONNotFound, err)
		assert.Nil(t, result)
	})
}
