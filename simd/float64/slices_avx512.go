//go:build simd_avx512

package simdfloat64

import "simd/archsimd"

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v float64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat64x8(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat64x8(s.data[i : i+l])
		v = v.Add(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = s.data[i] + v
		i += 1
	}
}

// AddToAll multiplies v all elements, updating their value with the update value
// performs: s[i] = s[i] * v
func (s *Slice) MulToAll(v float64) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat64x8(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat64x8(s.data[i : i+l])
		v = v.Mul(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] *= v
		i += 1
	}
}
