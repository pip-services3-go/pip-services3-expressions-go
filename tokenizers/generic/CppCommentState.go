package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// This state will either delegate to a comment-handling state, or return a token with just a slash in it.
type CppCommentState struct{}

func NewCppCommentState() *CppCommentState {
	c := &CppCommentState{}
	return c
}

// Ignore everything up to a closing star and slash, and then return the tokenizer's next token.
func (c *CppCommentState) GetMultiLineComment(reader io.IPushbackReader) (string, error) {
	result := strings.Builder{}

	lastSymbol := rune(0)
	nextSymbol, err := reader.Read()
	if err != nil {
		return "", err
	}
	for !utilities.CharValidator.IsEof(nextSymbol) {
		result.WriteRune(nextSymbol)
		if lastSymbol == '*' && nextSymbol == '/' {
			break
		}
		lastSymbol = nextSymbol

		nextSymbol, err = reader.Read()
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

// Ignore everything up to an end-of-line and return the tokenizer's next token.
func (c *CppCommentState) GetSingleLineComment(reader io.IPushbackReader) (string, error) {

	result := strings.Builder{}

	nextSymbol, err := reader.Read()
	if err != nil {
		return "", err
	}
	for !utilities.CharValidator.IsEof(nextSymbol) && !utilities.CharValidator.IsEol(nextSymbol) {
		result.WriteRune(nextSymbol)

		nextSymbol, err = reader.Read()
		if err != nil {
			return "", err
		}
	}

	if utilities.CharValidator.IsEol(nextSymbol) {
		reader.Pushback(nextSymbol)
	}

	return result.String(), nil
}

// Either delegate to a comment-handling state, or return a token with just a slash in it.
//
// Returns: Either just a slash token, or the results of delegating to a comment-handling state.
func (c *CppCommentState) NextToken(
	reader io.IPushbackReader, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	firstSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if firstSymbol != '/' {
		reader.Pushback(firstSymbol)
		panic("Incorrect usage of CppCommentState.")
	}

	secondSymbol, err1 := reader.Read()
	if err1 != nil {
		return nil, err1
	}
	if secondSymbol == '*' {
		str, err2 := c.GetMultiLineComment(reader)
		if err2 != nil {
			return nil, err2
		}
		return tokenizers.NewToken(tokenizers.Comment, "/*"+str), nil
	} else if secondSymbol == '/' {
		str, err2 := c.GetSingleLineComment(reader)
		if err2 != nil {
			return nil, err2
		}
		return tokenizers.NewToken(tokenizers.Comment, "//"+str), nil
	} else {
		if !utilities.CharValidator.IsEof(secondSymbol) {
			reader.Pushback(secondSymbol)
		}
		if !utilities.CharValidator.IsEof(firstSymbol) {
			reader.Pushback(firstSymbol)
		}
		return tokenizer.SymbolState().NextToken(reader, tokenizer)
	}
}
