package domian

type Result[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResult[T any](code int, message string, data T) Result[T] {
	return Result[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func Error[T any](code int, message string) Result[T] {
	return Result[T]{
		Code:    code,
		Message: message,
	}
}
func SuccessData[T any](data T) Result[T] {
	return Result[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}
func Success[T any]() Result[T] {
	return Result[T]{
		Code:    200,
		Message: "success",
	}
}
