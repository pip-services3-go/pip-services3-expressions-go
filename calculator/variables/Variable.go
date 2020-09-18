package variables

import "github.com/pip-services3-go/pip-services3-expressions-go/variants"

// Implements a variable holder object.
type Variable struct {
	name  string
	value *variants.Variant
}

// Constructs a new empty variable.
//
// Parameters:
//   - name: The name of this variable.
func EmptyVariable(name string) *Variable {
	return NewVariable(name, nil)
}

// Constructs this variable with name and value.
//
// Parameters:
//   - name: The name of this variable.
//   - value: The variable value.
func NewVariable(name string, value *variants.Variant) *Variable {
	if name == "" {
		panic("Name parameter cannot be empty")
	}
	if value == nil {
		value = variants.EmptyVariant()
	}
	c := &Variable{
		name:  name,
		value: value,
	}
	return c
}

// The variable name.
func (c *Variable) Name() string {
	return c.name
}

// Gets the variable value.
func (c *Variable) Value() *variants.Variant {
	return c.value
}

// Sets the variable value.
func (c *Variable) SetValue(value *variants.Variant) {
	c.value = value
}