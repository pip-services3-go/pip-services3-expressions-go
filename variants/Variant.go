package variants

import (
	"time"

	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
)

// Defines container for variant values.
type Variant struct {
	typ   int
	value interface{}
}

// Empty variant constant
var Empty *Variant = EmptyVariant()

// Constructs an empty variant object
func EmptyVariant() *Variant {
	return &Variant{
		typ:   Null,
		value: nil,
	}
}

// Constructs this class and assignes a value.
//
// Params:
// 	- value: another variant value.
func NewVariant(value interface{}) *Variant {
	c := &Variant{}
	c.SetAsObject(value)
	return c
}

// Creates a new variant from Integer value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromInteger(value int) *Variant {
	c := &Variant{}
	c.SetAsInteger(value)
	return c
}

// Creates a new variant from Long value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromLong(value int64) *Variant {
	c := &Variant{}
	c.SetAsLong(value)
	return c
}

// Creates a new variant from Boolean value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromBoolean(value bool) *Variant {
	c := &Variant{}
	c.SetAsBoolean(value)
	return c
}

// Creates a new variant from Float value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromFloat(value float32) *Variant {
	c := &Variant{}
	c.SetAsFloat(value)
	return c
}

// Creates a new variant from Double value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromDouble(value float64) *Variant {
	c := &Variant{}
	c.SetAsDouble(value)
	return c
}

// Creates a new variant from String value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromString(value string) *Variant {
	c := &Variant{}
	c.SetAsString(value)
	return c
}

// Creates a new variant from DateTime value.
//
// Params:
//   - value: a variant value.
// Returns: A created variant object
func VariantFromDateTime(value time.Time) *Variant {
	c := &Variant{}
	c.SetAsDateTime(value)
	return c
}

// Creates a new variant from TimeSpan value.
//
// Params:
//   - value: a variant value.
// Rturns: A created variant object
func VariantFromTimeSpan(value time.Duration) *Variant {
	c := &Variant{}
	c.SetAsTimeSpan(value)
	return c
}

// Creates a new variant from Object value.
//
// Params:
//   - value: a variant value.
// Rturns: A created variant object
func VariantFromObject(value interface{}) *Variant {
	c := &Variant{}
	c.SetAsObject(value)
	return c
}

// Creates a new variant from Array value.
//
// Params:
//   - value: a variant value.
// Rturns: A created variant object
func VariantFromArray(value []*Variant) *Variant {
	c := &Variant{}
	c.SetAsArray(value)
	return c
}

// Gets a variant type
func (c *Variant) Type() int {
	return c.typ
}

// Gets variant value as integer
func (c *Variant) AsInteger() int {
	return c.value.(int)
}

// Sets variant value as integer
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsInteger(value int) {
	c.typ = Integer
	c.value = value
}

// Gets variant value as int64
func (c *Variant) AsLong() int64 {
	return c.value.(int64)
}

// Sets variant value as int64
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsLong(value int64) {
	c.typ = Long
	c.value = value
}

// Gets variant value as boolean
func (c *Variant) AsBoolean() bool {
	return c.value.(bool)
}

// Sets variant value as boolean
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsBoolean(value bool) {
	c.typ = Boolean
	c.value = value
}

// Gets variant value as float
func (c *Variant) AsFloat() float32 {
	return c.value.(float32)
}

// Sets variant value as float
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsFloat(value float32) {
	c.typ = Float
	c.value = value
}

// Gets variant value as double
func (c *Variant) AsDouble() float64 {
	return c.value.(float64)
}

// Sets variant value as double
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsDouble(value float64) {
	c.typ = Double
	c.value = value
}

// Gets variant value as string
func (c *Variant) AsString() string {
	return c.value.(string)
}

// Sets variant value as string
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsString(value string) {
	c.typ = String
	c.value = value
}

// Gets variant value as DateTime
func (c *Variant) AsDateTime() time.Time {
	return c.value.(time.Time)
}

// Sets variant value as DateTime
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsDateTime(value time.Time) {
	c.typ = DateTime
	c.value = value
}

// Gets variant value as TimeSpan
func (c *Variant) AsTimeSpan() time.Duration {
	return c.value.(time.Duration)
}

// Sets variant value as TimeSpan
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsTimeSpan(value time.Duration) {
	c.typ = TimeSpan
	c.value = value
}

// Gets variant value as object
func (c *Variant) AsObject() interface{} {
	return c.value
}

// Sets variant value as object
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsObject(value interface{}) {
	c.value = value

	if value == nil {
		c.typ = Null
		return
	}

	switch value.(type) {
	case int:
		c.typ = Integer
	case int32:
		c.value = int(c.value.(int32))
		c.typ = Integer
	case uint:
		c.value = int64(c.value.(uint))
		c.typ = Long
	case uint32:
		c.value = int64(c.value.(uint32))
		c.typ = Long
	case int64:
		c.typ = Long
	case float32:
		c.typ = Float
	case float64:
		c.typ = Double
	case bool:
		c.typ = Boolean
	case time.Time:
		c.typ = DateTime
	case time.Duration:
		c.typ = TimeSpan
	case string:
		c.typ = String
	case []*Variant:
		c.typ = Array
		v1 := c.value.([]*Variant)
		v2 := make([]*Variant, len(v1))
		copy(v2, v1)
		c.value = v2
	case *Variant:
		v := c.value.(*Variant)
		c.typ = v.typ
		c.value = v.value
	default:
		c.typ = Object
	}
}

// Gets variant value as variant array
func (c *Variant) AsArray() []*Variant {
	return c.value.([]*Variant)
}

// Sets variant value as variant array
//
// Parameters:
//   - value a value to be set
func (c *Variant) SetAsArray(value []*Variant) {
	c.typ = Array
	a := make([]*Variant, len(value))
	c.value = copy(a, value)
}

// Gets length of the array
//
// Returns the length of the array or 0
func (c *Variant) Length() int {
	if c.typ == Array {
		a := c.value.([]*Variant)
		return len(a)
	}
	return 0
}

// Sets a new array length
//
// Parameters:
// 	- value a new array length
func (c *Variant) SetLength(value int) {
	if c.typ == Array {
		a := c.value.([]*Variant)
		for len(a) < value {
			a = append(a, &Variant{typ: Null, value: nil})
		}
		c.value = a
	} else {
		panic("Cannot set array length for non-array data type.")
	}
}

// Gets an array element by its index.
//
// Parameters:
//	- index an element index
// Returns a requested array element
func (c *Variant) GetByIndex(index int) *Variant {
	if c.typ == Array {
		a := c.value.([]*Variant)
		if len(a) > index {
			return a[index]
		} else {
			panic("Requested element of array is not accessible.")
		}
	} else {
		panic("Cannot access array element for none-array data type.")
	}
}

// Sets an array element by its index.
//
// Parameters:
// 	- index an element index
// 	- element an element value
func (c *Variant) SetByIndex(index int, element *Variant) {
	if c.typ == Array {
		a := c.value.([]*Variant)
		for len(a) <= index {
			a = append(a, &Variant{typ: Null, value: nil})
		}
		a[index] = element
		c.value = a
	} else {
		panic("Cannot access array element for none-array data type.")
	}
}

// Checks is this variant value Null.
//
// Returns <code>true</code> if this variant value is Null.
func (c *Variant) IsNull() bool {
	return c.typ == Null
}

// Checks is this variant value empty.
// Returns <code>true</code< is this variant value is empty.
func (c *Variant) IsEmpty() bool {
	return c.value == nil
}

// Assignes a new value to this object.
// Parameters:
//   - value A new value to be assigned.
func (c *Variant) Assign(value *Variant) {
	if value != nil {
		c.typ = value.typ
		c.value = value.value
	} else {
		c.typ = Null
		c.value = nil
	}
}

// Clears this object and assignes a VariantType.Null type.
func (c *Variant) Clear() {
	c.typ = Null
	c.value = nil
}

// Gets a string value for this object.
//
// Returns a string value for this object.
func (c *Variant) String() string {
	if c.value == nil {
		return "null"
	}
	return cconv.StringConverter.ToString(c.value)
}

// Compares this object to the specified one.
//
// Parameters:
//	- obj An object to be compared.
// Returns <code>true</code> if objects are equal.
func (c *Variant) Equals(obj interface{}) bool {
	if varObj, ok := obj.(Variant); ok {
		value1 := c.value
		value2 := varObj.value
		if value1 == nil || value2 == nil {
			return value1 == value2
		}
		return c.typ == varObj.typ && value1 == value2
	}
	return false
}

// Cloning the variant value
//
// Returns The cloned value of this variant
func (c *Variant) Clone() *Variant {
	return NewVariant(c)
}
