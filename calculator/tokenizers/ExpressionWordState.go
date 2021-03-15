package tokenizers

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
)

// Implements a word state object.
type ExpressionWordState struct {
	generic.GenericWordState
}

// Supported expression keywords.
var Keywords []string = []string{
	"AND", "OR", "NOT", "XOR", "LIKE", "IS", "IN", "NULL", "TRUE", "FALSE",
}

// Constructs an instance of this class.
func NewExpressionWordState() *ExpressionWordState {
	c := &ExpressionWordState{
		GenericWordState: *generic.NewGenericWordState(),
	}

	c.ClearWordChars()
	c.SetWordChars('a', 'z', true)
	c.SetWordChars('A', 'Z', true)
	c.SetWordChars('0', '9', true)
	c.SetWordChars('_', '_', true)
	c.SetWordChars(0x00c0, 0x00ff, true)
	c.SetWordChars(0x0100, 0xffff, true)

	return c
}

// Gets the next token from the stream started from the character linked to this state.
//
// Parameters:
//   - scanner: A textual string to be tokenized.
//   - tokenizer: A tokenizer class that controls the process.
// Returns: The next token from the top of the stream.
func (c *ExpressionWordState) NextToken(scanner io.IScanner,
	tokenizer tokenizers.ITokenizer) *tokenizers.Token {
	token := c.GenericWordState.NextToken(scanner, tokenizer)

	for _, keyword := range Keywords {
		if keyword == strings.ToUpper(token.Value()) {
			return tokenizers.NewToken(tokenizers.Keyword, token.Value())
		}
	}

	return token
}
