package simdint64

type Slice struct {
	data []int64
}

// NewSlice constructs a slice simd manipulator on an array of items
// WARNING: once the data is passed it, its assumed the [Slice] is the owner until [Slice.Output] is used
func NewSlice(data []int64) *Slice {
	return &Slice{data}
}

func (s *Slice) Output() []int64 {
	return s.data
}
