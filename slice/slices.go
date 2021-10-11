package slice

import "fmt"

func Slice_trial() {
	var x []float64
	x = make([]float64, 5, 10)
	fmt.Printf("x:%v, x's length:%d,  x's capaciy:%d\n", x, len(x), cap(x))
}

func Slice_append() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Printf("slice1: %v, slice2: %v\n", slice1, slice2)
}

func Slice_copy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Printf("slice1: %v, slie2: %v\n", slice1, slice2)
}
