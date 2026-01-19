package models

type PaginatedResponse[T any] struct {
	Items     []T   `json:"items"`
	PageCount int64 `json:"page_count"`
	Page      int64 `json:"page"`
}
