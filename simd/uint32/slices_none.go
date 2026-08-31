//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint32

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] += v
	}
}

// AddToAll multiplies v all elements, updating their value with the update value
// performs: s[i] = s[i] * v
func (s *Slice) MulToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] *= v
	}
}

// SubToAll subtracts v from all elements, updating their value with the result
// performs: s[i] = s[i] - v
func (s *Slice) SubToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] -= v
	}
}

// DivToAll divides all elements by v, updating their value with the result
// performs: s[i] = s[i] / v
func (s *Slice) DivToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] /= v
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = min(s.data[i], v)
	}
}

// MaxToAll sets each element to the maximum of itself and v
// performs: s[i] = max(s[i], v)
func (s *Slice) MaxToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = max(s.data[i], v)
	}
}

// Fill sets all elements to v
// performs: s[i] = v
func (s *Slice) Fill(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = v
	}
}

// AndToAll performs a bitwise AND between all elements and v
// performs: s[i] = s[i] & v
func (s *Slice) AndToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] &= v
	}
}

// OrToAll performs a bitwise OR between all elements and v
// performs: s[i] = s[i] | v
func (s *Slice) OrToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] |= v
	}
}

// XorToAll performs a bitwise XOR between all elements and v
// performs: s[i] = s[i] ^ v
func (s *Slice) XorToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] ^= v
	}
}

// AndNotToAll clears the bits of v from all elements
// performs: s[i] = s[i] &^ v
func (s *Slice) AndNotToAll(v uint32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] &^= v
	}
}

// Not flips every bit of all elements
// performs: s[i] = ^s[i]
func (s *Slice) Not() {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = ^s.data[i]
	}
}

// ShiftLeftToAll shifts every element left by count bits
// performs: s[i] = s[i] << count
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
func (s *Slice) Add(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] += v[i]
	}
}

// Sub subtracts v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] - v[i]
func (s *Slice) Sub(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] -= v[i]
	}
}

// Mul multiplies v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] * v[i]
func (s *Slice) Mul(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] *= v[i]
	}
}

// Div divides s by v element-wise, updating s in place over min(len(s), len(v)) elements
// performs: s[i] = s[i] / v[i]
func (s *Slice) Div(v []uint32) {
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
func (s *Slice) MinWith(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] = min(s.data[i], v[i])
	}
}

// MaxWith sets each element to the maximum of itself and the corresponding
// element of v, over min(len(s), len(v)) elements
// performs: s[i] = max(s[i], v[i])
func (s *Slice) MaxWith(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] = max(s.data[i], v[i])
	}
}

// And performs a bitwise AND against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] & v[i]
func (s *Slice) And(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] &= v[i]
	}
}

// Or performs a bitwise OR against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] | v[i]
func (s *Slice) Or(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] |= v[i]
	}
}

// Xor performs a bitwise XOR against v element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] ^ v[i]
func (s *Slice) Xor(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] ^= v[i]
	}
}

// AndNot clears the bits of v from s element-wise, over min(len(s), len(v)) elements
// performs: s[i] = s[i] &^ v[i]
func (s *Slice) AndNot(v []uint32) {
	if s == nil || len(s.data) == 0 || len(v) == 0 {
		return
	}

	n := min(len(s.data), len(v))
	for i := range n {
		s.data[i] &^= v[i]
	}
}

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() uint32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	var sum uint32
	for _, v := range s.data {
		sum += v
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() uint32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	m := s.data[0]
	for _, v := range s.data[1:] {
		m = min(m, v)
	}
	return m
}

// Max returns the largest element, or 0 for an empty slice
// performs: max(s[i])
func (s *Slice) Max() uint32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	m := s.data[0]
	for _, v := range s.data[1:] {
		m = max(m, v)
	}
	return m
}
