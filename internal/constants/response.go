package constants

// UserResponse represents a the response body of a Cars request
// Success response
// swagger:response ok
type UserResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse is the response models sent on errors
type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}
