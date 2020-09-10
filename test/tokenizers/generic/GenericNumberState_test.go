package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericNumberStateNextToken(t *testing.T) {
	state := generic.NewGenericNumberState()

	reader := io.NewStringPushbackReader("ABC")
	failed := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				failed = true
			}
		}()
		state.NextToken(reader, nil)
	}()
	assert.True(t, failed)

	reader = io.NewStringPushbackReader("123#")
	token, err := state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "123", token.Value())
	assert.Equal(t, tokenizers.Integer, token.Type())

	reader = io.NewStringPushbackReader("-123#")
	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "-123", token.Value())
	assert.Equal(t, tokenizers.Integer, token.Type())

	reader = io.NewStringPushbackReader("123.#")
	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "123.", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())

	reader = io.NewStringPushbackReader("123.456#")
	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "123.456", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())

	reader = io.NewStringPushbackReader("-123.456#")
	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "-123.456", token.Value())
	assert.Equal(t, tokenizers.Float, token.Type())
}
