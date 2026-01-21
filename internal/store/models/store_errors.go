package models

import "financialcontrol/internal/constants"

type StoreErrorType string

const (
	ErrNoRows StoreErrorType = constants.StoreErrorNoRowsMsg
)
