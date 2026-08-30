package simdfloat64_test

import (
	"slices"
	"testing"

	simdfloat64 "github.com/daanv2/go-math/simd/float64"
	"github.com/stretchr/testify/assert"
)

func Test_Slices_Ops(t *testing.T) {
	dataA := []float64{0, 1, 2, 3, 4, 5, 6, 7}

	t.Run("AddToAll", func(t *testing.T) {
		s := simdfloat64.NewSlice(slices.Clone(dataA))

		s.AddToAll(10)

		assert.Equal(t, []float64{10, 11, 12, 13, 14, 15, 16, 17}, s.Output())
	})

	t.Run("MulToAll", func(t *testing.T) {
		s := simdfloat64.NewSlice(slices.Clone(dataA))

		s.MulToAll(10)

		assert.Equal(t, []float64{0, 10, 20, 30, 40, 50, 60, 70}, s.Output())
	})
}

func Fuzz_Slices_Ops_AddToAll(f *testing.F) {
	f.Add(-1)
	f.Add(0)
	f.Add(1)
	f.Add(7)
	f.Add(8)
	f.Add(9)

	f.Fuzz(func(t *testing.T, a int) {
		if a < 0 {
			return
		}

		var data []float64
		for i := range a {
			data = append(data, float64(i))
		}

		s := simdfloat64.NewSlice(data)
		s.AddToAll(10)

	})
}
