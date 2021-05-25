package helpers

// Global error response message
func ErrorResponse(code int, err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": err.Error(),
	}
}
