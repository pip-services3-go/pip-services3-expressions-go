package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// A whitespace state ignores whitespace (such as blanks and tabs), and returns the tokenizer's
// next token. By default, all characters from 0 to 32 are whitespace.
type GenericWhitespaceState struct {
	mp *utilities.CharReferenceMap
}

// Constructs a whitespace state with a default idea of what characters are, in fact, whitespace.
func NewGenericWhitespaceState() *GenericWhitespaceState {
	c := &GenericWhitespaceState{
		mp: utilities.NewCharReferenceMap(),
	}
	c.SetWhitespaceChars(0, ' ', true)
	return c
}

// Ignore whitespace (such as blanks and tabs), and return the tokenizer's next token.
//
// Returns: The tokenizer's next token
func (c *GenericWhitespaceState) NextToken(
	reader io.IPushbackReader, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	tokenValue := strings.Builder{}

	nextSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}
	for c.mp.Lookup(nextSymbol) != nil {
		tokenValue.WriteRune(nextSymbol)

		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	if !utilities.CharValidator.IsEof(nextSymbol) {
		reader.Pushback(nextSymbol)
	}

	return tokenizers.NewToken(tokenizers.Whitespace, tokenValue.String()), nil
}

// Establish the given characters as whitespace to ignore.
//
// Parameters:
//   - fromSymbol: First character index of the interval.
//   - toSymbol: Last character index of the interval.
//   - enable: <code>true</code> if this state should ignore characters in the given range.
func (c *GenericWhitespaceState) SetWhitespaceChars(fromSymbol rune, toSymbol rune, enable bool) {
	if enable {
		c.mp.AddInterval(fromSymbol, toSymbol, true)
	} else {
		c.mp.AddInterval(fromSymbol, toSymbol, nil)
	}
}

// Clears definitions of whitespace characters.
func (c *GenericWhitespaceState) ClearWhitespaceChars() {
	c.mp.Clear()
}
