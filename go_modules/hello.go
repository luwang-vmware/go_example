package go_modules

import (
	"fmt"
)

func Hello(param string) error {
	fmt.Println("hello", param)
	return nil
}
