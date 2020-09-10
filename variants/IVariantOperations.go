package variants

// Defines an interface for variant operations manager.
type IVariantOperations interface {
	// Converts variant to specified type
	//
	// Parameters:
	//   - value: A variant value to be converted.
	//   - newType: A type of object to be returned.
	// Returns: A converted Variant value.
	Convert(value *Variant, newType int) *Variant

	// Performs '+' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Add(value1 *Variant, value2 *Variant) *Variant

	// Performs '-' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Sub(value1 *Variant, value2 *Variant) *Variant

	// Performs '*' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Mul(value1 *Variant, value2 *Variant) *Variant

	// Performs '/' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Div(value1 *Variant, value2 *Variant) *Variant

	// Performs '%' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Mod(value1 *Variant, value2 *Variant) *Variant

	// Performs '^' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Pow(value1 *Variant, value2 *Variant) *Variant

	// Performs AND operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	And(value1 *Variant, value2 *Variant) *Variant

	// Performs OR operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Or(value1 *Variant, value2 *Variant) *Variant

	// Performs XOR operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Xor(value1 *Variant, value2 *Variant) *Variant

	// Performs << operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Lsh(value1 *Variant, value2 *Variant) *Variant

	// Performs >> operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Rsh(value1 *Variant, value2 *Variant) *Variant

	// Performs NOT operation for a variant.
	//
	// Parameters:
	//   - value: The operand for this operation.
	// Returns: A result variant object.
	Not(value *Variant) *Variant

	// Performs unary '-' operation for a variant.
	//
	// Parameters:
	//   - value: The operand for this operation.
	// Returns: A result variant object.
	Negative(value *Variant) *Variant

	// Performs '=' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Equal(value1 *Variant, value2 *Variant) *Variant

	// Performs '<>' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	NotEqual(value1 *Variant, value2 *Variant) *Variant

	// Performs '>' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	More(value1 *Variant, value2 *Variant) *Variant

	// Performs '<' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	Less(value1 *Variant, value2 *Variant) *Variant

	// Performs '>=' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	MoreEqual(value1 *Variant, value2 *Variant) *Variant

	// Performs '<=' operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	LessEqual(value1 *Variant, value2 *Variant) *Variant

	// Performs IN operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	In(value1 *Variant, value2 *Variant) *Variant

	// Performs [] operation for two variants.
	//
	// Parameters:
	//   - value1: The first operand for this operation.
	//   - value2: The second operand for this operation.
	// Returns: A result variant object.
	GetElement(value1 *Variant, value2 *Variant) *Variant
}
