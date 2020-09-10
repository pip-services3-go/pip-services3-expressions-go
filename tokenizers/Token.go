package tokenizers

// A token represents a logical chunk of a string. For example, a typical tokenizer would break
// the string "1.23 &lt;= 12.3" into three tokens: the number 1.23, a less-than-or-equal symbol,
// and the number 12.3. A token is a receptacle, and relies on a tokenizer to decide precisely how
// to divide a string into tokens.
type Token struct {
	typ   int
	value string
}

// Constructs this token with type and value.
//
// Parameters:
//   - type: The type of this token.
//   - value: The token string value.
// Returns: Created token
func NewToken(typ int, value string) *Token {
	return &Token{
		typ:   typ,
		value: value,
	}
}

// The token type.
func (c *Token) Type() int {
	return c.typ
}

// The token value.
func (c *Token) Value() string {
	return c.value
}

func (c *Token) Equals(obj interface{}) bool {
	if tok, ok := obj.(Token); ok {
		return c.typ == tok.typ && c.value == tok.value
	}
	return false
}
