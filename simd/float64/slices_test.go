package simdfloat64_test

import (
	"math"
	"slices"
	"testing"

	simdfloat64 "github.com/daanv2/go-math/simd/float64"
	"github.com/stretchr/testify/assert"
)

// lengths exercises the SIMD tail handling: values below, at, and above the
// vector widths (4 for AVX256, 8 for AVX512) as well as odd remainders.
var lengths = []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 33}

// makeData returns a slice of n values with a mix of signs and magnitudes.
func makeData(n int) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = float64(i)*1.5 - float64(n)/2
	}

	return data
}

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

// Test_Slices_MutateOps checks every in-place mutation against a plain-Go
// reference across a range of lengths, so the vectorised body and its scalar
// tail are both exercised.
func Test_Slices_MutateOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simdfloat64.Slice)
		ref  func(v float64) float64
	}{
		{"AddToAll", func(s *simdfloat64.Slice) { s.AddToAll(3) }, func(v float64) float64 { return v + 3 }},
		{"SubToAll", func(s *simdfloat64.Slice) { s.SubToAll(3) }, func(v float64) float64 { return v - 3 }},
		{"MulToAll", func(s *simdfloat64.Slice) { s.MulToAll(3) }, func(v float64) float64 { return v * 3 }},
		{"DivToAll", func(s *simdfloat64.Slice) { s.DivToAll(2) }, func(v float64) float64 { return v / 2 }},
		{"MinToAll", func(s *simdfloat64.Slice) { s.MinToAll(1) }, func(v float64) float64 { return min(v, 1) }},
		{"MaxToAll", func(s *simdfloat64.Slice) { s.MaxToAll(1) }, func(v float64) float64 { return max(v, 1) }},
		{"Fill", func(s *simdfloat64.Slice) { s.Fill(9) }, func(v float64) float64 { return 9 }},
		{"Negate", func(s *simdfloat64.Slice) { s.Negate() }, func(v float64) float64 { return -v }},
		{"Abs", func(s *simdfloat64.Slice) { s.Abs() }, math.Abs},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				data := makeData(n)
				want := make([]float64, n)
				for i, v := range data {
					want[i] = c.ref(v)
				}

				s := simdfloat64.NewSlice(slices.Clone(data))
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
		mut  func(s *simdfloat64.Slice, v []float64)
		ref  func(a, b float64) float64
	}{
		{"Add", func(s *simdfloat64.Slice, v []float64) { s.Add(v) }, func(a, b float64) float64 { return a + b }},
		{"Sub", func(s *simdfloat64.Slice, v []float64) { s.Sub(v) }, func(a, b float64) float64 { return a - b }},
		{"Mul", func(s *simdfloat64.Slice, v []float64) { s.Mul(v) }, func(a, b float64) float64 { return a * b }},
		{"Div", func(s *simdfloat64.Slice, v []float64) { s.Div(v) }, func(a, b float64) float64 { return a / b }},
		{"MinWith", func(s *simdfloat64.Slice, v []float64) { s.MinWith(v) }, func(a, b float64) float64 { return min(a, b) }},
		{"MaxWith", func(s *simdfloat64.Slice, v []float64) { s.MaxWith(v) }, func(a, b float64) float64 { return max(a, b) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				for _, m := range lengths {
					data := makeData(n)
					other := make([]float64, m)
					for i := range other {
						other[i] = float64(i) + 1
					}

					want := slices.Clone(data)
					for i := range min(n, m) {
						want[i] = c.ref(want[i], other[i])
					}

					s := simdfloat64.NewSlice(slices.Clone(data))
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
		data := make([]float64, n)
		want := make([]float64, n)
		for i := range data {
			data[i] = float64(i) * 2
			want[i] = math.Sqrt(data[i])
		}

		s := simdfloat64.NewSlice(slices.Clone(data))
		s.Sqrt()

		assert.Equalf(t, want, s.Output(), "len=%d", n)
	}
}

// Test_Slices_Reductions checks Sum/Min/Max against plain-Go references.
func Test_Slices_Reductions(t *testing.T) {
	for _, n := range lengths {
		data := makeData(n)

		var wantSum float64
		for _, v := range data {
			wantSum += v
		}

		wantMin, wantMax := 0.0, 0.0
		if n > 0 {
			wantMin, wantMax = data[0], data[0]
			for _, v := range data[1:] {
				wantMin = min(wantMin, v)
				wantMax = max(wantMax, v)
			}
		}

		s := simdfloat64.NewSlice(slices.Clone(data))
		assert.InDeltaf(t, wantSum, s.Sum(), 1e-9, "Sum len=%d", n)
		assert.Equalf(t, wantMin, s.Min(), "Min len=%d", n)
		assert.Equalf(t, wantMax, s.Max(), "Max len=%d", n)
	}
}

// Test_Slices_NilAndEmpty makes sure the guard clauses hold.
func Test_Slices_NilAndEmpty(t *testing.T) {
	var nilSlice *simdfloat64.Slice

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
		nilSlice.Add([]float64{1})
		nilSlice.Sub([]float64{1})
		nilSlice.Mul([]float64{1})
		nilSlice.Div([]float64{1})
		nilSlice.MinWith([]float64{1})
		nilSlice.MaxWith([]float64{1})
	})

	assert.Equal(t, 0.0, nilSlice.Sum())
	assert.Equal(t, 0.0, nilSlice.Min())
	assert.Equal(t, 0.0, nilSlice.Max())

	empty := simdfloat64.NewSlice([]float64{})
	assert.Equal(t, 0.0, empty.Sum())
	assert.Equal(t, 0.0, empty.Min())
	assert.Equal(t, 0.0, empty.Max())
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
