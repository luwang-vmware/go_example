package main

import (
	"fmt"
	"log"

	"hash/fnv"
	"hash/maphash"

	"luwang.com/goexamples/go_modules"
	"luwang.com/goexamples/go_modules/mod_test"
)

func main() {
	go_modules.Hello("test")
	p := mod_test.Person{
		Name: "luwang",
		Age:  38,
	}
	p.Print()

	return
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("test")
	// o, r := os.UserHomeDir()
	// log.Println(o, r)
	// err := os.WriteFile("./tt/new_file", []byte("this is new"), 0o644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// n.test()
	// tel := n.Telephone{Name: "home", Tel: "12344"}
	// p := n.Person{Name: "luwang", Age: 30, Telep: tel}

	// p.Print()

	arrays := []string{"aa", "bb", "cc"}
	for id, name := range arrays {
		fmt.Println(id, name)
	}

	arrays = append(arrays, "newadd")
	fmt.Println(arrays)

	arrays[1] = "updated-bb"
	fmt.Println(arrays)

	copy(arrays[3:], arrays[4:])
	fmt.Println(arrays[:len(arrays)-1])

	var h maphash.Hash

	// Add a string to the hash, and print the current hash value.
	h.WriteString("pool-1")
	fmt.Printf("%#x\n", h.Sum64())
	fmt.Println(h.Sum64())

	h1 := fnv.New64a()
	h1.Write([]byte("pool-1"))
	fmt.Println(h1.Sum64())
	ss := fmt.Sprintf("np: %v", h1.Sum64())
	fmt.Println(ss)

}
