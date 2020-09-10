package utilities

// Validates characters that are processed by Tokenizers.
type TCharValidator struct {
}

var CharValidator *TCharValidator = &TCharValidator{}

const Eof rune = -1

func (c *TCharValidator) IsEof(value rune) bool {
	return value == -1
}

func (c *TCharValidator) IsEol(value rune) bool {
	return value == 10 || value == 13
}

func (c *TCharValidator) IsDigit(value rune) bool {
	return value >= '0' && value <= '9'
}
