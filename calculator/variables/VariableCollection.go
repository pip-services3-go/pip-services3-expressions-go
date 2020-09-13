package variables

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
)

// Implements a variables list.
type VariableCollection struct {
	variables []IVariable
}

func NewVariableCollection() *VariableCollection {
	c := &VariableCollection{
		variables: []IVariable{},
	}
	return c
}

// Adds a new variable to the collection.
//
// Parameters:
//   - variable: a variable to be added.
func (c *VariableCollection) Add(variable IVariable) {
	if variable == nil {
		panic("Variable cannot be null")
	}
	c.variables = append(c.variables, variable)
}

// A number of variables stored in the collection.
func (c *VariableCollection) Length() int {
	return len(c.variables)
}

// Get a variable by its index.
//
// Parameters:
//   - index: a variable index.
// Returns: a retrieved variable.
func (c *VariableCollection) Get(index int) IVariable {
	return c.variables[index]
}

// Get all variables stores in the collection
//
// Returns: a list with variables.
func (c *VariableCollection) GetAll() []IVariable {
	result := []IVariable{}
	result = append(result, c.variables...)
	return result
}

// Finds variable index in the list by it's name.
//
// Parameters:
//   - name: The variable name to be found.
// Returns: Variable index in the list or <code>-1</code> if variable was not found.
func (c *VariableCollection) FindIndexByName(name string) int {
	name = strings.ToUpper(name)
	for i, v := range c.variables {
		if strings.ToUpper(v.Name()) == name {
			return i
		}
	}
	return -1
}

// Finds variable in the list by it's name.
//
// Parameters:
//  - name: The variable name to be found.
// Returns: Variable or <code>null</code> if function was not found.
func (c *VariableCollection) FindByName(name string) IVariable {
	index := c.FindIndexByName(name)
	if index >= 0 {
		return c.variables[index]
	}
	return nil
}

// Finds variable in the list or create a new one if variable was not found.
//
// Parameters:
//   - name: The variable name to be found.
// Returns: Found or created variable.
func (c *VariableCollection) Locate(name string) IVariable {
	v := c.FindByName(name)
	if v == nil {
		v = EmptyVariable(name)
		c.variables = append(c.variables, v)
	}
	return v
}

// Removes a variable by its index.
//
// Parameters:
//   - index: a index of the variable to be removed.
func (c *VariableCollection) Remove(index int) {
	c.variables = append(c.variables[:index], c.variables[index+1:]...)
}

// Removes variable by it's name.
//
// Parameters:
//   - name: The variable name to be removed.
func (c *VariableCollection) RemoveByName(name string) {
	index := c.FindIndexByName(name)
	if index >= 0 {
		c.variables = append(c.variables[:index], c.variables[index+1:]...)
	}
}

// Clears the collection.
func (c *VariableCollection) Clear() {
	c.variables = []IVariable{}
}

// Clears all stored variables (assigns null values).
func (c *VariableCollection) ClearValues() {
	for _, v := range c.variables {
		v.SetValue(variants.EmptyVariant())
	}
}
