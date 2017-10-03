package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// Human implement fmt.Stringer
func (h Human) String() string {
	return "Name: " + h.name + ", Age: " + strconv.Itoa(h.age) + " years, Contact: " + h.phone
}

func main() {
	Bob := Human{"Bob", 27, "222-654-415"}

	fmt.Println("This Human is: ", Bob)
}
