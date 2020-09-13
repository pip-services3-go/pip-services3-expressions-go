package parsers

import "github.com/pip-services3-go/pip-services3-expressions-go/variants"

// Defines an expression token holder.
type ExpressionToken struct {
	typ   int
	value *variants.Variant
}

// Creates an instance of this token and initializes it with specified values.
//
// Parameters:
//   - type: The type of this token.
//   - value: The value of this token.
func NewExpressionToken(typ int, value *variants.Variant) *ExpressionToken {
	if value == nil {
		value = variants.EmptyVariant()
	}

	c := &ExpressionToken{
		typ:   typ,
		value: value,
	}
	return c
}

// Creates an instance of this class with specified type and Null value.
//
// Parameters:
//   - type: The type of this token.
func EmptyExpressionToken(typ int) *ExpressionToken {
	return NewExpressionToken(typ, nil)
}

// The type of this token.
func (c *ExpressionToken) Type() int {
	return c.typ
}

// The value of this token.
func (c *ExpressionToken) Value() *variants.Variant {
	return c.value
}
