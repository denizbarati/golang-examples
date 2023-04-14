package main

import "fmt"

func main() {
	Slices()
}

func Slices() {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// whole slices
	// [1 2 3 4 5 6 7 8 9]
	fmt.Println(p)

	// start from index 2 and end in index 5 but index 5 is not count
	// [3 4 5]
	fmt.Println(p[2:5])

	// end in index 6 but index 6 is not count
	// [1 2 3 4 5 6]
	fmt.Println(p[:6])

	// start from index 4
	// [5 6 7 8 9]
	fmt.Println(p[4:])
}
