package mod_test

import "fmt"

func test() {
	fmt.Println("-------")
}

type Telephone struct {
	Name string
	Tel  string
}

func (t Telephone) Print() {
	fmt.Println(t)
}
