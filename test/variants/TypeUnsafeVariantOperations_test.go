package test_variants

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/variants"
	"github.com/stretchr/testify/assert"
)

func TestUnsafeOperations(t *testing.T) {
	a := variants.NewVariant(123)
	manager := variants.NewTypeUnsafeVariantOperations()

	b := manager.Convert(a, variants.Float)
	assert.Equal(t, variants.Float, b.Type())
	assert.Equal(t, float32(123.0), b.AsFloat())

	c := variants.NewVariant(2)
	assert.Equal(t, float32(125.0), manager.Add(b, c).AsFloat())
	assert.Equal(t, float32(121.0), manager.Sub(b, c).AsFloat())
	assert.True(t, manager.Equal(a, b).AsBoolean())
}
