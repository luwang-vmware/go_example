package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("test")
	o, r := os.UserHomeDir()
	log.Println(o, r)
	err := os.WriteFile("./tt/new_file", []byte("this is new"), 0o644)
	if err != nil {
		log.Println(err)
	}
}
