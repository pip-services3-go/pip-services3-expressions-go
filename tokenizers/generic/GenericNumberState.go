package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// A NumberState object returns a number from a reader. This state's idea of a number allows
// an optional, initial minus sign, followed by one or more digits. A decimal point and another string
// of digits may follow these digits.
type GenericNumberState struct{}

func NewGenericNumberState() *GenericNumberState {
	c := &GenericNumberState{}
	return c
}

// Gets the next token from the stream started from the character linked to this state.
//
// Parameters:
//   - reader: A textual string to be tokenized.
//   - tokenizer: A tokenizer class that controls the process.
// Returns: The next token from the top of the stream.
func (c *GenericNumberState) NextToken(
	reader io.IPushbackReader, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	absorbedDot := false
	gotADigit := false
	tokenValue := strings.Builder{}

	nextSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Parses leading minus.
	if nextSymbol == '-' {
		tokenValue.WriteRune('-')
		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	// Parses digits before decimal separator.
	for utilities.CharValidator.IsDigit(nextSymbol) &&
		!utilities.CharValidator.IsEof(nextSymbol) {
		gotADigit = true
		tokenValue.WriteRune(nextSymbol)

		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	// Parses part after the decimal separator.
	if nextSymbol == '.' {
		absorbedDot = true
		tokenValue.WriteRune('.')

		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}

		// Absorb all digits.
		for utilities.CharValidator.IsDigit(nextSymbol) &&
			!utilities.CharValidator.IsEof(nextSymbol) {
			gotADigit = true
			tokenValue.WriteRune(nextSymbol)

			nextSymbol, err = reader.Read()
			if err != nil {
				return nil, err
			}
		}
	}

	// Pushback last unprocessed symbol.
	if !utilities.CharValidator.IsEof(nextSymbol) {
		reader.Pushback(nextSymbol)
	}

	// Process the result.
	if !gotADigit {
		reader.PushbackString(tokenValue.String())
		if tokenizer != nil && tokenizer.SymbolState() != nil {
			return tokenizer.SymbolState().NextToken(reader, tokenizer)
		} else {
			panic("Tokenizer must have an assigned symbol state.")
		}
	}

	tokenType := tokenizers.Integer
	if absorbedDot {
		tokenType = tokenizers.Float
	}
	return tokenizers.NewToken(tokenType, tokenValue.String()), nil
}
