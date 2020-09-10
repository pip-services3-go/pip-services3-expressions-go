package test_generic

import (
	"testing"

	test_tokenizers "github.com/pip-services3-go/pip-services3-expressions-go/test/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenericTokenizerExpression(t *testing.T) {
	tokenString := "A+B/123 - \t 'xyz'\n <>-10.11# This is a comment"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Word, "A"),
		tokenizers.NewToken(tokenizers.Symbol, "+"),
		tokenizers.NewToken(tokenizers.Word, "B"),
		tokenizers.NewToken(tokenizers.Symbol, "/"),
		tokenizers.NewToken(tokenizers.Integer, "123"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
		tokenizers.NewToken(tokenizers.Whitespace, " \t "),
		tokenizers.NewToken(tokenizers.Quoted, "'xyz'"),
		tokenizers.NewToken(tokenizers.Whitespace, "\n "),
		tokenizers.NewToken(tokenizers.Symbol, "<>"),
		tokenizers.NewToken(tokenizers.Float, "-10.11"),
		tokenizers.NewToken(tokenizers.Comment, "# This is a comment"),
		tokenizers.NewToken(tokenizers.Eof, ""),
	}

	tokenizer := generic.NewGenericTokenizer()
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestGenericTokenizerQuoteToken(t *testing.T) {
	tokenString := "A'xyz'\"abc\ndeg\" 'jkl\"def'"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Word, "A"),
		tokenizers.NewToken(tokenizers.Quoted, "xyz"),
		tokenizers.NewToken(tokenizers.Quoted, "abc\ndeg"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Quoted, "jkl\"def"),
	}

	tokenizer := generic.NewGenericTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestGenericTokenizerWordToken(t *testing.T) {
	tokenString := "A'xyz'Ebf_2\n2x_2"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Word, "A"),
		tokenizers.NewToken(tokenizers.Quoted, "xyz"),
		tokenizers.NewToken(tokenizers.Word, "Ebf_2"),
		tokenizers.NewToken(tokenizers.Whitespace, "\n"),
		tokenizers.NewToken(tokenizers.Integer, "2"),
		tokenizers.NewToken(tokenizers.Word, "x_2"),
	}

	tokenizer := generic.NewGenericTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestGenericTokenizerNumberToken(t *testing.T) {
	tokenString := "123-321 .543-.76-. -123.456"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Integer, "123"),
		tokenizers.NewToken(tokenizers.Integer, "-321"),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, ".543"),
		tokenizers.NewToken(tokenizers.Float, "-.76"),
		tokenizers.NewToken(tokenizers.Symbol, "-"),
		tokenizers.NewToken(tokenizers.Symbol, "."),
		tokenizers.NewToken(tokenizers.Whitespace, " "),
		tokenizers.NewToken(tokenizers.Float, "-123.456"),
	}

	tokenizer := generic.NewGenericTokenizer()
	tokenizer.SetSkipEof(true)
	tokenizer.SetDecodeStrings(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestGenericTokenizerWrongToken(t *testing.T) {
	tokenString := "1>2"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Integer, "1"),
		tokenizers.NewToken(tokenizers.Symbol, ">"),
		tokenizers.NewToken(tokenizers.Integer, "2"),
	}

	tokenizer := generic.NewGenericTokenizer()
	tokenizer.SetSkipEof(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}
