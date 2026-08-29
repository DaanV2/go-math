package simd_test

import (
	"math"
	"testing"

	"github.com/daanv2/go-math/pkg/simd"
	"github.com/stretchr/testify/require"
)

func Fuzz_Simd_Add_Float64(f *testing.F) {
	f.Add(DEFAULT_LEN_PRIME, DEFAULT_SEED1, DEFAULT_SEED2)

	f.Fuzz(func(t *testing.T, length int, seed1, seed2 uint64) {
		a, b, expected := addGenFloat64(t, length, seed1, seed2)

		receiver := make([]float64, len(a))
		simd.AddFloat64(a, b, receiver)

		require.Equal(t, expected, receiver)
	})
}

func Fuzz_Simd_Add_Float32(f *testing.F) {
	f.Add(DEFAULT_LEN_PRIME, DEFAULT_SEED1, DEFAULT_SEED2)

	f.Fuzz(func(t *testing.T, length int, seed1, seed2 uint64) {
		a, b, expected := addGenFloat32(t, length, seed1, seed2)

		receiver := make([]float32, len(a))
		simd.AddFloat32(a, b, receiver)

		require.Equal(t, expected, receiver)
	})
}

func Fuzz_Simd_Add_Uint64(f *testing.F) {
	f.Add(DEFAULT_LEN_PRIME, DEFAULT_SEED1, DEFAULT_SEED2)

	f.Fuzz(func(t *testing.T, length int, seed1, seed2 uint64) {
		a, b, expected := addGen[uint64](t, math.MaxUint64, length, seed1, seed2)

		receiver := make([]uint64, len(a))
		simd.AddUint64(a, b, receiver)

		require.Equal(t, expected, receiver)
	})
}

func Fuzz_Simd_Add_Uint32(f *testing.F) {
	f.Add(DEFAULT_LEN_PRIME, DEFAULT_SEED1, DEFAULT_SEED2)

	f.Fuzz(func(t *testing.T, length int, seed1, seed2 uint64) {
		a, b, expected := addGen[uint32](t, math.MaxUint32, length, seed1, seed2)

		receiver := make([]uint32, len(a))
		simd.AddUint32(a, b, receiver)

		require.Equal(t, expected, receiver)
	})
}
