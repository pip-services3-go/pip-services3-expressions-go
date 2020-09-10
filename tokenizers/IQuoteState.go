package tokenizers

// Defines an interface for tokenizer state that processes quoted strings.
type IQuoteState interface {
	ITokenizerState

	// Encodes a string value.
	//
	// Parameters:
	//   - value: A string value to be encoded.
	//   - quoteSymbol: A string quote character.
	// Returns: An encoded string.
	EncodeString(value string, quoteSymbol rune) string

	// Decodes a string value.
	//
	// Parameters:
	//   - value: A string value to be decoded.
	//   - quoteSymbol: A string quote character.
	// Returns: An decoded string.
	DecodeString(value string, quoteSymbol rune) string
}
