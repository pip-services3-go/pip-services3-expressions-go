package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestCCommentStateNextToken(t *testing.T) {
	state := generic.NewCCommentState()

	reader := io.NewStringPushbackReader("// Comment \n Comment ")
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

	reader = io.NewStringPushbackReader("/* Comment \n Comment */#")
	token, err := state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "/* Comment \n Comment */", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())
}
