package simd_test

const DEFAULT_LEN_PRIME int = 101
const DEFAULT_SEED1 uint64 = 11332761825998879000
const DEFAULT_SEED2 uint64 = 5109035920914201000

type Interger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
