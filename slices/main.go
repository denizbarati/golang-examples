package main

import "fmt"

func main() {
	Slices()
	length()
	NilSlice()
	MakingSlices()
	AppendingToSlice()
	AppendingToSlice2()
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

func length() {
	cities := []string{
		"LA",
		"San Diego",
		"San Francisco",
	}
	fmt.Println(cap(cities))

	// make slice and check length of that
	countries := make([]string, 50)
	fmt.Println(len(countries))
}

func NilSlice() { // The zero value of a slice is nil. A nil slice has a length and capacity of 0.
	var sum []int
	fmt.Println(sum)
	fmt.Println(len(sum))
	fmt.Println(cap(sum))
	fmt.Println(sum == nil)
}

func MakingSlices() {
	cities := make([]string, 3)
	cities[0] = "London"
	cities[1] = "Paris"
	cities[2] = "LA"
	fmt.Println("makingSlices func: ", cities)
}

func AppendingToSlice() {
	cities := []string{}
	cities = append(cities, "mountain view", "river road")
	fmt.Printf("AppendingToSlice func: %q \n", cities)
}

func AppendingToSlice2() {
	cities := []string{"LA", "London", "Paris"}
	otherCities := []string{"mountain view", "river road"}
	cities = append(cities, otherCities...)
	fmt.Printf("AppendingToSlice2 func: %q", cities)
}
