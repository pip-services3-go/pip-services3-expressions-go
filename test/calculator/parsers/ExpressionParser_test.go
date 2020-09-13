package test_calculator_parsers

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/calculator/parsers"
	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
	"github.com/stretchr/testify/assert"
)

func TestExpressionParserParseString(t *testing.T) {
	parser := parsers.NewExpressionParser()
	err := parser.SetExpression("(2+2)*ABS(-2)")
	assert.Nil(t, err)

	expectedTokens := []*parsers.ExpressionToken{
		parsers.NewExpressionToken(parsers.Constant, variants.VariantFromInteger(2)),
		parsers.NewExpressionToken(parsers.Constant, variants.VariantFromInteger(2)),
		parsers.NewExpressionToken(parsers.Plus, variants.Empty),
		parsers.NewExpressionToken(parsers.Constant, variants.VariantFromInteger(2)),
		parsers.NewExpressionToken(parsers.Unary, variants.Empty),
		parsers.NewExpressionToken(parsers.Constant, variants.VariantFromInteger(1)),
		parsers.NewExpressionToken(parsers.Function, variants.VariantFromString("ABS")),
		parsers.NewExpressionToken(parsers.Star, variants.Empty),
	}

	tokens := parser.ResultTokens()
	assert.Equal(t, len(expectedTokens), len(tokens))

	for i := 0; i < len(tokens); i++ {
		assert.Equal(t, expectedTokens[i].Type(), tokens[i].Type())
		assert.Equal(t, expectedTokens[i].Value().Type(), tokens[i].Value().Type())
		assert.Equal(t, expectedTokens[i].Value().AsObject(), tokens[i].Value().AsObject())
	}
}
