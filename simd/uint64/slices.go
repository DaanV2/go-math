package simduint64

type Slice struct {
	data []uint64
}

// NewSlice constructs a slice simd manipulator on an array of items
// WARNING: once the data is passed it, its assumed the [Slice] is the owner until [Slice.Output] is used
func NewSlice(data []uint64) *Slice {
	return &Slice{data}
}

func (s *Slice) Output() []uint64 {
	return s.data
}
