package utils

import (
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func PgTypeUUIDToUUID(pgtypeUUID pgtype.UUID) *uuid.UUID {
	if !pgtypeUUID.Valid {
		return nil
	}
	stringValue := pgtypeUUID.String()
	parsedUUID, err := uuid.Parse(stringValue)

	if err != nil {
		return nil
	}

	return &parsedUUID
}

func UUIDToPgTypeUUID(u *uuid.UUID) pgtype.UUID {
	if u == nil {
		return pgtype.UUID{Valid: false}
	}
	return pgtype.UUID{
		Bytes: *u,
		Valid: true,
	}
}

func Float64ToNumeric(f float64) (pgtype.Numeric, error) {
	return pgtype.Numeric{
		Int:   big.NewInt(int64(f)),
		Exp:   0,
		Valid: true,
	}, nil
}

func NumericToFloat64(n pgtype.Numeric) float64 {
	valuePgtype, err := n.Float64Value()
	if err != nil {
		return 0
	}
	return valuePgtype.Float64
}
