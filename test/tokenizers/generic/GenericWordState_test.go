package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericWordStateNextToken(t *testing.T) {
	state := generic.NewGenericWordState()

	reader := io.NewStringPushbackReader("AB_CD=")
	token, err := state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "AB_CD", token.Value())
	assert.Equal(t, tokenizers.Word, token.Type())
}
