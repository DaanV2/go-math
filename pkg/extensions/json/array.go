package xjson

import "encoding/json/v2"

type Numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | complex64 | complex128
}

type Storable[T Numbers] interface {
	Store(receiver []T)
}

func UnmarshalInto[T Numbers, U Storable[T]](data []byte, size int, transform func(data []T) U) (U, error) {
	var result U
	tmp := make([]T, 0, size)

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return result, err
	}

	return transform(tmp), nil
}

func MarshalAsArray[T Numbers, U Storable[T]](s U, size int) ([]byte, error) {
	tmp := make([]T, size)
	s.Store(tmp)
	data, err := json.Marshal(tmp)

	return data, err
}
