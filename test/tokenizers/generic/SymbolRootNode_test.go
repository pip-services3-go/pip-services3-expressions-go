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

	scanner := io.NewStringScanner("<A<<<>")

	token := node.NextToken(scanner)
	assert.Equal(t, "<", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "A", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "<<", token.Value())

	token = node.NextToken(scanner)
	assert.Equal(t, "<>", token.Value())
}
