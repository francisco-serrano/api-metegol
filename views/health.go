package views

type HealthResponse struct {
	Message string `json:"message"`
}

func NewHealthResponse(message string) *HealthResponse {
	return &HealthResponse{Message: message}
}
