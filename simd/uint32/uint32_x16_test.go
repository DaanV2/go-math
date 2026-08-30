package simduint32_test

import (
	"fmt"
	"testing"

	simduint32 "github.com/daanv2/go-math/simd/uint32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Uint32x16_Ops(t *testing.T) {
	dataA := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	dataB := []uint32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	vec1 := simduint32.NewUint32x16(dataA)
	vec2 := simduint32.NewUint32x16(dataB)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []uint32{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []uint32{9, 20, 33, 48, 65, 84, 105, 128, 153, 180, 209, 240, 273, 308, 345, 384}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []uint32{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []uint32{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simduint32.NewUint32x16(dataA)
		vec2 := simduint32.NewUint32x16(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []uint32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []uint32{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []uint32{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simduint32.NewUint32x16(dataA)
		vec2 := simduint32.NewUint32x16(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec2.Sub(vec1)

		assert.Equal(t, []uint32{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}, result.ToSlice())
	})
}

// Bitwise ops
func Test_Uint32x16_Bitwise(t *testing.T) {
	vecA := simduint32.NewUint32x16([]uint32{12, 10, 255, 0, 12, 10, 255, 0, 12, 10, 255, 0, 12, 10, 255, 0})
	vecB := simduint32.NewUint32x16([]uint32{10, 6, 15, 5, 10, 6, 15, 5, 10, 6, 15, 5, 10, 6, 15, 5})

	t.Run("And", func(t *testing.T) {
		result := vecA.And(vecB)

		assert.Equal(t, []uint32{8, 2, 15, 0, 8, 2, 15, 0, 8, 2, 15, 0, 8, 2, 15, 0}, result.ToSlice())
	})

	t.Run("Or", func(t *testing.T) {
		result := vecA.Or(vecB)

		assert.Equal(t, []uint32{14, 14, 255, 5, 14, 14, 255, 5, 14, 14, 255, 5, 14, 14, 255, 5}, result.ToSlice())
	})

	t.Run("Xor", func(t *testing.T) {
		result := vecA.Xor(vecB)

		assert.Equal(t, []uint32{6, 12, 240, 5, 6, 12, 240, 5, 6, 12, 240, 5, 6, 12, 240, 5}, result.ToSlice())
	})

	t.Run("AndNot", func(t *testing.T) {
		result := vecA.AndNot(vecB)

		assert.Equal(t, []uint32{4, 8, 240, 0, 4, 8, 240, 0, 4, 8, 240, 0, 4, 8, 240, 0}, result.ToSlice())
	})

	t.Run("Not", func(t *testing.T) {
		result := vecA.Not()

		assert.Equal(t, []uint32{4294967283, 4294967285, 4294967040, 4294967295, 4294967283, 4294967285, 4294967040, 4294967295, 4294967283, 4294967285, 4294967040, 4294967295, 4294967283, 4294967285, 4294967040, 4294967295}, result.ToSlice())
	})

	t.Run("ShiftLeft", func(t *testing.T) {
		vec := simduint32.NewUint32x16([]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		result := vec.ShiftLeft(2)

		assert.Equal(t, []uint32{4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60, 64}, result.ToSlice())
	})

	t.Run("ShiftRight", func(t *testing.T) {
		vec := simduint32.NewUint32x16([]uint32{8, 255, 16, 4, 8, 255, 16, 4, 8, 255, 16, 4, 8, 255, 16, 4})
		result := vec.ShiftRight(1)

		assert.Equal(t, []uint32{4, 127, 8, 2, 4, 127, 8, 2, 4, 127, 8, 2, 4, 127, 8, 2}, result.ToSlice())
	})
}

// Loads
func Test_Uint32x16_Load(t *testing.T) {
	data := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	assert.Equal(t, simduint32.NewUint32x16(data).ToSlice(), data[:16])
	assert.Equal(t, []uint32{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, simduint32.NewUint32x16(data[:1]).ToSlice())
	assert.Equal(t, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}, simduint32.NewUint32x16(data[:15]).ToSlice())
}

func Test_Uint32x16Slice(t *testing.T) {
	data := []uint32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25}

	result := simduint32.NewUint32x16Slice(data)

	require.Len(t, result, 3)
	assert.Equal(t, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result[0].ToSlice())
	assert.Equal(t, []uint32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result[1].ToSlice())
	assert.Equal(t, []uint32{25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleUint32x16() {
	data := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	v1 := simduint32.NewUint32x16(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
}
