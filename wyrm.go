package main

import (
	"flag"

	"github.com/jnieblas/wyrm/service"
)

func main() {
	create := flag.Bool("c", false, "Enable special behavior")
	name := flag.String("name", "", "Script name.")
	path := flag.String("path", "", "Script path.")
	command := flag.String("command", "", "Command needed to run script.")
	description := flag.String("description", "", "Script description.")

	flag.Parse()

	if *create {
		service.CreateScript(name, path, command, description)
	}
}
