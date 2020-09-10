package tokenizers

// Defines an interface for tokenizer state that processes whitespaces (' ', '\t')
type IWhitespaceState interface {
	ITokenizerState

	// Establish the given characters as whitespace to ignore.
	//
	// Parameters:
	//   - fromSymbol: First character index of the interval.
	//   - toSymbol: Last character index of the interval.
	//   - enable: <code>true</code> if this state should ignore characters in the given range.
	SetWhitespaceChars(fromSymbol rune, toSymbol rune, enable bool)

	// Clears definitions of whitespace characters.
	ClearWhitespaceChars()
}
