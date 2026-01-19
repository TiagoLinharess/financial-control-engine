package constants

// API
const (
	EnvDBUser      = "FINANCIAL_CONTROL_DATABASE_USER"
	EnvDBPassword  = "FINANCIAL_CONTROL_DATABASE_PASSWORD"
	EnvDBHost      = "FINANCIAL_CONTROL_DATABASE_HOST"
	EnvDBPort      = "FINANCIAL_CONTROL_DATABASE_PORT"
	EnvDBName      = "FINANCIAL_CONTROL_DATABASE_NAME"
	EnvAppPort     = "FINANCIAL_CONTROL_APP_PORT"
	DefaultAppPort = "3080"
)

// COMMONS
const (
	EmptyString        = ""
	Success            = "success"
	ID                 = "id"
	LimitText          = "limit"
	LimitDefaultString = "10"
	LimitDefault       = 10
	PageText           = "page"
	PageDefaultString  = "1"
	PageDefault        = 1
	InvalidPageParam   = "invalid.page.param"
	StartDateText      = "start_date"
	EndDateText        = "end_date"
	InvalidStartDate   = "invalid.start.date"
	InvalidEndDate     = "invalid.end.date"
)

// ERRORS
const (
	StoreErrorNoRowsMsg        = "no rows in result set"
	InternalServerErrorMsg     = "internal.server.error"
	NilValueErrorMsg           = "nil.value"
	UnsupportedTypeErrorMsg    = "unsupported.type"
	DecodeJsonErrorMsg         = "decode.json.error"
	EncodeJsonErrorMsg         = "encode.json.error"
	DecodeJsonErrorSystemMsg   = "DecodeJsonError"
	EncodeJsonErrorSystemMsg   = "EncodeJsonError"
	InvalidFieldErrorSystemMsg = "InvalidFieldError"
	LimitErrorSystemMsg        = "LimitError"
	NotFoundErrorSystemMsg     = "NotFoundError"
	StoreErrorSystemMsg        = "StoreError"
	UnauthorizedErrorSystemMsg = "UnauthorizedError"
	CustomError                = "CustomError"
)

// USER
const (
	UserIDNotFoundMsg = "user_id.not.found"
	UserIDInvalidMsg  = "user_id.invalid"
	UserID            = "user_id"
)

// CATEGORIES
const (
	CategoryNotFoundMsg             = "category.not.found"
	CategoryTransactionTypeMsg      = "category.transaction_type.invalid"
	CategoryTransactionTypeEmptyMsg = "category.transaction_type.empty"
	CategoryNameEmptyMsg            = "category.name.empty"
	CategoryIconEmptyMsg            = "category.icon.empty"
	CategoryCannotBeDeletedMsg      = "category.cannot.be.deleted"
	CategoryLimitReachedMsg         = "category.limit.reached"
)

// CREDIT CARDS
const (
	CreditcardNotFoundMsg                = "creditcard.not.found"
	CreditcardNameEmptyMsg               = "creditcard.name.empty"
	CreditcardFirstFourNumbersInvalidMsg = "creditcard.first.four.numbers.invalid"
	CreditcardLimitInvalidMsg            = "creditcard.limit.invalid"
	CreditcardClosingDayInvalidMsg       = "creditcard.closing.day.invalid"
	CreditcardExpireDayInvalidMsg        = "creditcard.expire.day.invalid"
	CreditcardCannotBeDeletedMsg         = "creditcard.cannot.be.deleted"
	CreditcardBackgroundColorEmptyMsg    = "creditcard.background.color.empty"
	CreditcardTextColorEmptyMsg          = "creditcard.text.color.empty"
	CreditcardLimitReachedMsg            = "creditcard.limit.reached"
)

// TRANSACTIONS
const (
	TransactionNameEmptyMsg                   = "transaction.name.empty"
	TransactionNameInvalidCharsCountMsg       = "transaction.name.invalid.chars.count"
	TransactionNotFoundMsg                    = "transaction.not.found"
	TransactionAmountInvalidMsg               = "transaction.amount.invalid"
	TransactionDateInvalidMsg                 = "transaction.date.invalid"
	TransactionCreditWithoutCreditcardMsg     = "transaction.credit.without.creditcard"
	TransactionDebitOrIncomeWithCreditcardMsg = "transaction.debit.or.income.with.creditcard"
	TransactionCreditcardLimitExceededMsg     = "transaction.creditcard.limit.exceeded"
)
