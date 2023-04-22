package main

import "fmt"

type USER struct {
	FirstName, LastName string
}

//method
func (u USER) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := USER{"Deniz", "Barati"}
	fmt.Println(u.Greeting())
}
