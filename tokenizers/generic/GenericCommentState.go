package generic

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
)

// A CommentState object returns a comment from a scanner.
type GenericCommentState struct{}

func NewGenericCommentState() *GenericCommentState {
	c := &GenericCommentState{}
	return c
}

// Either delegate to a comment-handling state, or return a token with just a slash in it.
//
// Returns: Either just a slash token, or the results of delegating to a comment-handling state
func (c *GenericCommentState) NextToken(
	scanner io.IScanner, tokenizer tokenizers.ITokenizer) *tokenizers.Token {

	tokenValue := strings.Builder{}
	nextSymbol := scanner.Read()
	line := scanner.Line()
	column := scanner.Column()

	for !utilities.CharValidator.IsEof(nextSymbol) && nextSymbol != '\n' && nextSymbol != '\r' {
		tokenValue.WriteRune(nextSymbol)
		nextSymbol = scanner.Read()
	}

	if !utilities.CharValidator.IsEof(nextSymbol) {
		scanner.Unread()
	}

	return tokenizers.NewToken(tokenizers.Comment, tokenValue.String(), line, column)
}
