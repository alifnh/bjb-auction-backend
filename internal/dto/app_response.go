package dto

type ErrorResponse struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
