package simdint32

type Slice struct {
	data []int32
}

// NewSlice constructs a slice simd manipulator on an array of items
// WARNING: once the data is passed it, its assumed the [Slice] is the owner until [Slice.Output] is used
func NewSlice(data []int32) *Slice {
	return &Slice{data}
}

func (s *Slice) Output() []int32 {
	return s.data
}
