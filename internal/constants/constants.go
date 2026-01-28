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
	StartDateText      = "start_date"
	EndDateText        = "end_date"
)

// STORE ERRORS
const (
	StoreErrorNoRowsMsg = "no rows in result set"
)

// ERRORS_USER_MESSAGES
const (
	UserUnauthorized           = "USER_UNAUTHORIZED"
	InvalidStartDate           = "INVALID_START_DATE"
	InvalidEndDate             = "INVALID_END_DATE"
	TransactionTypeMsg         = "TRANSACTION_TYPE_INVALID"
	TransactionTypeEmptyMsg    = "TRANSACTION_TYPE_EMPTY"
	NameEmptyMsg               = "NAME_EMPTY"
	IconEmptyMsg               = "ICON_EMPTY"
	NameInvalidCharsCountMsg   = "NAME_INVALID_CHARS_COUNT"
	IconInvalidCharsCountMsg   = "ICON_INVALID_CHARS_COUNT"
	LimitReachedMsg            = "LIMIT_REACHED"
	CannotBeDeletedMsg         = "CANNOT_BE_DELETED_BECAUSE_IT_HAS_ASSOCIATED_TRANSACTIONS"
	NotFoundMsg                = "NOT_FOUND"
	ValueInvalidMsg            = "VALUE_INVALID"
	DateEmptyMsg               = "DATE_EMPTY_OR_INVALID"
	DateInvalidMsg             = "DATE_INVALID"
	CreditcardLimitExceededMsg = "CREDITCARD_LIMIT_EXCEEDED"
	InvalidData                = "INVALID_DATA"
	InvalidID                  = "INVALID_ID"
)

// ERRORS
const (
	UserIDNotFound       = "USER_ID_NOT_FOUND"
	UserIDInvalid        = "USER_ID_INVALID"
	InternalServerError  = "INTERNAL_SERVER_ERROR"
	NilValueError        = "NIL_VALUE"
	UnsupportedTypeError = "UNSUPPORTED_TYPE"
	DecodeJsonError      = "DECODE_JSON_ERROR"
	EncodeJsonError      = "ENCODE_JSON_ERROR"
	InvalidFieldError    = "INVALID_FIELD_ERROR"
	LimitError           = "LIMIT_ERROR"
	NotFoundError        = "NOT_FOUND_ERROR"
	StoreError           = "STORE_ERROR"
	UnauthorizedError    = "UNAUTHORIZED_ERROR"
	InvalidPageParam     = "INVALID_PAGE_PARAM"
	CustomError          = "CUSTOM_ERROR"
	BadRequestError      = "BAD_REQUEST_ERROR"
)

// USER
const (
	UserID = "user_id"
)

// CREDIT CARDS
const (
	FirstFourNumbersInvalidMsg          = "FIRST_FOUR_NUMBERS_INVALID"
	LimitInvalidMsg                     = "LIMIT_INVALID"
	ClosingDayInvalidMsg                = "CLOSING_DAY_INVALID"
	ExpireDayInvalidMsg                 = "EXPIRE_DAY_INVALID"
	BackgroundColorEmptyMsg             = "BACKGROUND_COLOR_EMPTY"
	BackgroundColorInvalidCharsCountMsg = "BACKGROUND_COLOR_INVALID_CHARS_COUNT"
	TextColorEmptyMsg                   = "TEXT_COLOR_EMPTY"
	TextColorInvalidCharsCountMsg       = "TEXT_COLOR_INVALID_CHARS_COUNT"
)

// TRANSACTIONS
const (
	CreditWithoutCreditcardMsg     = "TRANSACTION_CREDIT_WITHOUT_CREDITCARD"
	DebitOrIncomeWithCreditcardMsg = "TRANSACTION_DEBIT_OR_INCOME_WITH_CREDITCARD"
)

// MONTHLY TRANSACTIONS
const (
	DayInvalidMsg = "DAY_INVALID"
)
