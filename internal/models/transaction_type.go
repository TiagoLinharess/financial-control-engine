package models

type TransactionType int

const (
	Income TransactionType = iota
	Debit
	Credit
)
