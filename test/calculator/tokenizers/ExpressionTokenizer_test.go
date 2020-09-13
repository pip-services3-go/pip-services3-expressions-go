package test_calculator_variables

import (
	"testing"

	ctokenizers "github.com/pip-services3-go/pip-services3-expressions-go/calculator/tokenizers"
	test_tokenizers "github.com/pip-services3-go/pip-services3-expressions-go/test/tokenizers"
	tokenizers "github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/stretchr/testify/assert"
)

func TestExpressionTokenizerQuoteToken(t *testing.T) {
	tokenString := "A'xyz'\"abc\ndeg\" 'jkl\"def'\"ab\"\"de\"'df''er'"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Word, "A"),
		tokenizers.NewToken(tokenizers.Quoted, "xyz"),
		tokenizers.NewToken(tokenizers.Word, "abc\ndeg"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Quoted, "jkl\"def"),
		tokenizers.NewToken(tokenizers.Word, "ab\"de"),
		tokenizers.NewToken(tokenizers.Quoted, "df'er"),
	}

	tokenizer := ctokenizers.NewExpressionTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestExpressionTokenizerWordToken(t *testing.T) {
	tokenString := "A'xyz'Ebf_2\n2_2"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Word, "A"),
		tokenizers.NewToken(tokenizers.Quoted, "xyz"),
		tokenizers.NewToken(tokenizers.Word, "Ebf_2"),
		tokenizers.NewToken(tokenizers.Whitespace, "\n"),
		tokenizers.NewToken(tokenizers.Integer, "2"),
		tokenizers.NewToken(tokenizers.Word, "_2"),
	}

	tokenizer := ctokenizers.NewExpressionTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestExpressionTokenizerNumberToken(t *testing.T) {
	tokenString := "123-321 .543-.76-. 123.456 123e45 543.11E+43 1e 3E-"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Integer, "123"),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
		tokenizers.NewToken(tokenizers.Integer, "321"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, ".543"),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
		tokenizers.NewToken(tokenizers.Float, ".76"),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
		tokenizers.NewToken(tokenizers.Symbol, "."),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, "123.456"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, "123e45"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, "543.11E+43"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Integer, "1"),
		tokenizers.NewToken(tokenizers.Word, "e"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Integer, "3"),
		tokenizers.NewToken(tokenizers.Word, "E"),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
	}

	tokenizer := ctokenizers.NewExpressionTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestExpressionTokenizerExpressionToken(t *testing.T) {
	tokenString := "A + b / (3 - Max(-123, 1)*2)"

	tokenizer := ctokenizers.NewExpressionTokenizer()
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	assert.Len(t, tokenList, 25)
}
