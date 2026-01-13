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

func Float64ToNumeric(f float64) pgtype.Numeric {
	bigFloat := big.NewFloat(f)
	bigInt := new(big.Int)
	exp := int32(0)

	scale := big.NewFloat(100)
	bigFloat.Mul(bigFloat, scale)
	bigFloat.Int(bigInt)
	exp = -2

	return pgtype.Numeric{
		Int:   bigInt,
		Exp:   exp,
		Valid: true,
	}
}

func NumericToFloat64(n pgtype.Numeric) float64 {
	valuePgtype, err := n.Float64Value()
	if err != nil {
		return 0
	}
	return valuePgtype.Float64
}
