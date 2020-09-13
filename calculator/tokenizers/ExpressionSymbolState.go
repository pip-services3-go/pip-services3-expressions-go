package tokenizers

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
)

// Implements a symbol state object.
type ExpressionSymbolState struct {
	generic.GenericSymbolState
}

// Constructs an instance of this class.
func NewExpressionSymbolState() *ExpressionSymbolState {
	c := &ExpressionSymbolState{
		GenericSymbolState: *generic.NewGenericSymbolState(),
	}

	c.Add("<=", tokenizers.Symbol)
	c.Add(">=", tokenizers.Symbol)
	c.Add("<>", tokenizers.Symbol)
	c.Add("!=", tokenizers.Symbol)
	c.Add(">>", tokenizers.Symbol)
	c.Add("<<", tokenizers.Symbol)

	return c
}
