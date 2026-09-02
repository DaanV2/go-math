//go:build simd_avx256

package simdbytes

import (
	"simd/archsimd"
)

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) MulToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) SubToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) DivToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] /= v
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) MaxToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) Fill(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

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

// AndToAll performs a bitwise AND between all elements and v
// performs: s[i] = s[i] & v
func (s *Slice) AndToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) OrToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) XorToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
func (s *Slice) AndNotToAll(v byte) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastUint8x32(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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

	l := 32
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
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
// NOTE: there is no hardware vectorised per-byte shift, this always runs scalar
func (s *Slice) ShiftLeftToAll(count uint) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] <<= count
	}
}

// ShiftRightToAll logically shifts every element right by count bits
// performs: s[i] = s[i] >> count
// NOTE: there is no hardware vectorised per-byte shift, this always runs scalar
func (s *Slice) ShiftRightToAll(count uint) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] >>= count
	}
}

// Add adds v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] + v[i]
func (s *Slice) Add(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Add(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] += v[i]
		i += 1
	}
}

// Sub subtracts v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] - v[i]
func (s *Slice) Sub(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Sub(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] -= v[i]
		i += 1
	}
}

// Mul multiplies v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] * v[i]
func (s *Slice) Mul(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Mul(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] *= v[i]
		i += 1
	}
}

// Div divides s by v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] / v[i]
// NOTE: there is no hardware vectorised integer division, this always runs scalar
func (s *Slice) Div(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] /= v[i]
	}
}

// MinWith sets each element to the minimum of itself and the corresponding
// element of v, over min(len(s), len(v)) elements
// performs: s[i] = min(s[i], v[i])
func (s *Slice) MinWith(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Min(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] = min(s.data[i], v[i])
		i += 1
	}
}

// MaxWith sets each element to the maximum of itself and the corresponding
// element of v, over min(len(s), len(v)) elements
// performs: s[i] = max(s[i], v[i])
func (s *Slice) MaxWith(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Max(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] = max(s.data[i], v[i])
		i += 1
	}
}

// And performs a bitwise AND against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] & v[i]
func (s *Slice) And(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.And(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] &= v[i]
		i += 1
	}
}

// Or performs a bitwise OR against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] | v[i]
func (s *Slice) Or(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Or(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] |= v[i]
		i += 1
	}
}

// Xor performs a bitwise XOR against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] ^ v[i]
func (s *Slice) Xor(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.Xor(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] ^= v[i]
		i += 1
	}
}

// AndNot clears the bits of v from s element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] &^ v[i]
func (s *Slice) AndNot(v []byte) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 32
	var i int

	for i < (n - l) {
		a := archsimd.LoadUint8x32(s.data[i : i+l])
		b := archsimd.LoadUint8x32(v[i : i+l])
		a = a.AndNot(b)
		a.Store(s.data[i : i+l])
		i += l
	}

	for i < n {
		s.data[i] &^= v[i]
		i += 1
	}
}

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() byte {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 32
	var i int
	acc := archsimd.BroadcastUint8x32(0)

	for i < (len(s.data) - l) {
		v := archsimd.LoadUint8x32(s.data[i : i+l])
		acc = acc.Add(v)
		i += l
	}

	var arr [32]byte
	acc.StoreArray(&arr)
	var sum byte
	for _, v := range arr {
		sum += v
	}

	for i < len(s.data) {
		sum += s.data[i]
		i += 1
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() byte {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 32
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadUint8x32(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadUint8x32(s.data[i : i+l])
			acc = acc.Min(v)
			i += l
		}

		var arr [32]byte
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
func (s *Slice) Max() byte {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 32
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadUint8x32(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadUint8x32(s.data[i : i+l])
			acc = acc.Max(v)
			i += l
		}

		var arr [32]byte
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
