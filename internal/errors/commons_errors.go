package errors

import "financialcontrol/internal/constants"

func UserNotFound(detail string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   constants.UserUnauthorized,
		SystemMessage: constants.UserIDNotFound,
		SystemDetail:  detail,
	}
}

func UserIDInvalid(detail string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   constants.UserUnauthorized,
		SystemMessage: constants.UserIDInvalid,
		SystemDetail:  detail,
	}
}

func InvalidDecodeJsonError(detail string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   constants.InvalidData,
		SystemMessage: constants.DecodeJsonError,
		SystemDetail:  detail,
	}
}

func InvalidFieldError(message string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   message,
		SystemMessage: constants.InvalidFieldError,
		SystemDetail:  message,
	}
}

func BadRequestError(detail string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   detail,
		SystemMessage: constants.BadRequestError,
		SystemDetail:  detail,
	}
}

func NotFoundError() ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   constants.NotFoundMsg,
		SystemMessage: constants.NotFoundError,
		SystemDetail:  constants.StoreErrorNoRowsMsg,
	}
}

func InternalServerError(detail string) ApiErrorItem {
	return ApiErrorItem{
		UserMessage:   constants.InternalServerError,
		SystemMessage: constants.InternalServerError,
		SystemDetail:  detail,
	}
}
