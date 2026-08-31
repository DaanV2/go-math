//go:build simd_avx512

package simdint32

import (
	"simd/archsimd"
)

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Add(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] += v
		i += 1
	}
}

// AddToAll multiplies v all elements, updating their value with the update value
// performs: s[i] = s[i] * v
func (s *Slice) MulToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Mul(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] *= v
		i += 1
	}
}

// SubToAll subtracts v from all elements, updating their value with the result
// performs: s[i] = s[i] - v
func (s *Slice) SubToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Sub(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] -= v
		i += 1
	}
}

// DivToAll divides all elements by v, updating their value with the result
// performs: s[i] = s[i] / v
// NOTE: there is no hardware vectorised integer division, this always runs scalar
func (s *Slice) DivToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] /= v
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Min(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = min(s.data[i], v)
		i += 1
	}
}

// MaxToAll sets each element to the maximum of itself and v
// performs: s[i] = max(s[i], v)
func (s *Slice) MaxToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Max(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = max(s.data[i], v)
		i += 1
	}
}

// Fill sets all elements to v
// performs: s[i] = v
func (s *Slice) Fill(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		vec.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = v
		i += 1
	}
}

// Negate flips the sign of all elements
// performs: s[i] = -s[i]
func (s *Slice) Negate() {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Neg()
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = -s.data[i]
		i += 1
	}
}

// Abs replaces all elements with their absolute value
// performs: s[i] = |s[i]|
func (s *Slice) Abs() {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Abs()
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		if s.data[i] < 0 {
			s.data[i] = -s.data[i]
		}
		i += 1
	}
}

// AndToAll performs a bitwise AND between all elements and v
// performs: s[i] = s[i] & v
func (s *Slice) AndToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.And(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] &= v
		i += 1
	}
}

// OrToAll performs a bitwise OR between all elements and v
// performs: s[i] = s[i] | v
func (s *Slice) OrToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Or(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] |= v
		i += 1
	}
}

// XorToAll performs a bitwise XOR between all elements and v
// performs: s[i] = s[i] ^ v
func (s *Slice) XorToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Xor(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] ^= v
		i += 1
	}
}

// AndNotToAll clears the bits of v from all elements
// performs: s[i] = s[i] &^ v
func (s *Slice) AndNotToAll(v int32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastInt32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.AndNot(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] &^= v
		i += 1
	}
}

// Not flips every bit of all elements
// performs: s[i] = ^s[i]
func (s *Slice) Not() {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.Not()
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = ^s.data[i]
		i += 1
	}
}

// ShiftLeftToAll shifts every element left by count bits
// performs: s[i] = s[i] << count
func (s *Slice) ShiftLeftToAll(count uint) {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int
	shift := uint64(count)

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.ShiftAllLeft(shift)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] <<= count
		i += 1
	}
}

// ShiftRightToAll arithmetically shifts every element right by count bits
// performs: s[i] = s[i] >> count
func (s *Slice) ShiftRightToAll(count uint) {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int
	shift := uint64(count)

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		v = v.ShiftAllRight(shift)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] >>= count
		i += 1
	}
}

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() int32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	acc := archsimd.BroadcastInt32x16(0)

	for i < (len(s.data) - l) {
		v := archsimd.LoadInt32x16(s.data[i : i+l])
		acc = acc.Add(v)
		i += l
	}

	var arr [16]int32
	acc.StoreArray(&arr)
	sum := arr[0] + arr[1] + arr[2] + arr[3] + arr[4] + arr[5] + arr[6] + arr[7] + arr[8] + arr[9] + arr[10] + arr[11] + arr[12] + arr[13] + arr[14] + arr[15]

	for i < len(s.data) {
		sum += s.data[i]
		i += 1
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() int32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadInt32x16(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadInt32x16(s.data[i : i+l])
			acc = acc.Min(v)
			i += l
		}

		var arr [16]int32
		acc.StoreArray(&arr)
		m = arr[0]
		for _, v := range arr[1:] {
			m = min(m, v)
		}
	}

	for i < len(s.data) {
		m = min(m, s.data[i])
		i += 1
	}
	return m
}

// Max returns the largest element, or 0 for an empty slice
// performs: max(s[i])
func (s *Slice) Max() int32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadInt32x16(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadInt32x16(s.data[i : i+l])
			acc = acc.Max(v)
			i += l
		}

		var arr [16]int32
		acc.StoreArray(&arr)
		m = arr[0]
		for _, v := range arr[1:] {
			m = max(m, v)
		}
	}

	for i < len(s.data) {
		m = max(m, s.data[i])
		i += 1
	}
	return m
}
