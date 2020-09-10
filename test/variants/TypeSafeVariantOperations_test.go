package test_variants

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
	"github.com/stretchr/testify/assert"
)

func TestSafeOperations(t *testing.T) {
	a := variants.NewVariant(int32(123))
	manager := variants.NewTypeSafeVariantOperations()

	b := manager.Convert(a, variants.Float)
	assert.Equal(t, variants.Float, b.Type())
	assert.Equal(t, float32(123.0), b.AsFloat())

	c := variants.NewVariant(int32(2))
	assert.Equal(t, int32(125), manager.Add(a, c).AsInteger())
	assert.Equal(t, int32(121), manager.Sub(a, c).AsInteger())
	assert.False(t, manager.Equal(a, c).AsBoolean())

	array := []*variants.Variant{
		variants.NewVariant("aaa"),
		variants.NewVariant("bbb"),
		variants.NewVariant("ccc"),
		variants.NewVariant("ddd"),
	}
	d := variants.NewVariant(array)
	assert.True(t, manager.In(d, variants.NewVariant("ccc")).AsBoolean())
	assert.False(t, manager.In(d, variants.NewVariant("eee")).AsBoolean())
	assert.Equal(t, "bbb", manager.GetElement(d, variants.NewVariant(1)).AsString())
}
