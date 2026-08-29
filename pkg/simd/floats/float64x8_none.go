//go:build simd_none

package simdfloats

type Float64x8 struct {
	data [8]float64
}
