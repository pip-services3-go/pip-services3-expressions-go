package parsers

// General syntax errors
const (
	// The unknown
	ErrCodeUnknown = "UNKNOWN"

	// The internal error
	ErrCodeInternal = "INTERNAL"

	// The unexpected end.
	ErrCodeUnexpectedEnd = "UNEXPECTED_END"

	// The error near
	ErrCodeErrorNear = "ERROR_NEAR"

	// The error at
	ErrCodeErrorAt = "ERROR_AT"

	// The unexpected symbol
	ErrCodeUnexpectedSymbol = "UNEXPECTED_SYMBOL"

	// The mismatched brackets
	ErrCodeMismatchedBrackets = "MISTMATCHED_BRACKETS"

	// The missing variable
	ErrCodeMissingVariable = "MISSING_VARIABLE"

	// Not closed section
	ErrCodeNotClosedSection = "NOT_CLOSED_SECTION"

	// Unexpected section end
	ErrCodeUnexpectedSectionEnd = "UNEXPECTED_SECTION_END"
)
