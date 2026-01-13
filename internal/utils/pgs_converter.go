package utils

import (
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
	var pgtypeNumeric pgtype.Numeric
	err := pgtypeNumeric.Scan(f)
	if err != nil {
		return pgtype.Numeric{}, err
	}
	return pgtypeNumeric, nil
}

func NumericToFloat64(n pgtype.Numeric) (float64, error) {
	valuePgtype, err := n.Float64Value()

	if err != nil {
		return 0, err
	}

	value := valuePgtype.Float64

	return value, nil
}
