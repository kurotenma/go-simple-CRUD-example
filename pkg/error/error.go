package errorTools

import (
	"github.com/labstack/echo/v4"
)

type Interface interface {
	Response(err Enum) error
	ValidationResponse(err Enum, vs Validations) error
	MessageResponse(err Enum, msg string) error
	MessageValidationResponse(err Enum, vs Validations, msg string) error
}

type Error struct {
	Context echo.Context
}

func NewError(ctx echo.Context) Interface {
	return &Error{Context: ctx}
}

func (e *Error) Response(err Enum) error {
	var errStruct ErrorStruct
	errStruct.ErrorEnum = err
	return e.Context.JSON(errStruct.Response())
}
func (e *Error) ValidationResponse(err Enum, vs Validations) error {
	var errStruct ErrorStruct
	errStruct.ErrorEnum = err
	errStruct.Validations = vs
	return e.Context.JSON(errStruct.Response())
}
func (e *Error) MessageResponse(err Enum, msg string) error {
	var errStruct ErrorStruct
	errStruct.ErrorEnum = err
	errStruct.Message = msg
	return e.Context.JSON(errStruct.Response())
}
func (e *Error) MessageValidationResponse(
	err Enum,
	vs Validations,
	msg string,
) error {
	var errStruct ErrorStruct
	errStruct.ErrorEnum = err
	errStruct.Message = msg
	errStruct.Validations = vs
	return e.Context.JSON(errStruct.Response())
}
