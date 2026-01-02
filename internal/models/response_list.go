package models

type ResponseList[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
}
