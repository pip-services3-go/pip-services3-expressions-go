package functions

// Defines a functions list.
type IFunctionCollection interface {
	// Adds a new function to the collection.
	//
	// Parameters:
	//   - function: a function to be added.
	Add(function IFunction)

	// A number of functions stored in the collection.
	Length() int

	// Get a function by its index.
	//
	// Parameters:
	//   - index: a function index.
	// Returns: a retrieved function.
	Get(index int) IFunction

	// Get all functions stores in the collection
	//
	// Returns: a list with functions.
	GetAll() []IFunction

	// Finds function index in the list by it's name.
	//
	// Parameters:
	//   - name: The function name to be found.
	// Returns: Function index in the list or <code>-1</code> if function was not found.
	FindIndexByName(name string) int

	// Finds function in the list by it's name.
	//
	// Parameters:
	//   - name: The function name to be found.
	// Returns: Function or <code>null</code> if function was not found.
	FindByName(name string) IFunction

	// Removes a function by its index.
	//
	// Parameters:
	//   - index: a index of the function to be removed.
	Remove(index int)

	// Removes function by it's name.
	//
	// Parameters:
	// 	 - name: The function name to be removed.
	RemoveByName(name string)

	// Clears the collection.
	Clear()
}
