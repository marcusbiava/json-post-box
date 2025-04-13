package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcusbiava/json-post-box/internal/adapter/http/dto"
	"github.com/marcusbiava/json-post-box/internal/usecase"
)

// Context defines the interface for the context
type Context interface {
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
	Param(name string) string
}

// EchoContext adapts echo.Context to our Context interface
type EchoContext struct {
	ctx echo.Context
}

func NewEchoContext(c echo.Context) Context {
	return &EchoContext{ctx: c}
}

func (c *EchoContext) Bind(i interface{}) error {
	return c.ctx.Bind(i)
}

func (c *EchoContext) JSON(code int, i interface{}) error {
	return c.ctx.JSON(code, i)
}

func (c *EchoContext) Param(name string) string {
	return c.ctx.Param(name)
}

// JSONHandler handles HTTP requests for JSON operations
type JSONHandler struct {
	uc usecase.JSONUsecase
}

// NewJSONHandler creates a new JSONHandler instance
func NewJSONHandler(uc usecase.JSONUsecase) *JSONHandler {
	return &JSONHandler{uc: uc}
}

// Store handles the creation of new JSON documents
func (h *JSONHandler) Store(c echo.Context) error {
	ctx := NewEchoContext(c)
	var jsonData interface{}

	if err := ctx.Bind(&jsonData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	storedData, err := h.uc.Store(jsonData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to store JSON")
	}

	return ctx.JSON(http.StatusCreated, dto.FromDomain(storedData))
}

// Get handles the retrieval of JSON documents by ID
func (h *JSONHandler) Get(c echo.Context) error {
	ctx := NewEchoContext(c)
	id := ctx.Param("id")

	jsonData, err := h.uc.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve JSON")
	}

	if jsonData == nil {
		return echo.NewHTTPError(http.StatusNotFound, "JSON not found")
	}

	return ctx.JSON(http.StatusOK, jsonData.Data)
}
