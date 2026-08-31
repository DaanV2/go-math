//go:build simd_avx512

package simdfloat32

import (
	"math"

	"simd/archsimd"
)

// AddToAll adds v to all elements, updating their value with the update value
// performs: s[i] = s[i] + v
func (s *Slice) AddToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
func (s *Slice) MulToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
func (s *Slice) SubToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
func (s *Slice) DivToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
		v = v.Div(vec)
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] /= v
		i += 1
	}
}

// MinToAll sets each element to the minimum of itself and v
// performs: s[i] = min(s[i], v)
func (s *Slice) MinToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
func (s *Slice) MaxToAll(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

	l := vec.Len()
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
func (s *Slice) Fill(v float32) {
	if s == nil || len(s.data) == 0 {
		return
	}

	vec := archsimd.BroadcastFloat32x16(v)

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
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
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
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
		v = v.Abs()
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = float32(math.Abs(float64(s.data[i])))
		i += 1
	}
}

// Sqrt replaces all elements with their square root
// performs: s[i] = sqrt(s[i])
func (s *Slice) Sqrt() {
	if s == nil || len(s.data) == 0 {
		return
	}

	l := 16
	var i int

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
		v = v.Sqrt()
		v.Store(s.data[i : i+l])
		i += l
	}

	for i < len(s.data) {
		s.data[i] = float32(math.Sqrt(float64(s.data[i])))
		i += 1
	}
}

// Sum returns the sum of all elements, or 0 for an empty slice
// performs: sum(s[i])
func (s *Slice) Sum() float32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	acc := archsimd.BroadcastFloat32x16(0)

	for i < (len(s.data) - l) {
		v := archsimd.LoadFloat32x16(s.data[i : i+l])
		acc = acc.Add(v)
		i += l
	}

	var arr [16]float32
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
func (s *Slice) Min() float32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadFloat32x16(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadFloat32x16(s.data[i : i+l])
			acc = acc.Min(v)
			i += l
		}

		var arr [16]float32
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
func (s *Slice) Max() float32 {
	if s == nil || len(s.data) == 0 {
		return 0
	}

	l := 16
	var i int
	m := s.data[0]

	if len(s.data) >= l {
		acc := archsimd.LoadFloat32x16(s.data[0:l])
		i = l
		for i < (len(s.data) - l) {
			v := archsimd.LoadFloat32x16(s.data[i : i+l])
			acc = acc.Max(v)
			i += l
		}

		var arr [16]float32
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
