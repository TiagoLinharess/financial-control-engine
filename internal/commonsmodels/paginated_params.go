package commonsmodels

import (
	"time"

	"github.com/google/uuid"
)

type PaginatedParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

type PaginatedParamsWithDateRange struct {
	UserID    uuid.UUID
	Limit     int32
	Offset    int32
	StartDate time.Time
	EndDate   time.Time
}
