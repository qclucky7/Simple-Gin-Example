package models

import "net/http"

type Result[T any] struct {
	Status int `json:"status"` //status
	Data   T   `json:"data"`   //data
}

func Ok[T any](data T) Result[T] {
	return Result[T]{
		Status: http.StatusOK,
		Data:   data,
	}
}

func Fail[T any](data T) Result[T] {
	return Result[T]{
		Status: http.StatusBadRequest,
		Data:   data,
	}
}
