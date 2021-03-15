package tokenizers

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// Implements an abstract tokenizer class.
type AbstractTokenizer struct {
	mp *utilities.CharReferenceMap

	skipUnknown      bool
	skipWhitespaces  bool
	skipComments     bool
	skipEof          bool
	mergeWhitespaces bool
	unifyNumbers     bool
	decodeStrings    bool

	commentState    ICommentState
	numberState     INumberState
	quoteState      IQuoteState
	symbolState     ISymbolState
	whitespaceState IWhitespaceState
	wordState       IWordState

	scanner        io.IScanner
	nextToken     *Token
	lastTokenType int
}

func NewAbstractTokenizer() *AbstractTokenizer {
	c := &AbstractTokenizer{
		mp:            utilities.NewCharReferenceMap(),
		lastTokenType: Unknown,
	}
	return c
}

func (c *AbstractTokenizer) SkipUnknown() bool {
	return c.skipUnknown
}

func (c *AbstractTokenizer) SetSkipUnknown(value bool) {
	c.skipUnknown = value
}

func (c *AbstractTokenizer) SkipWhitespaces() bool {
	return c.skipWhitespaces
}

func (c *AbstractTokenizer) SetSkipWhitespaces(value bool) {
	c.skipWhitespaces = value
}

func (c *AbstractTokenizer) SkipComments() bool {
	return c.skipComments
}

func (c *AbstractTokenizer) SetSkipComments(value bool) {
	c.skipComments = value
}

func (c *AbstractTokenizer) SkipEof() bool {
	return c.skipEof
}

func (c *AbstractTokenizer) SetSkipEof(value bool) {
	c.skipEof = value
}

func (c *AbstractTokenizer) MergeWhitespaces() bool {
	return c.mergeWhitespaces
}

func (c *AbstractTokenizer) SetMergeWhitespaces(value bool) {
	c.mergeWhitespaces = value
}

func (c *AbstractTokenizer) UnifyNumbers() bool {
	return c.unifyNumbers
}

func (c *AbstractTokenizer) SetUnifyNumbers(value bool) {
	c.unifyNumbers = value
}

func (c *AbstractTokenizer) DecodeStrings() bool {
	return c.decodeStrings
}

func (c *AbstractTokenizer) SetDecodeStrings(value bool) {
	c.decodeStrings = value
}

func (c *AbstractTokenizer) CommentState() ICommentState {
	return c.commentState
}

func (c *AbstractTokenizer) SetCommentState(value ICommentState) {
	c.commentState = value
}

func (c *AbstractTokenizer) NumberState() INumberState {
	return c.numberState
}

func (c *AbstractTokenizer) SetNumberState(value INumberState) {
	c.numberState = value
}

func (c *AbstractTokenizer) QuoteState() IQuoteState {
	return c.quoteState
}

func (c *AbstractTokenizer) SetQuoteState(value IQuoteState) {
	c.quoteState = value
}

func (c *AbstractTokenizer) SymbolState() ISymbolState {
	return c.symbolState
}

func (c *AbstractTokenizer) SetSymbolState(value ISymbolState) {
	c.symbolState = value
}

func (c *AbstractTokenizer) WhitespaceState() IWhitespaceState {
	return c.whitespaceState
}

func (c *AbstractTokenizer) SetWhitespaceState(value IWhitespaceState) {
	c.whitespaceState = value
}

func (c *AbstractTokenizer) WordState() IWordState {
	return c.wordState
}

func (c *AbstractTokenizer) SetWordState(value IWordState) {
	c.wordState = value
}

func (c *AbstractTokenizer) GetCharacterState(symbol rune) ITokenizerState {
	state, _ := c.mp.Lookup(symbol).(ITokenizerState)
	return state
}

func (c *AbstractTokenizer) SetCharacterState(fromSymbol rune, toSymbol rune, state ITokenizerState) {
	c.mp.AddInterval(fromSymbol, toSymbol, state)
}

func (c *AbstractTokenizer) ClearCharacterStates() {
	c.mp.Clear()
}

func (c *AbstractTokenizer) Reader() io.IScanner {
	return c.scanner
}

func (c *AbstractTokenizer) SetReader(value io.IScanner) {
	c.scanner = value
	c.nextToken = nil
	c.lastTokenType = Unknown
}

func (c *AbstractTokenizer) HasNextToken() (bool, error) {
	if c.nextToken == nil {
		var err error
		c.nextToken, err = c.ReadNextToken()
		if err != nil {
			return false, err
		}
	}
	return c.nextToken != nil, nil
}

func (c *AbstractTokenizer) NextToken() (*Token, error) {
	token := c.nextToken
	if token == nil {
		var err error
		token, err = c.ReadNextToken()
		if err != nil {
			return nil, err
		}
	}
	c.nextToken = nil
	return token, nil
}

func (c *AbstractTokenizer) ReadNextToken() (*Token, error) {
	if c.scanner == nil {
		return nil, nil
	}

	var token *Token = nil
	var err error

	for true {
		// Read character
		nextChar := c.scanner.Peek()

		// If reached Eof then exit
		if utilities.CharValidator.IsEof(nextChar) {
			token = nil
			break
		}

		// Get state for character
		state := c.GetCharacterState(nextChar)
		if state != nil {
			token, err = state.NextToken(c.scanner, c)
			if err != nil {
				return nil, err
			}
		}

		// Check for unknown characters and endless loops...
		if token == nil || token.Value() == "" {
			chr := c.scanner.Read()
			token = NewToken(Unknown, string(chr))
		}

		// Skip unknown characters if option set.
		if token.Type() == Unknown && c.skipUnknown {
			c.lastTokenType = token.Type()
			continue
		}

		// Decode strings is option set.
		if _, ok := state.(IQuoteState); ok && c.decodeStrings {
			token = NewToken(token.Type(), c.QuoteState().DecodeString(token.Value(), nextChar))
		}

		// Skips comments if option set.
		if token.Type() == Comment && c.skipComments {
			c.lastTokenType = token.Type()
			continue
		}

		// Skips whitespaces if option set.
		if token.Type() == Whitespace && c.lastTokenType == Whitespace && c.skipWhitespaces {
			c.lastTokenType = token.Type()
			continue
		}

		// Unifies whitespaces if option set.
		if token.Type() == Whitespace && c.mergeWhitespaces {
			token = NewToken(Whitespace, " ")
		}

		// Unifies numbers if option set.
		if c.unifyNumbers &&
			(token.Type() == Integer || token.Type() == Float || token.Type() == HexDecimal) {
			token = NewToken(Number, token.Value())
		}

		break
	}

	// Adds an Eof if option is not set.
	if token == nil && c.lastTokenType != Eof && !c.skipEof {
		token = NewToken(Eof, "")
	}

	// Assigns the last token type
	c.lastTokenType = Eof
	if token != nil {
		c.lastTokenType = token.Type()
	}

	return token, nil
}

func (c *AbstractTokenizer) TokenizeStream(scanner io.IScanner) ([]*Token, error) {
	c.SetReader(scanner)
	tokenList := []*Token{}

	token, err := c.NextToken()
	if err != nil {
		return tokenList, err
	}
	for token != nil {
		tokenList = append(tokenList, token)

		token, err = c.NextToken()
		if err != nil {
			return tokenList, err
		}
	}

	return tokenList, nil
}

func (c *AbstractTokenizer) TokenizeBuffer(buffer string) ([]*Token, error) {
	scanner := io.NewStringScanner(buffer)
	return c.TokenizeStream(scanner)
}

func (c *AbstractTokenizer) TokenizeStreamToStrings(scanner io.IScanner) ([]string, error) {
	c.SetReader(scanner)
	stringList := []string{}

	token, err := c.NextToken()
	if err != nil {
		return stringList, err
	}

	for token != nil {
		stringList = append(stringList, token.Value())

		token, err = c.NextToken()
		if err != nil {
			return stringList, err
		}
	}

	return stringList, nil
}

func (c *AbstractTokenizer) TokenizeBufferToStrings(buffer string) ([]string, error) {
	scanner := io.NewStringScanner(buffer)
	return c.TokenizeStreamToStrings(scanner)
}
