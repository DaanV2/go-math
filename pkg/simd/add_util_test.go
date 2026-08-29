package simd_test

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func addGen[T Interger](t *testing.T, max T, length int, seed1, seed2 uint64) (elems1, elems2, expected []T) {
	t.Helper()
	elems1 = []T{}
	elems2 = []T{}
	expected = []T{}

	if length < 0 {
		return
	}

	rnd := randomizer(seed1, seed2)

	for range length {
		var elem1 T = rnd.N(max)
		var elem2 T = rnd.N(max)
		result := elem1 + elem2

		elems1 = append(elems1, elem1)
		elems2 = append(elems2, elem2)
		expected = append(expected, result)
	}

	require.Len(t, elems1, length)
	require.Len(t, elems2, length)
	require.Len(t, expected, length)

	return
}

func addGenFloat64(t *testing.T, length int, seed1, seed2 uint64) (elems1, elems2, expected []float64) {
	t.Helper()
	elems1 = []float64{}
	elems2 = []float64{}
	expected = []float64{}

	if length < 0 {
		return
	}

	rnd := randomizer(seed1, seed2)

	for range length {
		elem1 := rnd.Float64()
		elem2 := rnd.Float64()
		result := elem1 + elem2

		elems1 = append(elems1, elem1)
		elems2 = append(elems2, elem2)
		expected = append(expected, result)
	}

	require.Len(t, elems1, length)
	require.Len(t, elems2, length)
	require.Len(t, expected, length)

	return
}

func addGenFloat32(t *testing.T, length int, seed1, seed2 uint64) (elems1, elems2, expected []float32) {
	t.Helper()
	elems1 = []float32{}
	elems2 = []float32{}
	expected = []float32{}

	if length < 0 {
		return
	}
	rnd := randomizer(seed1, seed2)

	for range length {
		elem1 := rnd.Float32()
		elem2 := rnd.Float32()
		result := elem1 + elem2

		elems1 = append(elems1, elem1)
		elems2 = append(elems2, elem2)
		expected = append(expected, result)
	}

	require.Len(t, elems1, length)
	require.Len(t, elems2, length)
	require.Len(t, expected, length)

	return
}

func randomizer(seed1, seed2 uint64) *rand.Rand {
	seed := rand.NewPCG(seed1, seed2)
	return rand.New(seed)
}
