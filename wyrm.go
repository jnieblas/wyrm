package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "Then name to greet.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Memory Address, %p!\n", name)
}
