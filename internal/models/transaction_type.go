package models

type TransactionType int

const (
	Income TransactionType = iota
	Debit
	Credit
)

func (t TransactionType) IsValid() bool {
	return t == Income || t == Debit || t == Credit
}
