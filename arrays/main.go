package main

import "fmt"

func main() {
	Arrays()
	PrintingArrays()
	MultiDimensionalArrays()
}

func Arrays() {
	var a [2]string
	a[0] = "Deniz"
	a[1] = "Barati"
	fmt.Println("Arrays func: ", a[0], a[1])
}

func PrintingArrays() {
	a := [2]string{"Deniz", "Barati"}
	fmt.Println("PrintingArrays func: ", a)
}

func MultiDimensionalArrays() {
	var a [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("MultiDimensionalArrays func: %q", a)
}
