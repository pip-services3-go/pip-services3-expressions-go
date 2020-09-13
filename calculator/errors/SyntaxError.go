package errors

import cerrors "github.com/pip-services3-go/pip-services3-commons-go/errors"

// Exception that can be thrown by Expression Parser.
func NewSyntaxError(correlationId, code, message string) *cerrors.ApplicationError {
	return &cerrors.ApplicationError{
		Category:      cerrors.BadRequest,
		CorrelationId: correlationId,
		Code:          code,
		Message:       message,
		Status:        400,
	}
}
