package main

import "fmt"

func main() {
	maps()
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func maps() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{12.02, 74.3}
	m["aaa"] = Vertex{12.501, 0.0002}
	fmt.Printf("%#v",m["aaa"])
}
