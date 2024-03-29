package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericSymbolStateNextToken(t *testing.T) {
	state := generic.NewGenericSymbolState()
	state.Add("<", tokenizers.Symbol)
	state.Add("<<", tokenizers.Symbol)
	state.Add("<>", tokenizers.Symbol)

	scanner := io.NewStringScanner("<A<<<>")

	token := state.NextToken(scanner, nil)
	assert.Equal(t, "<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "A", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "<<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token = state.NextToken(scanner, nil)
	assert.Equal(t, "<>", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())
}
