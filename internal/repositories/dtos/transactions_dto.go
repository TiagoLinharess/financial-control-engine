package dtos

import (
	"financialcontrol/internal/categories"
	m "financialcontrol/internal/models"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	cr "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"
)

func StoreTransactionListToStoreTransaction(transaction pgs.ListTransactionsByUserAndDateRow) pgs.GetTransactionByIDRow {
	return StoreTransactionPaginatedToStoreTransaction(pgs.ListTransactionsByUserIDPaginatedRow(transaction))
}

func StoreTransactionPaginatedToStoreTransaction(transaction pgs.ListTransactionsByUserIDPaginatedRow) pgs.GetTransactionByIDRow {
	return pgs.GetTransactionByIDRow{
		ID:                                 transaction.ID,
		UserID:                             transaction.UserID,
		Name:                               transaction.Name,
		Date:                               transaction.Date,
		Value:                              transaction.Value,
		Paid:                               transaction.Paid,
		CreatedAt:                          transaction.CreatedAt,
		UpdatedAt:                          transaction.UpdatedAt,
		CategoryID:                         transaction.CategoryID,
		CategoryTransactionType:            transaction.CategoryTransactionType,
		CategoryName:                       transaction.CategoryName,
		CategoryIcon:                       transaction.CategoryIcon,
		CreditcardID:                       transaction.CreditcardID,
		CreditcardName:                     transaction.CreditcardName,
		CreditcardFirstFourNumbers:         transaction.CreditcardFirstFourNumbers,
		CreditcardCreditLimit:              transaction.CreditcardCreditLimit,
		CreditcardCloseDay:                 transaction.CreditcardCloseDay,
		CreditcardExpireDay:                transaction.CreditcardExpireDay,
		CreditcardBackgroundColor:          transaction.CreditcardBackgroundColor,
		CreditcardTextColor:                transaction.CreditcardTextColor,
		MonthlyTransactionsID:              transaction.MonthlyTransactionsID,
		MonthlyTransactionsDay:             transaction.MonthlyTransactionsDay,
		AnnualTransactionsID:               transaction.AnnualTransactionsID,
		AnnualTransactionsMonth:            transaction.AnnualTransactionsMonth,
		AnnualTransactionsDay:              transaction.AnnualTransactionsDay,
		InstallmentTransactionsID:          transaction.InstallmentTransactionsID,
		InstallmentTransactionsInitialDate: transaction.InstallmentTransactionsInitialDate,
		InstallmentTransactionsFinalDate:   transaction.InstallmentTransactionsFinalDate,
	}
}

func StoreTransactionToTransaction(transaction pgs.GetTransactionByIDRow) tm.Transaction {
	category := categories.ShortCategory{
		ID:              *u.PgTypeUUIDToUUID(transaction.CategoryID),
		TransactionType: m.TransactionType(transaction.CategoryTransactionType.Int32),
		Name:            transaction.CategoryName.String,
		Icon:            transaction.CategoryIcon.String,
	}

	var creditcard *cr.ShortCreditCard
	if transaction.CreditcardID.Valid {
		creditcardValue := cr.ShortCreditCard{
			ID:               *u.PgTypeUUIDToUUID(transaction.CreditcardID),
			Name:             transaction.CreditcardName.String,
			FirstFourNumbers: transaction.CreditcardFirstFourNumbers.String,
			Limit:            transaction.CreditcardCreditLimit.Float64,
			CloseDay:         transaction.CreditcardCloseDay.Int32,
			ExpireDay:        transaction.CreditcardExpireDay.Int32,
			BackgroundColor:  transaction.CreditcardBackgroundColor.String,
			TextColor:        transaction.CreditcardTextColor.String,
		}

		creditcard = &creditcardValue
	}

	return tm.Transaction{
		ID:                     transaction.ID,
		UserID:                 transaction.UserID,
		Name:                   transaction.Name,
		Date:                   transaction.Date.Time,
		Value:                  u.NumericToFloat64(transaction.Value),
		Paid:                   transaction.Paid,
		Category:               category,
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              transaction.CreatedAt.Time,
		UpdatedAt:              transaction.UpdatedAt.Time,
	}
}
