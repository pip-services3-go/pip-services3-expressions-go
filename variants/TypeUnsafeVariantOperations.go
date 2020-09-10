package variants

import (
	"math"
	"time"

	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
)

// Implements a type unsafe variant operations manager object.
type TypeUnsafeVariantOperations struct {
	AbstractVariantOperations
}

func NewTypeUnsafeVariantOperations() *TypeUnsafeVariantOperations {
	c := &TypeUnsafeVariantOperations{
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
func (c *TypeUnsafeVariantOperations) Convert(value *Variant, newType int) *Variant {
	if newType == Null {
		result := EmptyVariant()
		return result
	}
	if newType == value.Type() || newType == Object {
		return value
	}
	if newType == String {
		result := EmptyVariant()
		result.SetAsString(cconv.StringConverter.ToString(value.AsObject()))
		return result
	}

	switch value.Type() {
	case Null:
		return c.convertFromNull(newType)
	case Integer:
		return c.convertFromInteger(value, newType)
	case Long:
		return c.convertFromLong(value, newType)
	case Float:
		return c.convertFromFloat(value, newType)
	case Double:
		return c.convertFromDouble(value, newType)
	case DateTime:
		return c.convertFromDateTime(value, newType)
	case TimeSpan:
		return c.convertFromTimeSpan(value, newType)
	case String:
		return c.convertFromString(value, newType)
	case Boolean:
		return c.convertFromBoolean(value, newType)
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromNull(newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(0)
		return result
	case Long:
		result.SetAsLong(0)
		return result
	case Float:
		result.SetAsFloat(0)
		return result
	case Double:
		result.SetAsDouble(0)
		return result
	case Boolean:
		result.SetAsBoolean(false)
		return result
	case DateTime:
		result.SetAsDateTime(time.Time{})
		return result
	case TimeSpan:
		result.SetAsTimeSpan(time.Duration(0 * time.Millisecond))
		return result
	case String:
		result.SetAsString("null")
		return result
	case Object:
		result.SetAsObject(nil)
		return result
	case Array:
		result.SetAsArray([]*Variant{})
		return result
	}
	panic("Variant convertion from Null " +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromInteger(value *Variant, newType int) *Variant {
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
	case DateTime:
		result.SetAsDateTime(time.Unix(int64(value.AsInteger()), 0))
		return result
	case TimeSpan:
		result.SetAsTimeSpan(time.Duration(value.AsInteger()) * time.Millisecond)
		return result
	case Boolean:
		result.SetAsBoolean(value.AsInteger() != 0)
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromLong(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(value.AsLong()))
		return result
	case Float:
		result.SetAsFloat(float32(value.AsLong()))
		return result
	case Double:
		result.SetAsDouble(float64(value.AsLong()))
		return result
	case DateTime:
		result.SetAsDateTime(time.Unix(value.AsLong(), 0))
		return result
	case TimeSpan:
		result.SetAsTimeSpan(time.Duration(value.AsLong() * time.Hour.Milliseconds()))
		return result
	case Boolean:
		result.SetAsBoolean(value.AsLong() != 0)
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromFloat(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(math.Trunc(float64(value.AsFloat()))))
		return result
	case Long:
		result.SetAsLong(int64(math.Trunc(float64(value.AsFloat()))))
		return result
	case Double:
		result.SetAsDouble(float64(value.AsFloat()))
		return result
	case Boolean:
		result.SetAsBoolean(value.AsFloat() != 0)
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromDouble(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(math.Trunc(value.AsDouble())))
		return result
	case Long:
		result.SetAsLong(int64(math.Trunc(value.AsDouble())))
		return result
	case Float:
		result.SetAsFloat(float32(value.AsDouble()))
		return result
	case Boolean:
		result.SetAsBoolean(value.AsDouble() != 0)
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromString(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(cconv.IntegerConverter.ToInteger(value.AsString())))
		return result
	case Long:
		result.SetAsLong(int64(cconv.LongConverter.ToLong(value.AsString())))
		return result
	case Float:
		result.SetAsFloat(cconv.FloatConverter.ToFloat(value.AsString()))
		return result
	case Double:
		result.SetAsDouble(cconv.DoubleConverter.ToDouble(value.AsString()))
		return result
	case DateTime:
		result.SetAsDateTime(cconv.DateTimeConverter.ToDateTime(value.AsString()))
		return result
	case TimeSpan:
		result.SetAsTimeSpan(cconv.DurationConverter.ToDuration(value.AsString()))
		return result
	case Boolean:
		result.SetAsBoolean(cconv.BooleanConverter.ToBoolean(value.AsString()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromBoolean(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		if value.AsBoolean() {
			result.SetAsInteger(1)
		} else {
			result.SetAsInteger(0)
		}
		return result
	case Long:
		if value.AsBoolean() {
			result.SetAsLong(1)
		} else {
			result.SetAsLong(0)
		}
		return result
	case Float:
		if value.AsBoolean() {
			result.SetAsFloat(1)
		} else {
			result.SetAsFloat(0)
		}
		return result
	case Double:
		if value.AsBoolean() {
			result.SetAsDouble(1)
		} else {
			result.SetAsDouble(0)
		}
		return result
	case String:
		if value.AsBoolean() {
			result.SetAsString("true")
		} else {
			result.SetAsString("false")
		}
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromDateTime(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(value.AsDateTime().Unix()))
		return result
	case Long:
		result.SetAsLong(value.AsDateTime().Unix())
		return result
	case String:
		result.SetAsString(cconv.StringConverter.ToString(value.AsDateTime()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}

func (c *TypeUnsafeVariantOperations) convertFromTimeSpan(value *Variant, newType int) *Variant {
	result := EmptyVariant()
	switch newType {
	case Integer:
		result.SetAsInteger(int32(value.AsTimeSpan().Milliseconds()))
		return result
	case Long:
		result.SetAsLong(value.AsTimeSpan().Milliseconds())
		return result
	case String:
		result.SetAsString(cconv.StringConverter.ToString(value.AsTimeSpan()))
		return result
	}
	panic("Variant convertion from " + typeToString(value.Type()) +
		" to " + typeToString(newType) + " is not supported.")
}
