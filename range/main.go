package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	Range()
}

func Range() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// "i" is index of array and "v" is value of array

	// also you can skip index or value by this syntax:
	for index := range pow {
		fmt.Println("index is: ", index)
	}
	// and to skip value :
	for _, value := range pow {
		fmt.Println("value is: ", value)
	}
}
