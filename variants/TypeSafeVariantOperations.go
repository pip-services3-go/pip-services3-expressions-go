package variants

// Implements a strongly typed (type safe) variant operations manager object.
type TypeSafeVariantOperations struct {
	AbstractVariantOperations
}

func NewTypeSafeVariantOperations() *TypeSafeVariantOperations {
	c := &TypeSafeVariantOperations{
		AbstractVariantOperations: AbstractVariantOperations{},
	}
	c.AbstractVariantOperations.convertFunc = c.Convert
	return c
}

// Converts variant to specified type
//
// Parameters:
//   - value: A variant value to be converted.
//   - newType: A type of object to be returned.
// Returns: A converted Variant value.
func (c *TypeSafeVariantOperations) Convert(value *Variant, newType int) *Variant {
	if newType == Null {
		result := EmptyVariant()
		return result
	}
	if newType == value.Type() || newType == Object {
		return value
	}

	switch value.Type() {
	case Integer:
		return c.convertFromInteger(value, newType)
	case Long:
		return c.convertFromLong(value, newType)
	case Float:
		return c.convertFromFloat(value, newType)
	case Double:
		break
	case String:
		break
	case Boolean:
		break
	case Object:
		return value
	case Array:
		break
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeSafeVariantOperations) convertFromInteger(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Long:
		result.SetAsLong(int64(value.AsInteger()))
		return result
	case Float:
		result.SetAsFloat(float32(value.AsInteger()))
		return result
	case Double:
		result.SetAsDouble(float64(value.AsInteger()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeSafeVariantOperations) convertFromLong(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Float:
		result.SetAsFloat(float32(value.AsLong()))
		return result
	case Double:
		result.SetAsDouble(float64(value.AsLong()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeSafeVariantOperations) convertFromFloat(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Double:
		result.SetAsDouble(float64(value.AsFloat()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}
