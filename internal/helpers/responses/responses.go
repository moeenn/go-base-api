package responses

type apiOkResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type apiErrorResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func NewErrorResponse(status int, error string) apiErrorResponse {
	return apiErrorResponse{
		Success:    false,
		Error:      error,
		StatusCode: status,
	}
}

func NewOkResponse[T any](data T) apiOkResponse[T] {
	return apiOkResponse[T]{
		Success: true,
		Data:    data,
	}
}
