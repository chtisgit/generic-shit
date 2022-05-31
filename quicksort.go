// Quicksort implemented from memory inside ...

package genericshit

import "fmt"
import "os"

type number interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | uint | int | float32 | float64
}

func Swap[T number](list []T, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

func Partition[T number](list []T, pred func(T) bool) int {
	i, j := 0, len(list)-1
	for i < j {
		if !pred(list[i]) {
			Swap(list, i, j)
			j--
		} else {
			i++
		}
	}

	for i < len(list) && pred(list[i]) {
		i++
	}

	return i
}

func Quicksort[T number](list []T) {
	last := len(list)-1
	pivot := list[last]
	i := Partition(list[:last], func(elem T) bool { return elem < pivot })
	Swap(list, i, last)

	if i >= 2 {
		Quicksort(list[:i])
	}
	if i+1 < len(list) {
		Quicksort(list[i+1:])
	}
}
