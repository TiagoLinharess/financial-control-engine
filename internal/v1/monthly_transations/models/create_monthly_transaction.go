package models

type CreateMonthlyTransaction struct {
	UserID       string
	Name         string
	Value        float64
	Day          int
	CategoryID   string
	CreditCardID string
}
