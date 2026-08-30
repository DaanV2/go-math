//go:build simd_avx512

package simdint64

// Abs returns the absolute values of the elements of x
func (x Int64x4) Abs() Int64x4 {
	return Int64x4{
		data: x.data.Abs(),
	}
}

// Mul performs a fused: x * y.
func (x Int64x4) Mul(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x4) Max(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x4) Min(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Min(y.data),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x4) ShiftRight(count uint) Int64x4 {
	return Int64x4{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
