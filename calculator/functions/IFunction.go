package functions

import "github.com/pip-services3-go/pip-services3-expressions-go/variants"

// Defines an interface for expression function.
// </summary>
type IFunction interface {
	// The function name.
	Name() string

	// The function calculation method.
	//
	// Parameters:
	//   - parameters: A list with function parameters<
	//   - variantOperations: Variants operations manager.
	Calculate(parameters []*variants.Variant,
		variantOperations variants.IVariantOperations) (*variants.Variant, error)
}
