package dcu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AppendStringToSliceIfMissing_Missing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = appendStringToSliceIfMissing(slice, "ipsum")
	assert.Equal(t, 5, len(slice))
}

func AppendStringToSliceIfMissing_NotMissing(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	slice = appendStringToSliceIfMissing(slice, "bar")
	assert.Equal(t, 4, len(slice))
}

func TestSliceContainsString_Successful(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := sliceContainsString(slice, "bar")
	assert.True(t, result)
}

func TestSliceContainsString_Failure(t *testing.T) {
	slice := []string{"test", "foo", "bar", "lorem"}
	result := sliceContainsString(slice, "ipsum")
	assert.False(t, result)
}
