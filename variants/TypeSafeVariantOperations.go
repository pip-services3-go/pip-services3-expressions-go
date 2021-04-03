package variants

import "github.com/pip-services3-go/pip-services3-commons-go/errors"

// Implements a strongly typed (type safe) variant operations manager object.
type TypeSafeVariantOperations struct {
	AbstractVariantOperations
}

func NewTypeSafeVariantOperations() *TypeSafeVariantOperations {
	c := &TypeSafeVariantOperations{}
	c.AbstractVariantOperations = *InheritAbstractVariantOperations(c)
	return c
}

// Converts variant to specified type
//
// Parameters:
//   - value: A variant value to be converted.
//   - newType: A type of object to be returned.
// Returns: A converted Variant value.
func (c *TypeSafeVariantOperations) Convert(
	value *Variant, newType int) (*Variant, error) {

	if newType == Null {
		result := EmptyVariant()
		return result, nil
	}
	if newType == value.Type() || newType == Object {
		return value, nil
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
		return value, nil
	case Array:
		break
	}

	err := errors.NewUnsupportedError("", "CONV_NOT_SUPPORTED",
		"Variant convertion from "+typeToString(value.Type())+
			" to "+typeToString(newType)+" is not supported.")
	return nil, err
}

func (c *TypeSafeVariantOperations) convertFromInteger(
	value *Variant, newType int) (*Variant, error) {

	result := EmptyVariant()
	switch newType {
	case Long:
		result.SetAsLong(int64(value.AsInteger()))
		return result, nil
	case Float:
		result.SetAsFloat(float32(value.AsInteger()))
		return result, nil
	case Double:
		result.SetAsDouble(float64(value.AsInteger()))
		return result, nil
	}

	err := errors.NewUnsupportedError("", "CONV_NOT_SUPPORTED",
		"Variant convertion from "+typeToString(value.Type())+
			" to "+typeToString(newType)+" is not supported.")
	return nil, err
}

func (c *TypeSafeVariantOperations) convertFromLong(
	value *Variant, newType int) (*Variant, error) {

	result := EmptyVariant()
	switch newType {
	case Float:
		result.SetAsFloat(float32(value.AsLong()))
		return result, nil
	case Double:
		result.SetAsDouble(float64(value.AsLong()))
		return result, nil
	}

	err := errors.NewUnsupportedError("", "CONV_NOT_SUPPORTED",
		"Variant convertion from "+typeToString(value.Type())+
			" to "+typeToString(newType)+" is not supported.")
	return nil, err
}

func (c *TypeSafeVariantOperations) convertFromFloat(
	value *Variant, newType int) (*Variant, error) {

	result := EmptyVariant()
	switch newType {
	case Double:
		result.SetAsDouble(float64(value.AsFloat()))
		return result, nil
	}

	err := errors.NewUnsupportedError("", "CONV_NOT_SUPPORTED",
		"Variant convertion from "+typeToString(value.Type())+
			" to "+typeToString(newType)+" is not supported.")
	return nil, err
}
