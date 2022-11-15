package go_modules

import (
	"fmt"

	"luwang.com/goexamples/go_modules/mod_test"
)

func Hello(param string) error {
	fmt.Println("hello", param)
	p := mod_test.Person{
		Name: "luwang",
		Age:  38,
	}
	p.Print()

	return nil
}
