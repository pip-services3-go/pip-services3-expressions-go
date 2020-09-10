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

	reader := io.NewStringPushbackReader("<A<<<>")

	token, err := state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "A", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "<<", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())

	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "<>", token.Value())
	assert.Equal(t, tokenizers.Symbol, token.Type())
}
