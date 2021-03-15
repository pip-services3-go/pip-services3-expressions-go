package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestCppCommentStateNextToken(t *testing.T) {
	state := generic.NewCppCommentState()

	scanner := io.NewStringScanner("-- Comment \n Comment ")
	failed := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				failed = true
			}
		}()
		state.NextToken(scanner, nil)
	}()
	assert.True(t, failed)

	scanner = io.NewStringScanner("// Comment \n Comment ")
	token, err := state.NextToken(scanner, nil)
	assert.Nil(t, err)
	assert.Equal(t, "// Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())

	scanner = io.NewStringScanner("/* Comment \n Comment */#")
	token, err = state.NextToken(scanner, nil)
	assert.Nil(t, err)
	assert.Equal(t, "/* Comment \n Comment */", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())
}
