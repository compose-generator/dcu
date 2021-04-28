package dcu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------ AppendStringToSliceIfMissing ------------------------------------------

func TestAppendStringToSliceIfMissing_Missing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = appendStringToSliceIfMissing(slice, "ipsum")
	assert.Equal(t, 5, len(slice))
}

func TestAppendStringToSliceIfMissing_NotMissing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = appendStringToSliceIfMissing(slice, "bar")
	assert.Equal(t, 4, len(slice))
}

// ------------------------------------------ TestSliceContainsString ------------------------------------------

func TestSliceContainsString_True(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := sliceContainsString(slice, "bar")
	assert.True(t, result)
}

func TestSliceContainsString_False(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := sliceContainsString(slice, "ipsum")
	assert.False(t, result)
}
