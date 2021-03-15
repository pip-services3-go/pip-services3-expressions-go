package tokenizers

import (
	"github.com/pip-services3-go/pip-services3-expressions-go/io"
)

// A tokenizer divides a string into tokens. This class is highly customizable with regard
// to exactly how this division occurs, but it also has defaults that are suitable for many
// languages. This class assumes that the character values read from the string lie in
// the range 0-255. For example, the Unicode value of a capital A is 65,
// so <code> System.out.println((char)65); </code> prints out a capital A.
// <p>
// The behavior of a tokenizer depends on its character state table. This table is an array
// of 256 <code>TokenizerState</code> states. The state table decides which state to enter
// upon reading a character from the input string.
// <p>
// For example, by default, upon reading an 'A', a tokenizer will enter a "word" state.
// This means the tokenizer will ask a <code>WordState</code> object to consume the 'A',
// along with the characters after the 'A' that form a word. The state's responsibility
// is to consume characters and return a complete token.
// <p>
// The default table sets a SymbolState for every character from 0 to 255,
// and then overrides this with:<blockquote><pre>
// From    To     State
// 0     ' '    whitespaceState
// 'a'    'z'    wordState
// 'A'    'Z'    wordState
// 160     255    wordState
// '0'    '9'    numberState
// '-'    '-'    numberState
// '.'    '.'    numberState
// '"'    '"'    quoteState
// '\''   '\''    quoteState
// '/'    '/'    slashState
// </pre></blockquote>
// In addition to allowing modification of the state table, this class makes each of the states
// above available. Some of these states are customizable. For example, wordState allows customization
// of what characters can be part of a word, after the first character.
type ITokenizer interface {
	// Gets skip unknown characters flag.
	SkipUnknown() bool

	// Sets skip unknown characters flag.
	SetSkipUnknown(value bool)

	// Gets skip whitespaces flag.
	SkipWhitespaces() bool

	// Sets skip whitespaces flag.
	SetSkipWhitespaces(value bool)

	// Gets skip comments flag.
	SkipComments() bool

	// Sets skip comments flag.
	SetSkipComments(value bool)

	// Gets skip End-Of-File token at the end of stream flag.
	SkipEof() bool

	// Sets skip End-Of-File token at the end of stream flag.
	SetSkipEof(value bool)

	// Gets merges whitespaces flag.
	MergeWhitespaces() bool

	// Sets merges whitespaces flag.
	SetMergeWhitespaces(value bool)

	// Gets unifies numbers: "Integers" and "Floats" makes just "Numbers" flag
	UnifyNumbers() bool

	// Sets unifies numbers: "Integers" and "Floats" makes just "Numbers" flag
	SetUnifyNumbers(value bool)

	// Gets decodes quoted strings flag.
	DecodeStrings() bool

	// Sets decodes quoted strings flag.
	SetDecodeStrings(value bool)

	// Gets a token state to process comments.
	CommentState() ICommentState

	// Gets a token state to process numbers.
	NumberState() INumberState

	// Gets a token state to process quoted strings.
	QuoteState() IQuoteState

	// Gets a token state to process symbols (single like "=" or muti-character like "<>")
	SymbolState() ISymbolState

	// Gets a token state to process white space delimiters.
	WhitespaceState() IWhitespaceState

	// Gets a token state to process words or indentificators.
	WordState() IWordState

	// Gets the stream scanner to tokenize.
	Reader() io.IScanner

	// Sets the stream scanner to tokenize.
	SetReader(scanner io.IScanner)

	// Checks if there is the next token exist.
	//
	// Returns: <code>true</code> if scanner has the next token.
	HasNextToken() (bool, error)

	// Gets the next token from the scanner.
	//
	// Returns: Next token of <code>null</code> if there are no more tokens left.
	NextToken() (*Token, error)

	// Tokenizes a textual stream into a list of token structures.
	//
	// Parameters:
	//   - scanner: A textual stream to be tokenized.
	// Returns: A list of token structures.
	TokenizeStream(scanner io.IScanner) ([]*Token, error)

	// Tokenizes a string buffer into a list of tokens structures.
	//
	// Parameters:
	//   - buffer: A string buffer to be tokenized.
	// Returns: A list of token structures.
	TokenizeBuffer(buffer string) ([]*Token, error)

	// Tokenizes a textual stream into a list of strings.
	//
	// Parameters:
	//   - scanner: A textual stream to be tokenized.
	// Returns: A list of token strings.
	TokenizeStreamToStrings(scanner io.IScanner) ([]string, error)

	// Tokenizes a string buffer into a list of strings.
	//
	// Parameters:
	//   - buffer: A string buffer to be tokenized.
	// Returns: A list of token strings.
	TokenizeBufferToStrings(buffer string) ([]string, error)
}
