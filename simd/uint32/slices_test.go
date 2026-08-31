package simduint32_test

import (
	"slices"
	"testing"

	simduint32 "github.com/daanv2/go-math/simd/uint32"
	"github.com/stretchr/testify/assert"
)

// lengths exercises the SIMD tail handling: values below, at, and above the
// vector widths (8 for AVX256, 16 for AVX512) as well as odd remainders.
var lengths = []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 33}

// makeData returns a slice of n increasing values.
func makeData(n int) []uint32 {
	data := make([]uint32, n)
	for i := range data {
		data[i] = uint32(i)*3 + 1
	}

	return data
}

func Test_Slices_Ops(t *testing.T) {
	dataA := []uint32{0, 1, 2, 3, 4, 5, 6, 7}

	t.Run("AddToAll", func(t *testing.T) {
		s := simduint32.NewSlice(slices.Clone(dataA))

		s.AddToAll(10)

		assert.Equal(t, []uint32{10, 11, 12, 13, 14, 15, 16, 17}, s.Output())
	})

	t.Run("MulToAll", func(t *testing.T) {
		s := simduint32.NewSlice(slices.Clone(dataA))

		s.MulToAll(10)

		assert.Equal(t, []uint32{0, 10, 20, 30, 40, 50, 60, 70}, s.Output())
	})
}

// Test_Slices_MutateOps checks every in-place mutation against a plain-Go
// reference across a range of lengths, so the vectorised body and its scalar
// tail are both exercised.
func Test_Slices_MutateOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simduint32.Slice)
		ref  func(v uint32) uint32
	}{
		{"AddToAll", func(s *simduint32.Slice) { s.AddToAll(3) }, func(v uint32) uint32 { return v + 3 }},
		{"SubToAll", func(s *simduint32.Slice) { s.SubToAll(3) }, func(v uint32) uint32 { return v - 3 }},
		{"MulToAll", func(s *simduint32.Slice) { s.MulToAll(3) }, func(v uint32) uint32 { return v * 3 }},
		{"DivToAll", func(s *simduint32.Slice) { s.DivToAll(2) }, func(v uint32) uint32 { return v / 2 }},
		{"MinToAll", func(s *simduint32.Slice) { s.MinToAll(1) }, func(v uint32) uint32 { return min(v, 1) }},
		{"MaxToAll", func(s *simduint32.Slice) { s.MaxToAll(1) }, func(v uint32) uint32 { return max(v, 1) }},
		{"Fill", func(s *simduint32.Slice) { s.Fill(9) }, func(v uint32) uint32 { return 9 }},
		{"AndToAll", func(s *simduint32.Slice) { s.AndToAll(6) }, func(v uint32) uint32 { return v & 6 }},
		{"OrToAll", func(s *simduint32.Slice) { s.OrToAll(6) }, func(v uint32) uint32 { return v | 6 }},
		{"XorToAll", func(s *simduint32.Slice) { s.XorToAll(6) }, func(v uint32) uint32 { return v ^ 6 }},
		{"AndNotToAll", func(s *simduint32.Slice) { s.AndNotToAll(6) }, func(v uint32) uint32 { return v &^ 6 }},
		{"Not", func(s *simduint32.Slice) { s.Not() }, func(v uint32) uint32 { return ^v }},
		{"ShiftLeftToAll", func(s *simduint32.Slice) { s.ShiftLeftToAll(2) }, func(v uint32) uint32 { return v << 2 }},
		{"ShiftRightToAll", func(s *simduint32.Slice) { s.ShiftRightToAll(2) }, func(v uint32) uint32 { return v >> 2 }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				data := makeData(n)
				want := make([]uint32, n)
				for i, v := range data {
					want[i] = c.ref(v)
				}

				s := simduint32.NewSlice(slices.Clone(data))
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
		mut  func(s *simduint32.Slice, v []uint32)
		ref  func(a, b uint32) uint32
	}{
		{"Add", func(s *simduint32.Slice, v []uint32) { s.Add(v) }, func(a, b uint32) uint32 { return a + b }},
		{"Sub", func(s *simduint32.Slice, v []uint32) { s.Sub(v) }, func(a, b uint32) uint32 { return a - b }},
		{"Mul", func(s *simduint32.Slice, v []uint32) { s.Mul(v) }, func(a, b uint32) uint32 { return a * b }},
		{"Div", func(s *simduint32.Slice, v []uint32) { s.Div(v) }, func(a, b uint32) uint32 { return a / b }},
		{"MinWith", func(s *simduint32.Slice, v []uint32) { s.MinWith(v) }, func(a, b uint32) uint32 { return min(a, b) }},
		{"MaxWith", func(s *simduint32.Slice, v []uint32) { s.MaxWith(v) }, func(a, b uint32) uint32 { return max(a, b) }},
		{"And", func(s *simduint32.Slice, v []uint32) { s.And(v) }, func(a, b uint32) uint32 { return a & b }},
		{"Or", func(s *simduint32.Slice, v []uint32) { s.Or(v) }, func(a, b uint32) uint32 { return a | b }},
		{"Xor", func(s *simduint32.Slice, v []uint32) { s.Xor(v) }, func(a, b uint32) uint32 { return a ^ b }},
		{"AndNot", func(s *simduint32.Slice, v []uint32) { s.AndNot(v) }, func(a, b uint32) uint32 { return a &^ b }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				for _, m := range lengths {
					data := makeData(n)
					other := make([]uint32, m)
					for i := range other {
						other[i] = uint32(i) + 1
					}

					want := slices.Clone(data)
					for i := range min(n, m) {
						want[i] = c.ref(want[i], other[i])
					}

					s := simduint32.NewSlice(slices.Clone(data))
					c.mut(s, other)

					assert.Equalf(t, want, s.Output(), "len=%d, otherLen=%d", n, m)
				}
			}
		})
	}
}

// Test_Slices_Reductions checks Sum/Min/Max against plain-Go references.
func Test_Slices_Reductions(t *testing.T) {
	for _, n := range lengths {
		data := makeData(n)

		var wantSum uint32
		for _, v := range data {
			wantSum += v
		}

		wantMin, wantMax := uint32(0), uint32(0)
		if n > 0 {
			wantMin, wantMax = data[0], data[0]
			for _, v := range data[1:] {
				wantMin = min(wantMin, v)
				wantMax = max(wantMax, v)
			}
		}

		s := simduint32.NewSlice(slices.Clone(data))
		assert.Equalf(t, wantSum, s.Sum(), "Sum len=%d", n)
		assert.Equalf(t, wantMin, s.Min(), "Min len=%d", n)
		assert.Equalf(t, wantMax, s.Max(), "Max len=%d", n)
	}
}

// Test_Slices_NilAndEmpty makes sure the guard clauses hold.
func Test_Slices_NilAndEmpty(t *testing.T) {
	var nilSlice *simduint32.Slice

	assert.NotPanics(t, func() {
		nilSlice.AddToAll(1)
		nilSlice.SubToAll(1)
		nilSlice.MulToAll(1)
		nilSlice.DivToAll(1)
		nilSlice.MinToAll(1)
		nilSlice.MaxToAll(1)
		nilSlice.Fill(1)
		nilSlice.AndToAll(1)
		nilSlice.OrToAll(1)
		nilSlice.XorToAll(1)
		nilSlice.AndNotToAll(1)
		nilSlice.Not()
		nilSlice.ShiftLeftToAll(1)
		nilSlice.ShiftRightToAll(1)
		nilSlice.Add([]uint32{1})
		nilSlice.Sub([]uint32{1})
		nilSlice.Mul([]uint32{1})
		nilSlice.Div([]uint32{1})
		nilSlice.MinWith([]uint32{1})
		nilSlice.MaxWith([]uint32{1})
		nilSlice.And([]uint32{1})
		nilSlice.Or([]uint32{1})
		nilSlice.Xor([]uint32{1})
		nilSlice.AndNot([]uint32{1})
	})

	assert.Equal(t, uint32(0), nilSlice.Sum())
	assert.Equal(t, uint32(0), nilSlice.Min())
	assert.Equal(t, uint32(0), nilSlice.Max())

	empty := simduint32.NewSlice([]uint32{})
	assert.Equal(t, uint32(0), empty.Sum())
	assert.Equal(t, uint32(0), empty.Min())
	assert.Equal(t, uint32(0), empty.Max())
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

		var data []uint32
		for i := range a {
			data = append(data, uint32(i))
		}

		s := simduint32.NewSlice(data)
		s.AddToAll(10)

	})
}
