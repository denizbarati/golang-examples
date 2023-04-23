package main

import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

func (r Rectangle) area() float32 {
	return r.breadth + r.length
}

func calculate(s Shape) {
	fmt.Println("the area is :", s.area())
}

func main() {
	inputs := Rectangle{34.2, 23.1}
	calculate(inputs)
}
