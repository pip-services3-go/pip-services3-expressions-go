package io

// Wraps string to provide unlimited pushback that allows tokenizers
// to look ahead through stream to perform lexical analysis.
type StringPushbackReader struct {
	content            []rune
	position           int
	pushbackCharsCount int
	pushbackSingleChar rune
	pushbackChars      []rune
}

// Creates an instance of this class.
//
// Parameters:
//   - content: A text content to be read
func NewStringPushbackReader(content string) *StringPushbackReader {
	return &StringPushbackReader{
		content:  []rune(content),
		position: 0,
	}
}

// Reads character from the top of the stream.
//
// Returns a read character or <code>-1</code> if stream processed to the end.
func (c *StringPushbackReader) Read() (rune, error) {
	if c.pushbackCharsCount == 1 {
		c.pushbackCharsCount = c.pushbackCharsCount - 1
		return c.pushbackSingleChar, nil
	} else if c.pushbackCharsCount > 1 {
		result := c.pushbackChars[0]
		c.pushbackChars = c.pushbackChars[1:]
		c.pushbackCharsCount = c.pushbackCharsCount - 1

		if c.pushbackCharsCount == 1 {
			c.pushbackSingleChar = c.pushbackChars[0]
			c.pushbackChars = c.pushbackChars[1:]
		}

		return result, nil
	} else {
		if c.position < len(c.content) {
			c.position = c.position + 1
			return c.content[c.position-1], nil
		}

		return -1, nil
	}
}

// Returns the character from the top of the stream without moving the stream pointer.
//
// Returns a character from the top of the stream or <code>-1</code> if stream is empty.
func (c *StringPushbackReader) Peek() (rune, error) {
	if c.pushbackCharsCount == 1 {
		return c.pushbackSingleChar, nil
	} else if c.pushbackCharsCount > 1 {
		return c.pushbackChars[0], nil
	} else {
		if c.position < len(c.content) {
			return c.content[c.position], nil
		} else {
			return -1, nil
		}
	}
}

// Puts the specified character to the top of the stream.
//
// Parameters:
//   - value: A character to be pushed back.
func (c *StringPushbackReader) Pushback(value rune) {
	// Skip EOF
	if value == -1 {
		return
	}

	if c.pushbackCharsCount == 0 {
		c.pushbackSingleChar = value
	} else if c.pushbackCharsCount == 1 {
		c.pushbackChars = append([]rune{value}, c.pushbackSingleChar)
	} else {
		c.pushbackChars = append([]rune{value}, c.pushbackChars...)
	}
	c.pushbackCharsCount = c.pushbackCharsCount + 1
}

// Pushes the specified string to the top of the stream.
//
// Parameters:
//   - value: A string to be pushed back.
func (c *StringPushbackReader) PushbackString(value string) {
	temp := []rune(value)
	len := len(temp)
	if len > 0 {
		if len == 1 {
			c.Pushback(temp[0])
		} else {
			if c.pushbackCharsCount == 1 {
				c.pushbackChars = append(temp, c.pushbackSingleChar)
			} else {
				c.pushbackChars = append(temp, c.pushbackChars...)
			}
			c.pushbackCharsCount = c.pushbackCharsCount + len
		}
	}
}
