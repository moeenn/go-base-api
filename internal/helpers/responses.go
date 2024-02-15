package helpers

type APIOkResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type APIErrorResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func NewErrorResponse(status int, error string) APIErrorResponse {
	return APIErrorResponse{
		Success:    false,
		Error:      error,
		StatusCode: status,
	}
}

func NewOkResponse[T any](data T) APIOkResponse[T] {
	return APIOkResponse[T]{
		Success: true,
		Data:    data,
	}
}
