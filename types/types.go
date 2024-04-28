package types

type List[T any] interface {
	Length() int
	InsertAt(item T, index int) error
	Remove(item T) bool
	RemoveAt(index int) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index int) (T, error)
}
