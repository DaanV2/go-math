//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdfloat32

import "math"

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] += v
	}
}

// AddToAll multiplies v all elements, updating their value with the update value
// performs: s[i] = s[i] * v
func (s *Slice) MulToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] *= v
	}
}

// SubToAll subtracts v from all elements, updating their value with the result
// performs: s[i] = s[i] - v
func (s *Slice) SubToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] -= v
	}
}

// DivToAll divides all elements by v, updating their value with the result
// performs: s[i] = s[i] / v
func (s *Slice) DivToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] /= v
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = min(s.data[i], v)
	}
}

// MaxToAll sets each element to the maximum of itself and v
// performs: s[i] = max(s[i], v)
func (s *Slice) MaxToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = max(s.data[i], v)
	}
}

// Fill sets all elements to v
// performs: s[i] = v
func (s *Slice) Fill(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = v
	}
}

// Negate flips the sign of all elements
// performs: s[i] = -s[i]
func (s *Slice) Negate() {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = -s.data[i]
	}
}

// Abs replaces all elements with their absolute value
// performs: s[i] = |s[i]|
func (s *Slice) Abs() {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = float32(math.Abs(float64(s.data[i])))
	}
}

// Sqrt replaces all elements with their square root
// performs: s[i] = sqrt(s[i])
func (s *Slice) Sqrt() {
	if s == nil || len(s.data) == 0 {
		return
	}

	for i := range s.data {
		s.data[i] = float32(math.Sqrt(float64(s.data[i])))
	}
}

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() float32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	var sum float32
	for _, v := range s.data {
		sum += v
	}
	return sum
}

// Min returns the smallest element, or 0 for an empty slice
// performs: min(s[i])
func (s *Slice) Min() float32 {
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
func (s *Slice) Max() float32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	m := s.data[0]
	for _, v := range s.data[1:] {
		m = max(m, v)
	}
	return m
}
