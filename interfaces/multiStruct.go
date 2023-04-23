package main

import "fmt"

type Rectangle struct {
	length, breadth float32
}

type Triangle struct {
	base, height float32
}

type Shape interface {
	area() float32
}

func (r Rectangle) area() float32 {
	return r.breadth + r.length
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

func calculate(s Shape) float32 {
	return s.area()
}

func main() {
	r := Rectangle{1, 5}
	t := Triangle{7, 1}

	rCalculate := calculate(r)
	tCalculate := calculate(t)

	fmt.Println("Rectangle area is : ", rCalculate)
	fmt.Println("Triangle area is : ", tCalculate)
}
