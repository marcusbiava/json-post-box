package dto

import "github.com/marcusbiava/json-post-box/internal/domain"

// JSONResponse represents the JSON response structure
type JSONResponse struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// FromDomain converts a domain JSON to a JSONResponse
func FromDomain(json *domain.JSON) *JSONResponse {
	if json == nil {
		return nil
	}
	return &JSONResponse{
		ID:   json.ID,
		Data: json.Data,
	}
}

// ToDomain converts a JSONResponse to a domain JSON
func (r *JSONResponse) ToDomain() *domain.JSON {
	return &domain.JSON{
		ID:   r.ID,
		Data: r.Data,
	}
}
