package io

// Defines reader with ability to push back characters.
// This reader is used by tokenizers to process input streams.
type IPushbackReader interface {
	// Reads character from the top of the stream.
	//
	// Returns a read character or <code>-1</code> if stream processed to the end.</returns>
	Read() (rune, error)

	// Returns the character from the top of the stream without moving the stream pointer.
	//
	// Returns a character from the top of the stream or <code>-1</code> if stream is empty.</returns>
	Peek() (rune, error)

	// Puts the specified character to the top of the stream.
	//
	// Parameters:
	// 	- value: A character to be pushed back.
	Pushback(value rune)

	// Pushes the specified string to the top of the stream.
	//
	// Parameters:
	// 	- value: A string to be pushed back.
	PushbackString(value string)
}
