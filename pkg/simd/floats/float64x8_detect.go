//go:build simd_detect || (!simd_avx256 && !simd_avx512 && !simd_none)

package simdfloats

import (
	"simd/archsimd"

	xruntime "github.com/daanv2/go-math/pkg/extensions/runtime"
)

type Float64x8 struct {
	data [8]float64
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8
	copy(result.data[:], data)

	return result
}

func (x Float64x8) Store(receiver []float64) {
	copy(receiver, x.data[:])
}

func (x Float64x8) Add(y Float64x8) Float64x8 { // nolint:dupl // keep the duplicate code
	var result Float64x8

	switch {
	case xruntime.AVX512():
		v1 := archsimd.LoadFloat64x8(x.data[:])
		v2 := archsimd.LoadFloat64x8(y.data[:])
		v1.Add(v2).Store(result.data[:])
	case xruntime.AVX256():
		// First 4
		v1 := archsimd.LoadFloat64x4(x.data[:4])
		v2 := archsimd.LoadFloat64x4(y.data[:4])
		v1.Add(v2).Store(result.data[:4])

		// Last 4
		v1 = archsimd.LoadFloat64x4(x.data[4:])
		v2 = archsimd.LoadFloat64x4(y.data[4:])
		v1.Add(v2).Store(result.data[4:])
	default:
		for i := range x.data {
			result.data[i] = x.data[i] + y.data[i]
		}
	}

	return result
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 { // nolint:dupl // keep the duplicate code
	var result Float64x8

	switch {
	case xruntime.AVX512():
		v1 := archsimd.LoadFloat64x8(x.data[:])
		v2 := archsimd.LoadFloat64x8(y.data[:])
		v1.Sub(v2).Store(result.data[:])
	case xruntime.AVX256():
		// First 4
		v1 := archsimd.LoadFloat64x4(x.data[:4])
		v2 := archsimd.LoadFloat64x4(y.data[:4])
		v1.Sub(v2).Store(result.data[:4])

		// Last 4
		v1 = archsimd.LoadFloat64x4(x.data[4:])
		v2 = archsimd.LoadFloat64x4(y.data[4:])
		v1.Sub(v2).Store(result.data[4:])
	default:
		for i := range x.data {
			result.data[i] = x.data[i] - y.data[i]
		}
	}

	return result
}

// Mul performs a fused: x * y.
func (x Float64x8) Mul(y Float64x8) Float64x8 { // nolint:dupl // keep the duplicate code
	var result Float64x8

	switch {
	case xruntime.AVX512():
		v1 := archsimd.LoadFloat64x8(x.data[:])
		v2 := archsimd.LoadFloat64x8(y.data[:])
		v1.Mul(v2).Store(result.data[:])
	case xruntime.AVX256():
		// First 4
		v1 := archsimd.LoadFloat64x4(x.data[:4])
		v2 := archsimd.LoadFloat64x4(y.data[:4])
		v1.Mul(v2).Store(result.data[:4])

		// Last 4
		v1 = archsimd.LoadFloat64x4(x.data[4:])
		v2 = archsimd.LoadFloat64x4(y.data[4:])
		v1.Mul(v2).Store(result.data[4:])
	default:
		for i := range x.data {
			result.data[i] = x.data[i] * y.data[i]
		}
	}

	return result
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x8) MulAdd(y, z Float64x8) Float64x8 {
	var result Float64x8

	switch {
	case xruntime.AVX512():
		v1 := archsimd.LoadFloat64x8(x.data[:])
		v2 := archsimd.LoadFloat64x8(y.data[:])
		v3 := archsimd.LoadFloat64x8(z.data[:])
		v1.Mul(v2).Add(v3).Store(result.data[:])
	case xruntime.AVX256():
		// First 4
		v1 := archsimd.LoadFloat64x4(x.data[:4])
		v2 := archsimd.LoadFloat64x4(y.data[:4])
		v3 := archsimd.LoadFloat64x4(z.data[:4])
		v1.Mul(v2).Add(v3).Store(result.data[:4])

		// Last 4
		v1 = archsimd.LoadFloat64x4(x.data[4:])
		v2 = archsimd.LoadFloat64x4(y.data[4:])
		v3 = archsimd.LoadFloat64x4(z.data[4:])
		v1.Mul(v2).Add(v3).Store(result.data[4:])
	default:
		for i := range x.data {
			result.data[i] = (x.data[i] * y.data[i]) + z.data[i]
		}
	}

	return result
}
