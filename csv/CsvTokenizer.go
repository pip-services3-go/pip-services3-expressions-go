package csv

import "github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"

// Implements a tokenizer class for CSV files.
type CsvTokenizer struct {
	tokenizers.AbstractTokenizer
	fieldSeparators []rune
	quoteSymbols    []rune
	endOfLine       string
}

// Constructs this object with default parameters.
func NewCsvTokenizer() *CsvTokenizer {
	c := &CsvTokenizer{
		fieldSeparators: []rune{','},
		quoteSymbols:    []rune{'"'},
		endOfLine:       string(CR) + string(LF),
	}
	c.AbstractTokenizer = *tokenizers.InheritAbstractTokenizer(c)

	c.SetNumberState(nil)
	c.SetWhitespaceState(nil)
	c.SetCommentState(nil)
	c.SetWordState(NewCsvWordState(c.fieldSeparators, c.quoteSymbols))
	c.SetSymbolState(NewCsvSymbolState())
	c.SetQuoteState(NewCsvQuoteState())
	c.AssignStates()

	return c
}

// Gets separators for fields in CSV stream.
func (c *CsvTokenizer) FieldSeparators() []rune {
	return c.fieldSeparators
}

// Sets separators for fields in CSV stream.
func (c *CsvTokenizer) SetFieldSeparators(value []rune) {
	for _, fieldSeparator := range value {
		if fieldSeparator == CR || fieldSeparator == LF || fieldSeparator == Nil {
			panic("Invalid field separator.")
		}

		for _, quoteSymbol := range c.quoteSymbols {
			if fieldSeparator == quoteSymbol {
				panic("Invalid field separator.")
			}
		}
	}

	c.fieldSeparators = value
	c.SetWordState(NewCsvWordState(value, c.quoteSymbols))
	c.AssignStates()
}

// Gets a separator for rows in CSV stream.
func (c *CsvTokenizer) EndOfLine() string {
	return c.endOfLine
}

// Sets a separator for rows in CSV stream.
func (c *CsvTokenizer) SetEndOfLine(value string) {
	c.endOfLine = value
}

// Gets characters to quote strings
func (c *CsvTokenizer) QuoteSymbols() []rune {
	return c.quoteSymbols
}

// Sets characters to quote strings
func (c *CsvTokenizer) SetQuoteSymbols(value []rune) {
	for _, quoteSymbol := range value {
		if quoteSymbol == CR || quoteSymbol == LF || quoteSymbol == Nil {
			panic("Invalid quote symbol.")
		}

		for _, fieldSeparator := range c.fieldSeparators {
			if quoteSymbol == fieldSeparator {
				panic("Invalid quote symbol.")
			}
		}
	}

	c.quoteSymbols = value
	c.SetWordState(NewCsvWordState(c.fieldSeparators, value))
	c.AssignStates()
}

// Assigns tokenizer states to correct characters.
func (c *CsvTokenizer) AssignStates() {
	c.ClearCharacterStates()
	c.SetCharacterState(0x0000, 0xffff, c.WordState())
	c.SetCharacterState(CR, CR, c.SymbolState())
	c.SetCharacterState(LF, LF, c.SymbolState())

	for _, fieldSeparator := range c.fieldSeparators {
		c.SetCharacterState(fieldSeparator, fieldSeparator, c.SymbolState())
	}

	for _, quoteSymbol := range c.quoteSymbols {
		c.SetCharacterState(quoteSymbol, quoteSymbol, c.QuoteState())
	}
}
