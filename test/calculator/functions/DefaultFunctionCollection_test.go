package test_calculator_functions

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/calculator/functions"
	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFunctionsCollection(t *testing.T) {
	collection := functions.NewDefaultFunctionCollection()
	parameters := []*variants.Variant{
		variants.VariantFromInteger(1),
		variants.VariantFromInteger(2),
		variants.VariantFromInteger(3),
	}
	operations := variants.NewTypeUnsafeVariantOperations()

	f := collection.FindByName("sum")
	assert.NotNil(t, f)

	result, err := f.Calculate(parameters, operations)
	assert.Nil(t, err)
	assert.Equal(t, variants.Integer, result.Type())
	assert.Equal(t, 6, result.AsInteger())
}
