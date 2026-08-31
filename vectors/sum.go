package vectors

type Addable[T any] interface {
	Add(other T) T
}

func Sum[T Addable[T]](items []T) (result T) {
	for _, v := range items {
		result = result.Add(v)
	}

	return result
}
