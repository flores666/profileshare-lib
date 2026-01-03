package api

type AppResponse struct {
	Status      string            `json:"status"` //error, ok
	Data        any               `json:"data,omitempty"`
	Validations []ValidationError `json:"validations,omitempty"`
	Message     string            `json:"message,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

// NewError creates [StatusError] api handlers.
// validations can be nil
func NewError(msg string, validations *ValidationErrors) AppResponse {
	response := AppResponse{
		Status: StatusError,
	}

	if validations != nil {
		response.Validations = validations.Validations
	}

	if msg != "" {
		response.Message = msg
	}

	return response
}

// NewOk creates [StatusOk] api handlers.
// Payload can be nil
func NewOk(msg string, payload any) AppResponse {
	response := AppResponse{
		Status:  StatusOk,
		Message: msg,
	}

	if payload != nil {
		response.Data = payload
	}

	return response
}

// Ok checks if handlers status is [StatusOk]
func (r AppResponse) Ok() bool {
	return r.Status == StatusOk
}
