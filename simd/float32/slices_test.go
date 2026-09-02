package simdfloat32_test

import (
	"math"
	"slices"
	"testing"

	simdfloat32 "github.com/daanv2/go-math/simd/float32"
	"github.com/stretchr/testify/assert"
)

// lengths exercises the SIMD tail handling: values below, at, and above the
// vector widths (8 for AVX256, 16 for AVX512) as well as odd remainders.
var lengths = []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 33}

// makeData returns a slice of n values with a mix of signs and magnitudes.
func makeData(n int) []float32 {
	data := make([]float32, n)
	for i := range data {
		data[i] = float32(i)*1.5 - float32(n)/2
	}

	return data
}

func Test_Slices_Ops(t *testing.T) {
	dataA := []float32{0, 1, 2, 3, 4, 5, 6, 7}

	t.Run("AddToAll", func(t *testing.T) {
		s := simdfloat32.NewSlice(slices.Clone(dataA))

		s.AddToAll(10)

		assert.Equal(t, []float32{10, 11, 12, 13, 14, 15, 16, 17}, s.Output())
	})

	t.Run("MulToAll", func(t *testing.T) {
		s := simdfloat32.NewSlice(slices.Clone(dataA))

		s.MulToAll(10)

		assert.Equal(t, []float32{0, 10, 20, 30, 40, 50, 60, 70}, s.Output())
	})
}

// Test_Slices_MutateOps checks every in-place mutation against a plain-Go
// reference across a range of lengths, so the vectorised body and its scalar
// tail are both exercised.
func Test_Slices_MutateOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simdfloat32.Slice)
		ref  func(v float32) float32
	}{
		{"AddToAll", func(s *simdfloat32.Slice) { s.AddToAll(3) }, func(v float32) float32 { return v + 3 }},
		{"SubToAll", func(s *simdfloat32.Slice) { s.SubToAll(3) }, func(v float32) float32 { return v - 3 }},
		{"MulToAll", func(s *simdfloat32.Slice) { s.MulToAll(3) }, func(v float32) float32 { return v * 3 }},
		{"DivToAll", func(s *simdfloat32.Slice) { s.DivToAll(2) }, func(v float32) float32 { return v / 2 }},
		{"MinToAll", func(s *simdfloat32.Slice) { s.MinToAll(1) }, func(v float32) float32 { return min(v, 1) }},
		{"MaxToAll", func(s *simdfloat32.Slice) { s.MaxToAll(1) }, func(v float32) float32 { return max(v, 1) }},
		{"Fill", func(s *simdfloat32.Slice) { s.Fill(9) }, func(v float32) float32 { return 9 }},
		{"Negate", func(s *simdfloat32.Slice) { s.Negate() }, func(v float32) float32 { return -v }},
		{"Abs", func(s *simdfloat32.Slice) { s.Abs() }, func(v float32) float32 { return float32(math.Abs(float64(v))) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				data := makeData(n)
				want := make([]float32, n)
				for i, v := range data {
					want[i] = c.ref(v)
				}

				s := simdfloat32.NewSlice(slices.Clone(data))
				c.mut(s)

				assert.Equalf(t, want, s.Output(), "len=%d", n)
			}
		})
	}
}

// Test_Slices_PairwiseOps checks the slice-vs-slice mutation methods against
// plain-Go references, including cases where the two slices differ in length
// so only min(len(s), len(v)) elements should be touched.
func Test_Slices_PairwiseOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simdfloat32.Slice, v []float32)
		ref  func(a, b float32) float32
	}{
		{"Add", func(s *simdfloat32.Slice, v []float32) { s.Add(v) }, func(a, b float32) float32 { return a + b }},
		{"Sub", func(s *simdfloat32.Slice, v []float32) { s.Sub(v) }, func(a, b float32) float32 { return a - b }},
		{"Mul", func(s *simdfloat32.Slice, v []float32) { s.Mul(v) }, func(a, b float32) float32 { return a * b }},
		{"Div", func(s *simdfloat32.Slice, v []float32) { s.Div(v) }, func(a, b float32) float32 { return a / b }},
		{"MinWith", func(s *simdfloat32.Slice, v []float32) { s.MinWith(v) }, func(a, b float32) float32 { return min(a, b) }},
		{"MaxWith", func(s *simdfloat32.Slice, v []float32) { s.MaxWith(v) }, func(a, b float32) float32 { return max(a, b) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				for _, m := range lengths {
					data := makeData(n)
					other := make([]float32, m)
					for i := range other {
						other[i] = float32(i) + 1
					}

					want := slices.Clone(data)
					for i := range min(n, m) {
						want[i] = c.ref(want[i], other[i])
					}

					s := simdfloat32.NewSlice(slices.Clone(data))
					c.mut(s, other)

					assert.Equalf(t, want, s.Output(), "len=%d, otherLen=%d", n, m)
				}
			}
		})
	}
}

// Test_Slices_Sqrt is separated out because it needs non-negative inputs.
func Test_Slices_Sqrt(t *testing.T) {
	for _, n := range lengths {
		data := make([]float32, n)
		want := make([]float32, n)
		for i := range data {
			data[i] = float32(i) * 2
			want[i] = float32(math.Sqrt(float64(data[i])))
		}

		s := simdfloat32.NewSlice(slices.Clone(data))
		s.Sqrt()

		assert.Equalf(t, want, s.Output(), "len=%d", n)
	}
}

// Test_Slices_Reductions checks Sum/Min/Max against plain-Go references.
func Test_Slices_Reductions(t *testing.T) {
	for _, n := range lengths {
		data := makeData(n)

		var wantSum float32
		for _, v := range data {
			wantSum += v
		}

		wantMin, wantMax := float32(0), float32(0)
		if n > 0 {
			wantMin, wantMax = data[0], data[0]
			for _, v := range data[1:] {
				wantMin = min(wantMin, v)
				wantMax = max(wantMax, v)
			}
		}

		s := simdfloat32.NewSlice(slices.Clone(data))
		assert.InDeltaf(t, wantSum, s.Sum(), 1e-3, "Sum len=%d", n)
		assert.Equalf(t, wantMin, s.Min(), "Min len=%d", n)
		assert.Equalf(t, wantMax, s.Max(), "Max len=%d", n)
	}
}

// Test_Slices_NilAndEmpty makes sure the guard clauses hold.
func Test_Slices_NilAndEmpty(t *testing.T) {
	var nilSlice *simdfloat32.Slice

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
		nilSlice.Sqrt()
		nilSlice.Add([]float32{1})
		nilSlice.Sub([]float32{1})
		nilSlice.Mul([]float32{1})
		nilSlice.Div([]float32{1})
		nilSlice.MinWith([]float32{1})
		nilSlice.MaxWith([]float32{1})
	})

	assert.Equal(t, float32(0), nilSlice.Sum())
	assert.Equal(t, float32(0), nilSlice.Min())
	assert.Equal(t, float32(0), nilSlice.Max())

	empty := simdfloat32.NewSlice([]float32{})
	assert.Equal(t, float32(0), empty.Sum())
	assert.Equal(t, float32(0), empty.Min())
	assert.Equal(t, float32(0), empty.Max())
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

		var data []float32
		for i := range a {
			data = append(data, float32(i))
		}

		s := simdfloat32.NewSlice(data)
		s.AddToAll(10)

	})
}
