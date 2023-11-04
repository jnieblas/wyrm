package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "wyrm", "The name to greet.")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
	if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	}

}
