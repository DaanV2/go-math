package simdbytes_test

import (
	"slices"
	"testing"

	simdbytes "github.com/daanv2/go-math/simd/bytes"
	"github.com/stretchr/testify/assert"
)

// lengths exercises the SIMD tail handling: values below, at, and above the
// vector widths (32 for AVX256, 64 for AVX512) as well as odd remainders.
var lengths = []int{0, 1, 2, 3, 4, 5, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65}

// makeData returns a slice of n increasing values.
func makeData(n int) []byte {
	data := make([]byte, n)
	for i := range data {
		data[i] = byte(i)*3 + 1
	}

	return data
}

func Test_Slices_Ops(t *testing.T) {
	dataA := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	t.Run("AddToAll", func(t *testing.T) {
		s := simdbytes.NewSlice(slices.Clone(dataA))

		s.AddToAll(10)

		assert.Equal(t, []byte{10, 11, 12, 13, 14, 15, 16, 17}, s.Output())
	})

	t.Run("MulToAll", func(t *testing.T) {
		s := simdbytes.NewSlice(slices.Clone(dataA))

		s.MulToAll(10)

		assert.Equal(t, []byte{0, 10, 20, 30, 40, 50, 60, 70}, s.Output())
	})
}

// Test_Slices_MutateOps checks every in-place mutation against a plain-Go
// reference across a range of lengths, so the vectorised body and its scalar
// tail are both exercised.
func Test_Slices_MutateOps(t *testing.T) {
	cases := []struct {
		name string
		mut  func(s *simdbytes.Slice)
		ref  func(v byte) byte
	}{
		{"AddToAll", func(s *simdbytes.Slice) { s.AddToAll(3) }, func(v byte) byte { return v + 3 }},
		{"SubToAll", func(s *simdbytes.Slice) { s.SubToAll(3) }, func(v byte) byte { return v - 3 }},
		{"MulToAll", func(s *simdbytes.Slice) { s.MulToAll(3) }, func(v byte) byte { return v * 3 }},
		{"DivToAll", func(s *simdbytes.Slice) { s.DivToAll(2) }, func(v byte) byte { return v / 2 }},
		{"MinToAll", func(s *simdbytes.Slice) { s.MinToAll(1) }, func(v byte) byte { return min(v, 1) }},
		{"MaxToAll", func(s *simdbytes.Slice) { s.MaxToAll(1) }, func(v byte) byte { return max(v, 1) }},
		{"Fill", func(s *simdbytes.Slice) { s.Fill(9) }, func(v byte) byte { return 9 }},
		{"AndToAll", func(s *simdbytes.Slice) { s.AndToAll(6) }, func(v byte) byte { return v & 6 }},
		{"OrToAll", func(s *simdbytes.Slice) { s.OrToAll(6) }, func(v byte) byte { return v | 6 }},
		{"XorToAll", func(s *simdbytes.Slice) { s.XorToAll(6) }, func(v byte) byte { return v ^ 6 }},
		{"AndNotToAll", func(s *simdbytes.Slice) { s.AndNotToAll(6) }, func(v byte) byte { return v &^ 6 }},
		{"Not", func(s *simdbytes.Slice) { s.Not() }, func(v byte) byte { return ^v }},
		{"ShiftLeftToAll", func(s *simdbytes.Slice) { s.ShiftLeftToAll(2) }, func(v byte) byte { return v << 2 }},
		{"ShiftRightToAll", func(s *simdbytes.Slice) { s.ShiftRightToAll(2) }, func(v byte) byte { return v >> 2 }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				data := makeData(n)
				want := make([]byte, n)
				for i, v := range data {
					want[i] = c.ref(v)
				}

				s := simdbytes.NewSlice(slices.Clone(data))
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
		mut  func(s *simdbytes.Slice, v []byte)
		ref  func(a, b byte) byte
	}{
		{"Add", func(s *simdbytes.Slice, v []byte) { s.Add(v) }, func(a, b byte) byte { return a + b }},
		{"Sub", func(s *simdbytes.Slice, v []byte) { s.Sub(v) }, func(a, b byte) byte { return a - b }},
		{"Mul", func(s *simdbytes.Slice, v []byte) { s.Mul(v) }, func(a, b byte) byte { return a * b }},
		{"Div", func(s *simdbytes.Slice, v []byte) { s.Div(v) }, func(a, b byte) byte { return a / b }},
		{"MinWith", func(s *simdbytes.Slice, v []byte) { s.MinWith(v) }, func(a, b byte) byte { return min(a, b) }},
		{"MaxWith", func(s *simdbytes.Slice, v []byte) { s.MaxWith(v) }, func(a, b byte) byte { return max(a, b) }},
		{"And", func(s *simdbytes.Slice, v []byte) { s.And(v) }, func(a, b byte) byte { return a & b }},
		{"Or", func(s *simdbytes.Slice, v []byte) { s.Or(v) }, func(a, b byte) byte { return a | b }},
		{"Xor", func(s *simdbytes.Slice, v []byte) { s.Xor(v) }, func(a, b byte) byte { return a ^ b }},
		{"AndNot", func(s *simdbytes.Slice, v []byte) { s.AndNot(v) }, func(a, b byte) byte { return a &^ b }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, n := range lengths {
				for _, m := range lengths {
					data := makeData(n)
					other := make([]byte, m)
					for i := range other {
						other[i] = byte(i) + 1
					}

					want := slices.Clone(data)
					for i := range min(n, m) {
						want[i] = c.ref(want[i], other[i])
					}

					s := simdbytes.NewSlice(slices.Clone(data))
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

		var wantSum byte
		for _, v := range data {
			wantSum += v
		}

		wantMin, wantMax := byte(0), byte(0)
		if n > 0 {
			wantMin, wantMax = data[0], data[0]
			for _, v := range data[1:] {
				wantMin = min(wantMin, v)
				wantMax = max(wantMax, v)
			}
		}

		s := simdbytes.NewSlice(slices.Clone(data))
		assert.Equalf(t, wantSum, s.Sum(), "Sum len=%d", n)
		assert.Equalf(t, wantMin, s.Min(), "Min len=%d", n)
		assert.Equalf(t, wantMax, s.Max(), "Max len=%d", n)
	}
}

// Test_Slices_NilAndEmpty makes sure the guard clauses hold.
func Test_Slices_NilAndEmpty(t *testing.T) {
	var nilSlice *simdbytes.Slice

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
		nilSlice.Add([]byte{1})
		nilSlice.Sub([]byte{1})
		nilSlice.Mul([]byte{1})
		nilSlice.Div([]byte{1})
		nilSlice.MinWith([]byte{1})
		nilSlice.MaxWith([]byte{1})
		nilSlice.And([]byte{1})
		nilSlice.Or([]byte{1})
		nilSlice.Xor([]byte{1})
		nilSlice.AndNot([]byte{1})
	})

	assert.Equal(t, byte(0), nilSlice.Sum())
	assert.Equal(t, byte(0), nilSlice.Min())
	assert.Equal(t, byte(0), nilSlice.Max())

	empty := simdbytes.NewSlice([]byte{})
	assert.Equal(t, byte(0), empty.Sum())
	assert.Equal(t, byte(0), empty.Min())
	assert.Equal(t, byte(0), empty.Max())
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

		var data []byte
		for i := range a {
			data = append(data, byte(i))
		}

		s := simdbytes.NewSlice(data)
		s.AddToAll(10)

	})
}
