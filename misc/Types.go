package misc

type List[T any] interface {
	Length() int
	InsertAt(item T, index int) error
	Remove(item T) bool
	RemoveAt(index int) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index int) (T, error)
}

type Point struct {
	X int
	Y int
}

type BinaryTreeNode[T any] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
	Prev *ListNode[T]
}
