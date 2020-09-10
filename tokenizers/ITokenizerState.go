package tokenizers

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
)

// A tokenizerState returns a token, given a reader, an initial character read from the reader,
// and a tokenizer that is conducting an overall tokenization of the reader. The tokenizer will
// typically have a character state table that decides which state to use, depending on an initial
// character. If a single character is insufficient, a state such as <code>SlashState</code>
// will read a second character, and may delegate to another state, such as <code>SlashStarState</code>.
// This prospect of delegation is the reason that the <code>nextToken()</code>
// method has a tokenizer argument.
type ITokenizerState interface {
	// Gets the next token from the stream started from the character linked to this state.
	//
	// Parameters:
	//   - reader: A textual string to be tokenized.
	//   - tokenizer: A tokenizer class that controls the process.
	// Returns: The next token from the top of the stream.
	NextToken(reader io.IPushbackReader, tokenizer ITokenizer) (*Token, error)
}
