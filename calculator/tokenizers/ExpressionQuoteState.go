package tokenizers

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// Implements an Expression-specific quote string state object.
type ExpressionQuoteState struct {
}

func NewExpressionQuoteState() *ExpressionQuoteState {
	c := &ExpressionQuoteState{}
	return c
}

// Gets the next token from the stream started from the character linked to this state.
//
// Parameters:
//   - scanner: A textual string to be tokenized.
//   - tokenizer: A tokenizer class that controls the process.
// Returns: The next token from the top of the stream.
func (c *ExpressionQuoteState) NextToken(scanner io.IScanner,
	tokenizer tokenizers.ITokenizer) *tokenizers.Token {
	firstSymbol := scanner.Read()
	line := scanner.Line()
	column := scanner.Column()
	tokenValue := strings.Builder{}
	tokenValue.WriteRune(firstSymbol)
	nextSymbol := scanner.Read()

	for !utilities.CharValidator.IsEof(nextSymbol) {
		tokenValue.WriteRune(nextSymbol)

		if nextSymbol == firstSymbol {
			chr := scanner.Peek()

			if chr == firstSymbol {
				nextSymbol = scanner.Read()
				tokenValue.WriteRune(nextSymbol)
			} else {
				break
			}
		}

		nextSymbol = scanner.Read()
	}

	tokenType := tokenizers.Quoted
	if firstSymbol == '"' {
		tokenType = tokenizers.Word
	}

	return tokenizers.NewToken(tokenType, tokenValue.String(), line, column)
}

// Encodes a string value.
//
// Parameters:
//   - value: A string value to be encoded.
//   - quoteSymbol: A string quote character.
// Returns: An encoded string.
func (c *ExpressionQuoteState) EncodeString(value string, quoteSymbol rune) string {
	result := strings.Builder{}
	quoteString := string(quoteSymbol)
	result.WriteRune(quoteSymbol)
	result.WriteString(strings.ReplaceAll(value, quoteString, quoteString+quoteString))
	result.WriteRune(quoteSymbol)
	return result.String()
}

// Decodes a string value.
//
// Parameters:
//   - value: A string value to be decoded.
//   - quoteChar: A string quote character.
// Returns: An decoded string.
func (c *ExpressionQuoteState) DecodeString(value string, quoteSymbol rune) string {
	runes := []rune(value)
	if len(runes) >= 2 && runes[0] == quoteSymbol && runes[len(value)-1] == quoteSymbol {
		value = string(runes[1 : len(runes)-1])
		quoteString := string(quoteSymbol)
		value = strings.ReplaceAll(value, quoteString+quoteString, quoteString)
	}
	return value
}
