package serializers

type ErrorResponse struct {
	Error       string `json:"error"`
	ErrorDetail string `json:"detail"`
}
