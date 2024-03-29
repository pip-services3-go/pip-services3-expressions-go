package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericWhitespaceStateNextToken(t *testing.T) {
	state := generic.NewGenericWhitespaceState()

	scanner := io.NewStringScanner(" \t\n\r ")
	token := state.NextToken(scanner, nil)
	assert.Equal(t, " \t\n\r ", token.Value())
	assert.Equal(t, tokenizers.Whitespace, token.Type())
}
