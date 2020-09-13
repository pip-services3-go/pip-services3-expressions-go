package functions

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/pip-services3-go/pip-services3-expressions-go/calculator/errors"
	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
)

// Implements a list filled with standard functions.
type DefaultFunctionCollection struct {
	FunctionCollection
}

// Constructs this list and fills it with the standard functions.
func NewDefaultFunctionCollection() *DefaultFunctionCollection {
	c := &DefaultFunctionCollection{
		FunctionCollection: *NewFunctionCollection(),
	}

	c.Add(NewDelegatedFunction("Time", timeFunctionCalculator))
	c.Add(NewDelegatedFunction("Now", timeFunctionCalculator))
	c.Add(NewDelegatedFunction("Min", minFunctionCalculator))
	c.Add(NewDelegatedFunction("Max", maxFunctionCalculator))
	c.Add(NewDelegatedFunction("Sum", sumFunctionCalculator))
	c.Add(NewDelegatedFunction("If", ifFunctionCalculator))
	c.Add(NewDelegatedFunction("Choose", chooseFunctionCalculator))
	c.Add(NewDelegatedFunction("E", eFunctionCalculator))
	c.Add(NewDelegatedFunction("Pi", piFunctionCalculator))
	c.Add(NewDelegatedFunction("Rnd", rndFunctionCalculator))
	c.Add(NewDelegatedFunction("Random", rndFunctionCalculator))
	c.Add(NewDelegatedFunction("Abs", absFunctionCalculator))
	c.Add(NewDelegatedFunction("Acos", acosFunctionCalculator))
	c.Add(NewDelegatedFunction("Asin", asinFunctionCalculator))
	c.Add(NewDelegatedFunction("Atan", atanFunctionCalculator))
	c.Add(NewDelegatedFunction("Exp", expFunctionCalculator))
	c.Add(NewDelegatedFunction("Log", logFunctionCalculator))
	c.Add(NewDelegatedFunction("Ln", logFunctionCalculator))
	c.Add(NewDelegatedFunction("Log10", log10FunctionCalculator))
	c.Add(NewDelegatedFunction("Ceil", ceilFunctionCalculator))
	c.Add(NewDelegatedFunction("Ceiling", ceilFunctionCalculator))
	c.Add(NewDelegatedFunction("Floor", floorFunctionCalculator))
	c.Add(NewDelegatedFunction("Round", roundFunctionCalculator))
	c.Add(NewDelegatedFunction("Trunc", truncFunctionCalculator))
	c.Add(NewDelegatedFunction("Truncate", truncFunctionCalculator))
	c.Add(NewDelegatedFunction("Cos", cosFunctionCalculator))
	c.Add(NewDelegatedFunction("Sin", sinFunctionCalculator))
	c.Add(NewDelegatedFunction("Tan", tanFunctionCalculator))
	c.Add(NewDelegatedFunction("Sqr", sqrtFunctionCalculator))
	c.Add(NewDelegatedFunction("Sqrt", sqrtFunctionCalculator))
	c.Add(NewDelegatedFunction("Empty", emptyFunctionCalculator))
	c.Add(NewDelegatedFunction("Null", nullFunctionCalculator))
	c.Add(NewDelegatedFunction("Contains", containsFunctionCalculator))

	return c
}

// Checks if parameters contains the correct number of function parameters (must be stored on the top of the parameters).
//
// Parameters:
//   - parameters: A list with function parameters.
//   - expectedParamCount: The expected number of function parameters.
func checkParamCount(parameters []*variants.Variant, expectedParamCount int) error {
	paramCount := len(parameters)
	if expectedParamCount != paramCount {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected "+string(expectedParamCount)+
				" parameters but was found "+string(paramCount))
		return err
	}
	return nil
}

// Gets function parameter by it's index.
//
// Parameters:
//   - parameters: A list with function parameters.
//   - paramIndex: Index for the function parameter (0 for the first parameter).
// Returns: Function parameter value.
func getParameter(parameters []*variants.Variant, paramIndex int) *variants.Variant {
	return parameters[paramIndex]
}

func timeFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 0)
	if err != nil {
		return nil, err
	}

	result := variants.VariantFromLong(time.Now().Unix())

	return result, nil
}

func minFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	paramCount := len(parameters)
	if paramCount < 2 {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected at least 2 parameters")
		return nil, err
	}

	result := getParameter(parameters, 0)
	for i := 1; i < paramCount; i = i + 1 {
		value := getParameter(parameters, i)
		if variantOperations.More(result, value).AsBoolean() {
			result = value
		}
	}

	return result, nil
}

func maxFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	paramCount := len(parameters)
	if paramCount < 2 {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected at least 2 parameters")
		return nil, err
	}

	result := getParameter(parameters, 0)
	for i := 1; i < paramCount; i++ {
		value := getParameter(parameters, i)
		if variantOperations.Less(result, value).AsBoolean() {
			result = value
		}
	}

	return result, nil
}

func sumFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	paramCount := len(parameters)
	if paramCount < 2 {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected at least 2 parameters")
		return nil, err
	}

	result := getParameter(parameters, 0)
	for i := 1; i < paramCount; i++ {
		value := getParameter(parameters, i)
		result = variantOperations.Add(result, value)
	}

	return result, nil
}

func ifFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 3)
	if err != nil {
		return nil, err
	}

	value1 := getParameter(parameters, 0)
	value2 := getParameter(parameters, 1)
	value3 := getParameter(parameters, 2)
	condition := variantOperations.Convert(value1, variants.Boolean)
	result := value3
	if condition.AsBoolean() {
		result = value2
	}

	return result, nil
}

func chooseFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	paramCount := len(parameters)
	if paramCount < 3 {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected at least 3 parameters")
		return nil, err
	}

	value1 := getParameter(parameters, 0)
	condition := variantOperations.Convert(value1, variants.Integer)
	paramIndex := int(condition.AsInteger())

	if paramCount < paramIndex+1 {
		err := errors.NewExpressionError("", "WRONG_PARAM_COUNT",
			"Expected at least "+string(paramIndex+1)+" parameters")
		return nil, err
	}

	result := getParameter(parameters, paramIndex)

	return result, nil
}

func eFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 0)
	if err != nil {
		return nil, err
	}

	result := variants.VariantFromFloat(math.E)

	return result, nil
}

func piFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 0)
	if err != nil {
		return nil, err
	}

	result := variants.VariantFromFloat(math.Pi)

	return result, nil
}

func rndFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 0)
	if err != nil {
		return nil, err
	}

	result := variants.VariantFromFloat(rand.Float32())

	return result, nil
}

func absFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := getParameter(parameters, 0)
	result := variants.EmptyVariant()
	switch value.Type() {
	case variants.Integer:
		result.SetAsInteger(int(math.Abs(float64(value.AsInteger()))))
		break
	case variants.Long:
		result.SetAsLong(int64(math.Abs(float64(value.AsLong()))))
		break
	case variants.Float:
		result.SetAsFloat(float32(math.Abs(float64(value.AsFloat()))))
		break
	case variants.Double:
		result.SetAsDouble(math.Abs(value.AsDouble()))
		break
	default:
		value = variantOperations.Convert(value, variants.Double)
		result.SetAsDouble(math.Abs(value.AsDouble()))
		break
	}

	return result, nil
}

func acosFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Acos(value.AsDouble()))

	return result, nil
}

func asinFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Asin(value.AsDouble()))

	return result, nil
}

func atanFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Atan(value.AsDouble()))

	return result, nil
}

func expFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Exp(value.AsDouble()))

	return result, nil
}

func logFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Log(value.AsDouble()))

	return result, nil
}

func log10FunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Log10(value.AsDouble()))

	return result, nil
}

func ceilFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Ceil(value.AsDouble()))

	return result, nil
}

func floorFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Floor(value.AsDouble()))

	return result, nil
}

func roundFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Round(value.AsDouble()))

	return result, nil
}

func truncFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromLong(int64(math.Trunc(value.AsDouble())))

	return result, nil
}

func cosFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Cos(value.AsDouble()))

	return result, nil
}

func sinFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Sin(value.AsDouble()))

	return result, nil
}

func tanFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Tan(value.AsDouble()))

	return result, nil
}

func sqrtFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := variantOperations.Convert(getParameter(parameters, 0), variants.Double)
	result := variants.VariantFromDouble(math.Sqrt(value.AsDouble()))

	return result, nil
}

func emptyFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 1)
	if err != nil {
		return nil, err
	}

	value := getParameter(parameters, 0)
	result := variants.VariantFromBoolean(value.IsEmpty())

	return result, nil
}

func nullFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 0)
	if err != nil {
		return nil, err
	}

	result := variants.EmptyVariant()

	return result, nil
}

func containsFunctionCalculator(parameters []*variants.Variant,
	variantOperations variants.IVariantOperations) (*variants.Variant, error) {
	err := checkParamCount(parameters, 2)
	if err != nil {
		return nil, err
	}

	str := variantOperations.Convert(getParameter(parameters, 0), variants.String)
	substr := variantOperations.Convert(getParameter(parameters, 1), variants.String)

	if str.IsEmpty() || str.IsNull() {
		return variants.VariantFromBoolean(false), nil
	}

	contains := strings.Contains(str.AsString(), substr.AsString())
	result := variants.VariantFromBoolean(contains)

	return result, nil
}
