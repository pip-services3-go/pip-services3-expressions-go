package functions

import (
	"fmt"

	"github.com/pip-services3-go/pip-services3-expressions-go/calculator/errors"
	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
)

// Defines a delegate to implement a function
//
// Parameters:
//   - parameters: A list with function parameters
//   - variantOperations: A manager for variant operations.
// Returns: A calculated function value.
type FunctionCalculator func(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error)

// Defines an interface for expression function.
type DelegatedFunction struct {
	name       string
	calculator FunctionCalculator
}

// Constructs this function class with specified parameters.
//
// Parameters:
//   - name: The name of this function.
//   - calculator: The function calculator delegate.
func NewDelegatedFunction(name string, calculator FunctionCalculator) *DelegatedFunction {
	if name == "" {
		panic("Name parameter cannot be empty.")
	}
	if calculator == nil {
		panic("Calculator parameter cannot be nil.")
	}

	c := &DelegatedFunction{
		name:       name,
		calculator: calculator,
	}
	return c
}

// The function name.
func (c *DelegatedFunction) Name() string {
	return c.name
}

// The function calculation method.
//
// Parameters:
//   - parameters: A list with function parameters.
//   - variantOperations: Variants operations manager.
// Returns: A calculated function result.
func (c *DelegatedFunction) Calculate(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	var result *variants.Variant
	var err error

	// Capture calculation error
	defer func() {
		if r := recover(); r != nil {
			message := fmt.Sprint("%v", r)
			err = errors.NewExpressionError("", "CALC_FAILED", message)
		}
	}()

	result, err = c.calculator(parameters, variantOperations)

	return result, err
}
