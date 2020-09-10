package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// A quoteState returns a quoted string token from a reader. This state will collect characters
// until it sees a match to the character that the tokenizer used to switch to this state.
// For example, if a tokenizer uses a double-quote character to enter this state,
// then <code>nextToken()</code> will search for another double-quote until it finds one
// or finds the end of the reader.
type GenericQuoteState struct{}

func NewGenericQuoteState() *GenericQuoteState {
	c := &GenericQuoteState{}
	return c
}

// Return a quoted string token from a reader. This method will collect
// characters until it sees a match to the character that the tokenizer used
// to switch to this state.
//
// Returns: A quoted string token from a reader.
func (c *GenericQuoteState) NextToken(
	reader io.IPushbackReader, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	firstSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}

	tokenValue := strings.Builder{}
	tokenValue.WriteRune(firstSymbol)

	nextSymbol, err1 := reader.Read()
	if err1 != nil {
		return nil, err1
	}
	for !utilities.CharValidator.IsEof(nextSymbol) {
		tokenValue.WriteRune(nextSymbol)
		if nextSymbol == firstSymbol {
			break
		}

		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	return tokenizers.NewToken(tokenizers.Quoted, tokenValue.String()), nil
}

// Encodes a string value.
//
// Parameters:
//   - value: A string value to be encoded.
//   - quoteSymbol: A string quote character.
// Returns: An encoded string.
func (c *GenericQuoteState) EncodeString(value string, quoteSymbol rune) string {
	result := strings.Builder{}
	result.WriteRune(quoteSymbol)
	result.WriteString(value)
	result.WriteRune(quoteSymbol)
	return result.String()
}

// Decodes a string value.
//
// Parameters:
//   - value: A string value to be decoded.
//   - quoteSymbol: A string quote character.
// Returns>: An decoded string.
func (c *GenericQuoteState) DecodeString(value string, quoteSymbol rune) string {
	runes := []rune(value)
	if len(runes) >= 2 && runes[0] == quoteSymbol && runes[len(runes)-1] == quoteSymbol {
		return string(runes[1 : len(runes)-1])
	}
	return value
}
