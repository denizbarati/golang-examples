package main

import (
	"fmt"
)

func main() {
	maps()
	maps2()
	mutliMaps()
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func maps() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{12.02, 74.3}
	m["aaa"] = Vertex{12.501, 0.0002}
	fmt.Printf("%#v \n", m["aaa"])
}

func maps2() {
	mapExample := map[string]Vertex{
		"first ": {12, 34},
		"second": {20, 5},
	}
	fmt.Println(mapExample)
}

func mutliMaps() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	m["Splice"] = Vertex{34.05641, -118.48175}
	fmt.Printf("%v\n", m)
	delete(m, "Splice")
	fmt.Println(m["Splice"])
	name, ok := m["Splice"]
	fmt.Printf("key 'Splice' is present?: %t - value: %v\n", ok, name)
	name, ok = m["Google"]
	fmt.Printf("key 'Google' is present?: %t - value: %v\n", ok, name)
}