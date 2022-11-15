package mod_test

import "fmt"

type Person struct {
	Name  string
	Age   int
	Telep Telephone
}

func (p Person) Print() {
	fmt.Println(p.Name, p.Age)
	p.Telep.Print()
	test()
}
