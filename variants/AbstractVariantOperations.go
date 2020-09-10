package variants

// Implements an abstractd variant operations manager object.
type AbstractVariantOperations struct {
	convertFunc func(value *Variant, newType int) *Variant
}

// Convert variant type to string representation
//
// Parameters:
//   - value: a variant type to be converted.
// Returns: a string representation of the type.
func typeToString(value int) string {
	switch value {
	case Null:
		return "Null"
	case Integer:
		return "Integer"
	case Long:
		return "Long"
	case Float:
		return "Float"
	case Double:
		return "Double"
	case String:
		return "String"
	case Boolean:
		return "Boolean"
	case DateTime:
		return "DateTime"
	case TimeSpan:
		return "TimeSpan"
	case Object:
		return "Object"
	case Array:
		return "Array"
	default:
		return "Unknown"
	}
}

// Converts variant to specified type
//
// Parameters:
//   - value: A variant value to be converted.
//   - newType: A type of object to be returned.
// Returns: A converted Variant value.
func (c *AbstractVariantOperations) Convert(
	value *Variant, newType int) *Variant {
	return c.convertFunc(value, newType)
	//panic("Not implemented operation")
}

// Performs '+' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Add(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() + value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() + value2.AsLong())
		return result
	case Float:
		result.SetAsFloat(value1.AsFloat() + value2.AsFloat())
		return result
	case Double:
		result.SetAsDouble(value1.AsDouble() + value2.AsDouble())
		return result
	case TimeSpan:
		result.SetAsTimeSpan(value1.AsTimeSpan() + value2.AsTimeSpan())
		return result
	case String:
		result.SetAsString(value1.AsString() + value2.AsString())
		return result
	}
	panic("Operation '+' is not supported for type " + typeToString(value1.Type()))
}

// Performs '-' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Sub(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() - value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() - value2.AsLong())
		return result
	case Float:
		result.SetAsFloat(value1.AsFloat() - value2.AsFloat())
		return result
	case Double:
		result.SetAsDouble(value1.AsDouble() - value2.AsDouble())
		return result
	case TimeSpan:
		result.SetAsTimeSpan(value1.AsTimeSpan() - value2.AsTimeSpan())
		return result
	case DateTime:
		result.SetAsTimeSpan(value1.AsDateTime().Sub(value2.AsDateTime()))
		return result
	}
	panic("Operation '-' is not supported for type " + typeToString(value1.Type()))
}

// Performs '*' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Mul(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() * value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() * value2.AsLong())
		return result
	case Float:
		result.SetAsFloat(value1.AsFloat() * value2.AsFloat())
		return result
	case Double:
		result.SetAsDouble(value1.AsDouble() * value2.AsDouble())
		return result
	}
	panic("Operation '*' is not supported for type " + typeToString(value1.Type()))
}

// Performs '/' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Div(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() / value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() / value2.AsLong())
		return result
	case Float:
		result.SetAsFloat(value1.AsFloat() / value2.AsFloat())
		return result
	case Double:
		result.SetAsDouble(value1.AsDouble() / value2.AsDouble())
		return result
	}
	panic("Operation '/' is not supported for type " + typeToString(value1.Type()))
}

// Performs '%' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Mod(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() % value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() % value2.AsLong())
		return result
	}
	panic("Operation '%' is not supported for type " + typeToString(value1.Type()))
}

// Performs '^' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Pow(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Performs operation.
	switch value1.Type() {
	case Integer:
	case Long:
	case Float:
	case Double:
		// Converts second operant to the type of the first operand.
		value1 = c.Convert(value1, Double)
		value2 = c.Convert(value2, Double)
		result.SetAsDouble(value1.AsDouble() * value2.AsDouble())
		return result
	}
	panic("Operation '^' is not supported for type " + typeToString(value1.Type()))
}

// Performs AND operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) And(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() & value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() & value2.AsLong())
		return result
	case Boolean:
		result.SetAsBoolean(value1.AsBoolean() && value2.AsBoolean())
		return result
	}
	panic("Operation AND is not supported for type " + typeToString(value1.Type()))
}

// Performs OR operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Or(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() | value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() | value2.AsLong())
		return result
	case Boolean:
		result.SetAsBoolean(value1.AsBoolean() || value2.AsBoolean())
		return result
	}
	panic("Operation OR is not supported for type " + typeToString(value1.Type()))
}

// Performs XOR operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Xor(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() ^ value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() ^ value2.AsLong())
		return result
	case Boolean:
		result.SetAsBoolean((value1.AsBoolean() && !value2.AsBoolean()) ||
			(!value1.AsBoolean() && value2.AsBoolean()))
		return result
	}
	panic("Operation XOR is not supported for type " + typeToString(value1.Type()))
}

// Performs '<<' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Lsh(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, Integer)

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() << value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() << value2.AsInteger())
		return result
	}
	panic("Operation '<<' is not supported for type " + typeToString(value1.Type()))
}

// Performs '>>' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Rsh(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, Integer)

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsInteger(value1.AsInteger() >> value2.AsInteger())
		return result
	case Long:
		result.SetAsLong(value1.AsLong() >> value2.AsInteger())
		return result
	}
	panic("Operation '>>' is not supported for type " + typeToString(value1.Type()))
}

// Performs NOT operation for a variant.
//
// Parameters:
//   - value: The operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Not(value *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value.Type() == Null {
		result.SetAsBoolean(true)
		return result
	}

	// Performs operation.
	switch value.Type() {
	case Integer:
		result.SetAsInteger(^value.AsInteger())
		return result
	case Long:
		result.SetAsLong(^value.AsLong())
		return result
	case Boolean:
		result.SetAsBoolean(!value.AsBoolean())
		return result
	}
	panic("Operation NOT is not supported for type " + typeToString(value.Type()))
}

// Performs unary '-' operation for a variant.
//
// Parameters:
//   - value: The operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Negative(value *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value.Type() == Null {
		return result
	}

	// Performs operation.
	switch value.Type() {
	case Integer:
		result.SetAsInteger(-value.AsInteger())
		return result
	case Long:
		result.SetAsLong(-value.AsLong())
		return result
	case Float:
		result.SetAsFloat(-value.AsFloat())
		return result
	case Double:
		result.SetAsDouble(-value.AsDouble())
		return result
	}
	panic("Operation unary '-' is not supported for type " + typeToString(value.Type()))
}

// Performs '=' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Equal(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null && value2.Type() == Null {
		result.SetAsBoolean(true)
		return result
	}
	if value1.Type() == Null || value2.Type() == Null {
		result.SetAsBoolean(false)
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() == value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() == value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() == value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() == value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() == value2.AsString())
		return result
	case Boolean:
		result.SetAsBoolean(value1.AsBoolean() == value2.AsBoolean())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() == value2.AsTimeSpan())
		return result
	case DateTime:
		date1 := value1.AsDateTime()
		date2 := value2.AsDateTime()
		result.SetAsBoolean(date1.Equal(date2))
		return result
	case Object:
		result.SetAsBoolean(value1.AsObject() == value2.AsObject())
		return result
	}
	panic("Operation '=' is not supported for type " + typeToString(value1.Type()))
}

// Performs '<>' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) NotEqual(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null && value2.Type() == Null {
		result.SetAsBoolean(false)
		return result
	}
	if value1.Type() == Null || value2.Type() == Null {
		result.SetAsBoolean(true)
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() != value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() != value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() != value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() != value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() != value2.AsString())
		return result
	case Boolean:
		result.SetAsBoolean(value1.AsBoolean() != value2.AsBoolean())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() != value2.AsTimeSpan())
		return result
	case DateTime:
		date1 := value1.AsDateTime()
		date2 := value2.AsDateTime()
		result.SetAsBoolean(!date1.Equal(date2))
		return result
	case Object:
		result.SetAsBoolean(value1.AsObject() != value2.AsObject())
		return result
	}
	panic("Operation '<>' is not supported for type " + typeToString(value1.Type()))
}

// Performs '>' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) More(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() > value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() > value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() > value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() > value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() > value2.AsString())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() > value2.AsTimeSpan())
		return result
	case DateTime:
		result.SetAsBoolean(value1.AsDateTime().After(value2.AsDateTime()))
		return result
	}
	panic("Operation '>' is not supported for type " + typeToString(value1.Type()))
}

// Performs '<' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) Less(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() < value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() < value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() < value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() < value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() < value2.AsString())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() < value2.AsTimeSpan())
		return result
	case DateTime:
		result.SetAsBoolean(value1.AsDateTime().Before(value2.AsDateTime()))
		return result
	}
	panic("Operation '<' is not supported for type " + typeToString(value1.Type()))
}

// Performs '>=' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) MoreEqual(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() >= value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() >= value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() >= value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() >= value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() >= value2.AsString())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() >= value2.AsTimeSpan())
		return result
	case DateTime:
		date1 := value1.AsDateTime()
		date2 := value2.AsDateTime()
		result.SetAsBoolean(date1.After(date2) || date1.Equal(date2))
		return result
	}
	panic("Operation '>=' is not supported for type " + typeToString(value1.Type()))
}

// Performs '<=' operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) LessEqual(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Converts second operant to the type of the first operand.
	value2 = c.Convert(value2, value1.Type())

	// Performs operation.
	switch value1.Type() {
	case Integer:
		result.SetAsBoolean(value1.AsInteger() <= value2.AsInteger())
		return result
	case Long:
		result.SetAsBoolean(value1.AsLong() <= value2.AsLong())
		return result
	case Float:
		result.SetAsBoolean(value1.AsFloat() <= value2.AsFloat())
		return result
	case Double:
		result.SetAsBoolean(value1.AsDouble() <= value2.AsDouble())
		return result
	case String:
		result.SetAsBoolean(value1.AsString() <= value2.AsString())
		return result
	case TimeSpan:
		result.SetAsBoolean(value1.AsTimeSpan() <= value2.AsTimeSpan())
		return result
	case DateTime:
		date1 := value1.AsDateTime()
		date2 := value2.AsDateTime()
		result.SetAsBoolean(date1.Before(date2) || date1.Equal(date2))
		return result
	}
	panic("Operation '<=' is not supported for type " + typeToString(value1.Type()))
}

// Performs IN operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) In(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	// Processes null arrays.
	if value1.AsObject() == nil {
		result.SetAsBoolean(false)
		return result
	}

	if value1.Type() == Array {
		array := value1.AsArray()
		for _, element := range array {
			eq := c.Equal(value2, element)
			if eq.Type() == Boolean && eq.AsBoolean() {
				result.SetAsBoolean(true)
				return result
			}
		}
		result.SetAsBoolean(false)
		return result
	}
	return c.Equal(value1, value2)
}

// Performs [] operation for two variants.
//
// Parameters:
//   - value1: The first operand for this operation.
//   - value2: The second operand for this operation.
// Returns: A result variant object.
func (c *AbstractVariantOperations) GetElement(
	value1 *Variant, value2 *Variant) *Variant {
	result := EmptyVariant()

	// Processes VariantType.Null values.
	if value1.Type() == Null || value2.Type() == Null {
		return result
	}

	value2 = c.Convert(value2, Integer)
	index := int(value2.AsInteger())

	if value1.Type() == Array {
		return value1.GetByIndex(index)
	} else if value1.Type() == String {
		runes := []rune(value1.AsString())
		result.SetAsString(string(runes[value2.AsInteger()]))
		return result
	}
	panic("Operation '[]' is not supported for type " + typeToString(value1.Type()))
}
