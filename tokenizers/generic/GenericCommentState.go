package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// A CommentState object returns a comment from a reader.
type GenericCommentState struct{}

func NewGenericCommentState() *GenericCommentState {
	c := &GenericCommentState{}
	return c
}

// Either delegate to a comment-handling state, or return a token with just a slash in it.
//
// Returns: Either just a slash token, or the results of delegating to a comment-handling state
func (c *GenericCommentState) NextToken(
	reader io.IPushbackReader, tokenizer tokenizers.ITokenizer) (*tokenizers.Token, error) {

	tokenValue := strings.Builder{}

	nextSymbol, err := reader.Read()
	if err != nil {
		return nil, err
	}
	for !utilities.CharValidator.IsEof(nextSymbol) && nextSymbol != '\n' && nextSymbol != '\r' {
		tokenValue.WriteRune(nextSymbol)

		nextSymbol, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	if !utilities.CharValidator.IsEof(nextSymbol) {
		reader.Pushback(nextSymbol)
	}

	return tokenizers.NewToken(tokenizers.Comment, tokenValue.String()), nil
}
