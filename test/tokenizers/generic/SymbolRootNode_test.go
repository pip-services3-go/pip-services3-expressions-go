package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestSymbolRootNodeNextToken(t *testing.T) {
	node := generic.NewSymbolRootNode()
	node.Add("<", tokenizers.Symbol)
	node.Add("<<", tokenizers.Symbol)
	node.Add("<>", tokenizers.Symbol)

	reader := io.NewStringPushbackReader("<A<<<>")

	token, err := node.NextToken(reader)
	assert.Nil(t, err)
	assert.Equal(t, "<", token.Value())

	token, err = node.NextToken(reader)
	assert.Nil(t, err)
	assert.Equal(t, "A", token.Value())

	token, err = node.NextToken(reader)
	assert.Nil(t, err)
	assert.Equal(t, "<<", token.Value())

	token, err = node.NextToken(reader)
	assert.Nil(t, err)
	assert.Equal(t, "<>", token.Value())
}
