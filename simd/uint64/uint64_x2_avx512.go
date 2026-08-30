//go:build simd_avx512

package simduint64

// Mul performs a fused: x * y.
func (x Uint64x2) Mul(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x2) Max(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x2) Min(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Min(y.data),
	}
}
