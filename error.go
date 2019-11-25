package objectia

import "fmt"

// Error for API requests
type Error struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// String() converts the error into a human-readable string.
func (e *Error) String() string {
	return fmt.Sprintf("%s (status code: %d)", e.Message, e.Status)
}

// Error() performs as String().
func (e *Error) Error() string {
	return e.String()
}

// NewResponseError creates a new API response error.
func NewResponseError(status int, code, message string) error {
	return &Error{
		Status:  status,
		Success: false,
		Code:    code,
		Message: message,
	}
}

// NewError creates a new error.
func NewError(code, message string) error {
	return &Error{
		Status:  0,
		Success: false,
		Code:    code,
		Message: message,
	}
}
