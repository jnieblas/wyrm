package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	enableLogging := flag.Bool("l", false, "Enable logging.")
	provision := flag.Bool("provision", false, "Provision database.")
	create := flag.Bool("c", false, "Create a new script; requires -name, -path and -command")
	update := flag.Bool("u", false, "Update an existing script; requires -name, -path and -command")
	info := flag.Bool("i", false, "List information about one or multiple scripts")
	name := flag.String("name", "", "Script name, used to reference a script")
	flag.Parse()

	if !*enableLogging {
		log.SetOutput(io.Discard)
	}

	if *provision {
		provisionDB()
		os.Exit(0)
	}

	if *create {
		wyrm := getCreateInput()
		createScript(&wyrm)
	} else if *update {
		wyrm := getCreateInput()
		updateScript(&wyrm)
	} else if *info {
		if *name != "" {
			getScript(name)
		} else {
			getScripts()
		}
	} else {
		if *name != "" {
			executeScript(name)
		} else {
			args := flag.Args()

			if len(args) > 0 {
				scriptName := args[0]
				executeScript(&scriptName)
			} else {
				fmt.Println("Invalid usage of wyrm.")
				flag.Usage()
			}

		}
	}
}

func getCreateInput() (wyrm Script) {
	reader := CreateWyrmReader()
	reader.Ask("Give your wyrm a name: ", &wyrm.Name)
	reader.Ask("Execution path: ", &wyrm.Path)
	reader.Ask("Command to Execute: ", &wyrm.Command)
	reader.Ask("Wyrm description (optional): ", &wyrm.Description)

	log.Println(wyrm.String())
	return wyrm
}
