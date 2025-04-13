package server

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusbiava/json-post-box/internal/adapter/http/handler"
)

// SetupRoutes configures all application routes
func SetupRoutes(e *echo.Echo, h *handler.JSONHandler) {
	// API v1 routes
	v1 := e.Group("/api/v1")
	{
		// JSON endpoints
		v1.POST("/jsons", h.Store)
		v1.GET("/jsons/:id", h.Get)

		// Health check
		v1.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "healthy"})
		})
	}
}
