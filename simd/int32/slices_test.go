package simdint32_test

import (
	"slices"
	"testing"

	simdint32 "github.com/daanv2/go-math/simd/int32"
	"github.com/stretchr/testify/assert"
)

// lengths exercises the SIMD tail handling: values below, at, and above the
// vector widths (8 for AVX256, 16 for AVX512) as well as odd remainders.
var lengths = []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 33}

// makeData returns a slice of n values with a mix of signs and magnitudes.
func makeData(n int) []int32 {
	data := make([]int32, n)
	for i := range data {
		data[i] = int32(i)*3 - int32(n)/2
	}
	return data
}

func Test_Slices_Ops(t *testing.T) {
	dataA := []int32{0, 1, 2, 3, 4, 5, 6, 7}

	t.Run("AddToAll", func(t *testing.T) {
		s := simdint32.NewSlice(slices.Clone(dataA))

		s.AddToAll(10)

		assert.Equal(t, []int32{10, 11, 12, 13, 14, 15, 16, 17}, s.Output())
	})

	t.Run("MulToAll", func(t *testing.T) {
		s := simdint32.NewSlice(slices.Clone(dataA))

		s.MulToAll(10)

		assert.Equal(t, []int32{0, 10, 20, 30, 40, 50, 60, 70}, s.Output())
	})
}

// Test_Slices_MutateOps checks every in-place mutation against a plain-Go
// reference across a range of lengths, so the vectorised body and its scalar
// tail are both exercised.
func Test_Slices_MutateOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simdint32.Slice)
		ref  func(v int32) int32
	}{
		{"AddToAll", func(s *simdint32.Slice) { s.AddToAll(3) }, func(v int32) int32 { return v + 3 }},
		{"SubToAll", func(s *simdint32.Slice) { s.SubToAll(3) }, func(v int32) int32 { return v - 3 }},
		{"MulToAll", func(s *simdint32.Slice) { s.MulToAll(3) }, func(v int32) int32 { return v * 3 }},
		{"DivToAll", func(s *simdint32.Slice) { s.DivToAll(2) }, func(v int32) int32 { return v / 2 }},
		{"MinToAll", func(s *simdint32.Slice) { s.MinToAll(1) }, func(v int32) int32 { return min(v, 1) }},
		{"MaxToAll", func(s *simdint32.Slice) { s.MaxToAll(1) }, func(v int32) int32 { return max(v, 1) }},
		{"Fill", func(s *simdint32.Slice) { s.Fill(9) }, func(v int32) int32 { return 9 }},
		{"Negate", func(s *simdint32.Slice) { s.Negate() }, func(v int32) int32 { return -v }},
		{"Abs", func(s *simdint32.Slice) { s.Abs() }, func(v int32) int32 {
			if v < 0 {
				return -v
			}
			return v
		}},
		{"AndToAll", func(s *simdint32.Slice) { s.AndToAll(6) }, func(v int32) int32 { return v & 6 }},
		{"OrToAll", func(s *simdint32.Slice) { s.OrToAll(6) }, func(v int32) int32 { return v | 6 }},
		{"XorToAll", func(s *simdint32.Slice) { s.XorToAll(6) }, func(v int32) int32 { return v ^ 6 }},
		{"AndNotToAll", func(s *simdint32.Slice) { s.AndNotToAll(6) }, func(v int32) int32 { return v &^ 6 }},
		{"Not", func(s *simdint32.Slice) { s.Not() }, func(v int32) int32 { return ^v }},
		{"ShiftLeftToAll", func(s *simdint32.Slice) { s.ShiftLeftToAll(2) }, func(v int32) int32 { return v << 2 }},
		{"ShiftRightToAll", func(s *simdint32.Slice) { s.ShiftRightToAll(2) }, func(v int32) int32 { return v >> 2 }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				data := makeData(n)
				want := make([]int32, n)
				for i, v := range data {
					want[i] = c.ref(v)
				}

				s := simdint32.NewSlice(slices.Clone(data))
				c.mut(s)

				assert.Equalf(t, want, s.Output(), "len=%d", n)
			}
		})
	}
}

// Test_Slices_Reductions checks Sum/Min/Max against plain-Go references.
func Test_Slices_Reductions(t *testing.T) {
	for _, n := range lengths {
		data := makeData(n)

		var wantSum int32
		for _, v := range data {
			wantSum += v
		}

		wantMin, wantMax := int32(0), int32(0)
		if n > 0 {
			wantMin, wantMax = data[0], data[0]
			for _, v := range data[1:] {
				wantMin = min(wantMin, v)
				wantMax = max(wantMax, v)
			}
		}

		s := simdint32.NewSlice(slices.Clone(data))
		assert.Equalf(t, wantSum, s.Sum(), "Sum len=%d", n)
		assert.Equalf(t, wantMin, s.Min(), "Min len=%d", n)
		assert.Equalf(t, wantMax, s.Max(), "Max len=%d", n)
	}
}

// Test_Slices_NilAndEmpty makes sure the guard clauses hold.
func Test_Slices_NilAndEmpty(t *testing.T) {
	var nilSlice *simdint32.Slice

	assert.NotPanics(t, func() {
		nilSlice.AddToAll(1)
		nilSlice.SubToAll(1)
		nilSlice.MulToAll(1)
		nilSlice.DivToAll(1)
		nilSlice.MinToAll(1)
		nilSlice.MaxToAll(1)
		nilSlice.Fill(1)
		nilSlice.Negate()
		nilSlice.Abs()
		nilSlice.AndToAll(1)
		nilSlice.OrToAll(1)
		nilSlice.XorToAll(1)
		nilSlice.AndNotToAll(1)
		nilSlice.Not()
		nilSlice.ShiftLeftToAll(1)
		nilSlice.ShiftRightToAll(1)
	})

	assert.Equal(t, int32(0), nilSlice.Sum())
	assert.Equal(t, int32(0), nilSlice.Min())
	assert.Equal(t, int32(0), nilSlice.Max())

	empty := simdint32.NewSlice([]int32{})
	assert.Equal(t, int32(0), empty.Sum())
	assert.Equal(t, int32(0), empty.Min())
	assert.Equal(t, int32(0), empty.Max())
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

		var data []int32
		for i := range a {
			data = append(data, int32(i))
		}

		s := simdint32.NewSlice(data)
		s.AddToAll(10)

	})
}
