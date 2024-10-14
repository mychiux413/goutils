package c

import (
	"errors"
	"fmt"
)

// 廣義的錯誤
var (
	ErrInvalidInput              = errors.New("invalid input")
	ErrHTTPError                 = errors.New("http error")
	ErrAmountError               = errors.New("amount error")
	ErrUnknownError              = errors.New("unknown error")
	ErrInternalServerError       = errors.New("internal server error")
	ErrAlreadyExists             = errors.New("already exists")
	ErrNotExists                 = errors.New("doesn't exists")
	ErrUnauthorized              = errors.New("unauthorized")
	ErrNotImplemented            = errors.New("not implemented")
	ErrForbidden                 = errors.New("forbidden")
	ErrInsufficientBalance       = errors.New("insufficient balance")
	ErrSignatureError            = errors.New("signature error")
	ErrServerMaintaining         = fmt.Errorf("%w: server is maintaining", ErrInternalServerError)
	ErrInvalidInputTimeRange     = fmt.Errorf("%w: invalid time range", ErrInvalidInput)
	ErrInvalidInputTimeFormat    = fmt.Errorf("%w: invalid time format", ErrInvalidInput)
	ErrAmountMustBeInteger       = fmt.Errorf("%w: amount must be integer", ErrAmountError)
	ErrAmountMustGreaterThanZero = fmt.Errorf("%w: amount must greater than zero", ErrAmountError)
	ErrIPNotInWhiteList          = errors.New("ip not in white list")
	ErrRequestTooFrequent        = fmt.Errorf("%w: request too frequent", ErrHTTPError)
	ErrRequestTimeout            = fmt.Errorf("%w: request timeout", ErrHTTPError)
	ErrBadRequest                = errors.New("bad request")
	ErrNothingUpdated            = fmt.Errorf("%w: nothingUpdated", ErrBadRequest)
	ErrNothingDeleted            = fmt.Errorf("%w: ErrNothingDeleted", ErrBadRequest)
)

func ErrorsIn(err error, targets []error) bool {
	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}
