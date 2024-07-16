package contracts

type JsonResponse[T any] struct {
	Result T      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}
