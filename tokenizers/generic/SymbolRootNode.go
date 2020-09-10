package generic

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
)

// This class is a special case of a <code>SymbolNode</code>. A <code>SymbolRootNode</code>
// object has no symbol of its own, but has children that represent all possible symbols.
type SymbolRootNode struct {
	SymbolNode
}

// Creates and initializes a root node.
func NewSymbolRootNode() *SymbolRootNode {
	return &SymbolRootNode{
		SymbolNode: *NewSymbolNode(nil, 0),
	}
}

// Add the given string as a symbol.
//
// Parameters:
//   - value: The character sequence to add.
func (c *SymbolRootNode) Add(value string, tokenType int) {
	if value == "" {
		panic("Value must have at least 1 character")
	}

	v := []rune(value)
	childNode := c.EnsureChildWithChar(v[0])

	if childNode.TokenType() == tokenizers.Unknown {
		childNode.SetValid(true)
		childNode.SetTokenType(tokenizers.Symbol)
	}

	childNode.AddDescendantLine(v[1:], tokenType)
}

// Return a symbol string from a reader.
//
// Parameters:
//   - reader: A reader to read from
//   - firstChar: The first character of this symbol, already read from the reader.
// Returns: A symbol string from a reader
func (c *SymbolRootNode) NextToken(reader io.IPushbackReader) (*tokenizers.Token, error) {
	nextSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}

	childNode := c.FindChildWithChar(nextSymbol)
	if childNode != nil {
		childNode, err = childNode.DeepestRead(reader)
		if err != nil {
			return nil, err
		}

		childNode = childNode.UnreadToValid(reader)
		childNodeValue := string(childNode.Ancestry())
		return tokenizers.NewToken(childNode.TokenType(), childNodeValue), nil
	} else {
		tokenValue := string([]rune{nextSymbol})
		return tokenizers.NewToken(tokenizers.Symbol, tokenValue), nil
	}
}
