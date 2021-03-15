package parsers

import "github.com/pip-services3-go/pip-services3-expressions-go/variants"

// Defines an expression token holder.
type ExpressionToken struct {
	typ   int
	value *variants.Variant
	line int
	column int
}

// Creates an instance of this token and initializes it with specified values.
//
// Parameters:
//   - typ: The type of this token.
//   - value: The value of this token.
//   - line: The line number where the token is.
//   - column: The column number where the token is.
func NewExpressionToken(typ int, value *variants.Variant, line int, column int) *ExpressionToken {
	if value == nil {
		value = variants.EmptyVariant()
	}

	c := &ExpressionToken{
		typ:   typ,
		value: value,
		line: line,
		column: column,
	}
	return c
}

// The type of this token.
func (c *ExpressionToken) Type() int {
	return c.typ
}

// The value of this token.
func (c *ExpressionToken) Value() *variants.Variant {
	return c.value
}

// The line number where the token is.
func (c *ExpressionToken) Line() int {
	return c.line
}

// The column number where the token is.
func (c *ExpressionToken) Column() int {
	return c.column
}