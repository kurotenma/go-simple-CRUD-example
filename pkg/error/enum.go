package errorTools

import (
	"errors"
	"net/http"
)

var (
	InternalServerError = http.StatusInternalServerError
	Unauthorized        = http.StatusUnauthorized
	BadRequest          = http.StatusBadRequest
)

var (
	ErrPasswordMismatch = Enum{
		"ERR_PASSWORD_MISMATCH",
		errors.New("password mismatch"),
		Unauthorized,
	}
	ErrUserNotFound = Enum{
		"ERR_USER_NOT_FOUND",
		errors.New("game not found"),
		Unauthorized,
	}
	ErrUserRegistered = Enum{
		"ERR_USER_REGISTERED",
		errors.New("game already registered"),
		Unauthorized,
	}
	ErrUnauthorized = Enum{
		"ERR_UNAUTHORIZED",
		errors.New("unauthorized"),
		Unauthorized,
	}
	ErrReqBind = Enum{
		"ERR_REQ_BIND",
		errors.New("error request bind"),
		BadRequest,
	}
	ErrReqValidation = Enum{
		"ERR_REQ_VALIDATION",
		errors.New("error request validation"),
		BadRequest,
	}
	ErrInitDB = Enum{
		"ERR_INIT_DB",
		errors.New("error when init db"),
		InternalServerError,
	}
	ErrDTOMapper = Enum{
		"ERR_DTO_MAPPER",
		errors.New("error when mapping DTO"),
		InternalServerError,
	}
	ErrUnknown = Enum{
		"ERR_UNKNOWN",
		errors.New("error unknown"),
		InternalServerError,
	}
	ErrLoadSigningKey = Enum{
		"ERR_LOAD_SIGNING_KEY",
		errors.New("error load signing key"),
		InternalServerError,
	}
	ErrExecQuery = Enum{
		"ERR_EXEC_QUERY",
		errors.New("error executing query"),
		InternalServerError,
	}
	ErrBuildQuery = Enum{
		"ERR_BUILD_QUERY",
		errors.New("error building query"),
		InternalServerError,
	}
	ErrEncryptPassword = Enum{
		"ERR_ENCRYPT_PASSWORD",
		errors.New("error encrypt password"),
		InternalServerError,
	}
	ErrFieldValidation = Enum{
		"ERR_FIELD_VALIDATION",
		errors.New("error field validation"),
		InternalServerError,
	}
	ErrParseTime = Enum{
		"ERR_PARSE_TIME",
		errors.New("error parse time"),
		InternalServerError,
	}
	ErrValidatePassword = Enum{
		"ERR_VALIDATE_PASSWORD",
		errors.New("error validate password"),
		InternalServerError,
	}
	ErrGeneratePassword = Enum{
		"ERR_GENERATE_PASSWORD",
		errors.New("error generate password"),
		InternalServerError,
	}
	ErrCompareHashPassword = Enum{
		"ERR_COMPARE_HASH_PASSWORD",
		errors.New("error compare hash and password"),
		InternalServerError,
	}
)

func GetEnum(e error) Enum {
	switch e {
	case ErrPasswordMismatch.Error:
		return ErrPasswordMismatch
	case ErrUnauthorized.Error:
		return ErrUnauthorized
	case ErrReqBind.Error:
		return ErrReqBind
	case ErrReqValidation.Error:
		return ErrReqValidation
	case ErrInitDB.Error:
		return ErrInitDB
	case ErrDTOMapper.Error:
		return ErrDTOMapper
	case ErrExecQuery.Error:
		return ErrExecQuery
	}
	return ErrUnknown
}
