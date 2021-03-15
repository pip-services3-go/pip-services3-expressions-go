package csv

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
)

// Implements a symbol state to tokenize delimiters in CSV streams.
type CsvSymbolState struct {
	generic.GenericSymbolState
}

// Constructs this object with specified parameters.
func NewCsvSymbolState() *CsvSymbolState {
	c := &CsvSymbolState{
		GenericSymbolState: *generic.NewGenericSymbolState(),
	}

	c.Add("\n", tokenizers.Eol)
	c.Add("\r", tokenizers.Eol)
	c.Add("\r\n", tokenizers.Eol)
	c.Add("\n\r", tokenizers.Eol)

	return c
}

func (c *CsvSymbolState) NextToken(
	scanner io.IScanner, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	// Optimization...
	nextSymbol := scanner.Read()
	if nextSymbol != '\n' && nextSymbol != '\r' {
		return tokenizers.NewToken(tokenizers.Symbol, string(nextSymbol)), nil
	} else {
		scanner.Unread()
		return c.GenericSymbolState.NextToken(scanner, tokenizer)
	}
}
