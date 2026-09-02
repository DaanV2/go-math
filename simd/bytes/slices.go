package simdbytes

type Slice struct {
	data []byte
}

// NewSlice constructs a slice simd manipulator on an array of items
// WARNING: once the data is passed it, its assumed the [Slice] is the owner until [Slice.Output] is used
func NewSlice(data []byte) *Slice {
	return &Slice{data}
}

func (s *Slice) Output() []byte {
	return s.data
}
