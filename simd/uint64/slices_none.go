//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint64

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] += v
	}
}

// AddToAll multiplies v all elements, updating their value with the update value
// performs: s[i] = s[i] * v
func (s *Slice) MulToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] *= v
	}
}

// SubToAll subtracts v from all elements, updating their value with the result
// performs: s[i] = s[i] - v
func (s *Slice) SubToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] -= v
	}
}

// DivToAll divides all elements by v, updating their value with the result
// performs: s[i] = s[i] / v
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

	for i := range s.data {
		s.data[i] = min(s.data[i], v)
	}
}

// MaxToAll sets each element to the maximum of itself and v
// performs: s[i] = max(s[i], v)
func (s *Slice) MaxToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = max(s.data[i], v)
	}
}

// Fill sets all elements to v
// performs: s[i] = v
func (s *Slice) Fill(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = v
	}
}

// AndToAll performs a bitwise AND between all elements and v
// performs: s[i] = s[i] & v
func (s *Slice) AndToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] &= v
	}
}

// OrToAll performs a bitwise OR between all elements and v
// performs: s[i] = s[i] | v
func (s *Slice) OrToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] |= v
	}
}

// XorToAll performs a bitwise XOR between all elements and v
// performs: s[i] = s[i] ^ v
func (s *Slice) XorToAll(v uint64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] ^= v
	}
}

// AndNotToAll clears the bits of v from all elements
// performs: s[i] = s[i] &^ v
func (s *Slice) AndNotToAll(v uint64) {
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

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() uint64 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	var sum uint64
	for _, v := range s.data {
		sum += v
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() uint64 {
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
func (s *Slice) Max() uint64 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	m := s.data[0]
	for _, v := range s.data[1:] {
		m = max(m, v)
	}
	return m
}
