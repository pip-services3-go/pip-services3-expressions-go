package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericCommentStateNextToken(t *testing.T) {
	state := generic.NewGenericCommentState()

	reader := io.NewStringPushbackReader("# Comment \r# Comment ")
	token, err := state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "# Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())

	reader = io.NewStringPushbackReader("# Comment \n# Comment ")
	token, err = state.NextToken(reader, nil)
	assert.Nil(t, err)
	assert.Equal(t, "# Comment ", token.Value())
	assert.Equal(t, tokenizers.Comment, token.Type())
}
