package variables

import "github.com/pip-services3-go/pip-services3-expressions-go/variants"

// Defines a variable interface.
type IVariable interface {
	// The variable name.
	Name() string

	// Gets the variable value.
	Value() *variants.Variant

	// Sets the variable value.
	SetValue(value *variants.Variant)
}
