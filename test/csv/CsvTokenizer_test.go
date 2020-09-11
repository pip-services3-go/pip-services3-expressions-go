package test_generic

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/csv"
	test_tokenizers "github.com/pip-services3-go/pip-services3-expressions-go/test/tokenizers"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers"
	"github.com/stretchr/testify/assert"
)

func TestCsvTokenizerWithDefaultParameters(t *testing.T) {
	tokenString := "\n\r\"John \"\"Da Man\"\"\",Repici,120 Jefferson St.,Riverside, NJ,08075\r\n" +
		"Stephen,Tyler,\"7452 Terrace \"\"At the Plaza\"\" road\",SomeTown,SD, 91234\r" +
		",Blankman,,SomeTown, SD, 00298\n"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Eol, "\n\r"),
		tokenizers.NewToken(tokenizers.Quoted, "\"John \"\"Da Man\"\"\""),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "Repici"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "120 Jefferson St."),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "Riverside"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, " NJ"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "08075"),
		tokenizers.NewToken(tokenizers.Eol, "\r\n"),
		tokenizers.NewToken(tokenizers.Word, "Stephen"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "Tyler"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Quoted, "\"7452 Terrace \"\"At the Plaza\"\" road\""),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "SomeTown"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "SD"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, " 91234"),
		tokenizers.NewToken(tokenizers.Eol, "\r"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "Blankman"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, "SomeTown"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, " SD"),
		tokenizers.NewToken(tokenizers.Symbol, ","),
		tokenizers.NewToken(tokenizers.Word, " 00298"),
		tokenizers.NewToken(tokenizers.Eol, "\n"),
	}

	tokenizer := csv.NewCsvTokenizer()
	tokenizer.SetSkipEof(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}

func TestCsvTokenizerWithOverridenParameters(t *testing.T) {
	tokenString := "\n\r'John, ''Da Man'''\tRepici\t120 Jefferson St.\tRiverside\t NJ\t08075\r\n" +
		"Stephen\t\"Tyler\"\t'7452 \t\nTerrace ''At the Plaza'' road'\tSomeTown\tSD\t 91234\r" +
		"\tBlankman\t\tSomeTown 'xxx\t'\t SD\t 00298\n"
	expectedTokens := []*tokenizers.Token{
		tokenizers.NewToken(tokenizers.Eol, "\n\r"),
		tokenizers.NewToken(tokenizers.Quoted, "'John, ''Da Man'''"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "Repici"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "120 Jefferson St."),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "Riverside"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, " NJ"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "08075"),
		tokenizers.NewToken(tokenizers.Eol, "\r\n"),
		tokenizers.NewToken(tokenizers.Word, "Stephen"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Quoted, "\"Tyler\""),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Quoted, "'7452 \t\nTerrace ''At the Plaza'' road'"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "SomeTown"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "SD"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, " 91234"),
		tokenizers.NewToken(tokenizers.Eol, "\r"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "Blankman"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, "SomeTown "),
		tokenizers.NewToken(tokenizers.Quoted, "'xxx\t'"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, " SD"),
		tokenizers.NewToken(tokenizers.Symbol, "\t"),
		tokenizers.NewToken(tokenizers.Word, " 00298"),
		tokenizers.NewToken(tokenizers.Eol, "\n"),
	}

	tokenizer := csv.NewCsvTokenizer()
	tokenizer.SetFieldSeparators([]rune{'\t'})
	tokenizer.SetQuoteSymbols([]rune{'\'', '"'})
	tokenizer.SetEndOfLine("\n")
	tokenizer.SetSkipEof(true)
	tokenList, err := tokenizer.TokenizeBuffer(tokenString)
	assert.Nil(t, err)

	test_tokenizers.AssertAreEqualsTokenLists(t, expectedTokens, tokenList)
}
