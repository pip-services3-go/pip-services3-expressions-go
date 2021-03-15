package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericQuoteStateNextToken(t *testing.T) {
	state := generic.NewGenericQuoteState()

	scanner := io.NewStringScanner("'ABC#DEF'#")
	token, err := state.NextToken(scanner, nil)
	assert.Nil(t, err)
	assert.Equal(t, "'ABC#DEF'", token.Value())
	assert.Equal(t, tokenizers.Quoted, token.Type())

	scanner = io.NewStringScanner("'ABC#DEF''")
	token, err = state.NextToken(scanner, nil)
	assert.Nil(t, err)
	assert.Equal(t, "'ABC#DEF'", token.Value())
	assert.Equal(t, tokenizers.Quoted, token.Type())
}

func TestGenericQuoteStateEncodeAndDecodeString(t *testing.T) {
	state := generic.NewGenericQuoteState()

	value := state.EncodeString("ABC", '\'')
	assert.Equal(t, "'ABC'", value)

	value = state.DecodeString(value, '\'')
	assert.Equal(t, "ABC", value)

	value = state.DecodeString("'ABC'DEF'", '\'')
	assert.Equal(t, "ABC'DEF", value)
}
