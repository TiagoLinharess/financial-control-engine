package utils

import (
	"encoding/json"
	"errors"
	"financialcontrol/internal/constants"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, errors.New(constants.NilValueErrorMsg)
	}

	switch x := v.(type) {
	case float64:
		return x, nil
	case float32:
		return float64(x), nil
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint16:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	case json.Number:
		return x.Float64()
	case pgtype.Numeric:
		valuePgtype, err := x.Float64Value()
		if err != nil {
			return 0, err
		}
		return valuePgtype.Float64, nil
	case bool:
		if x {
			return 1, nil
		}
		return 0, nil
	default:
		s := fmt.Sprintf("%v", v)
		if s == "<nil>" {
			return 0, errors.New(constants.NilValueErrorMsg)
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("%s: %T", constants.UnsupportedTypeErrorMsg, v)
		}
		return f, nil
	}
}
