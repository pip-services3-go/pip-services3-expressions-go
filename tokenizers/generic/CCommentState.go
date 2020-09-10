package generic

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// This state will either delegate to a comment-handling state, or return a token with just a slash in it.
type CCommentState struct {
	CppCommentState
}

func NewCCommentState() *CCommentState {
	c := &CCommentState{
		CppCommentState: *NewCppCommentState(),
	}
	return c
}

// Either delegate to a comment-handling state, or return a token with just a slash in it.
//
// Returns: Either just a slash token, or the results of delegating to a comment-handling state.
func (c *CCommentState) NextToken(
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
