//go:build simd_avx256

package simduint64

// NOTE: this file uses the local Uint64x4 wrapper (defined alongside this
// package) rather than archsimd.Uint64x4 directly, since native
// Mul/Min/Max on 64-bit lanes require AVX-512 instructions
// (VPMULLQ/VPMINUQ/VPMAXUQ) that don't exist on plain AVX2 hardware. The
// wrapper emulates them with AVX2-safe sequences.

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) MulToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) SubToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) DivToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] /= v
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) MaxToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) Fill(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

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
func (s *Slice) AndToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) OrToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) XorToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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
func (s *Slice) AndNotToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := NewUint64x4Boardcast(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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

	l := 4
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
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

	l := 4
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
		v = v.ShiftLeft(count)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] <<= count
		i += 1
	}
}

// ShiftRightToAll logically shifts every element right by count bits
// performs: s[i] = s[i] >> count
func (s *Slice) ShiftRightToAll(count uint) {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 4
	var i int

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
		v = v.ShiftRight(count)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] >>= count
		i += 1
	}
}

// Add adds v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] + v[i]
func (s *Slice) Add(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Sub(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Mul(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Div(v []uint64) {
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
func (s *Slice) MinWith(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) MaxWith(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) And(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Or(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Xor(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) AndNot(v []uint64) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	l := 4
	var i int

	for i < (n - l) {
		a := NewUint64x4(s.data[i : i+l])
		b := NewUint64x4(v[i : i+l])
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
func (s *Slice) Sum() uint64 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 4
	var i int
	acc := NewUint64x4Boardcast(0)

	for i < (len(s.data) - l) {
		v := NewUint64x4(s.data[i : i+l])
		acc = acc.Add(v)
		i += l
	}

	var arr [4]uint64
	acc.Store(arr[:])
	sum := arr[0] + arr[1] + arr[2] + arr[3]

	for i < len(s.data) {
		sum += s.data[i]
		i += 1
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() uint64 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 4
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := NewUint64x4(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := NewUint64x4(s.data[i : i+l])
			acc = acc.Min(v)
			i += l
		}

		var arr [4]uint64
		acc.Store(arr[:])
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
func (s *Slice) Max() uint64 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 4
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := NewUint64x4(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := NewUint64x4(s.data[i : i+l])
			acc = acc.Max(v)
			i += l
		}

		var arr [4]uint64
		acc.Store(arr[:])
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
