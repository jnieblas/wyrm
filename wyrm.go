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
	path := flag.String("path", "", "Absolute path to script")
	command := flag.String("command", "", "Command needed to run script")
	description := flag.String("description", "", "Script description")
	flag.Parse()

	if !*enableLogging {
		log.SetOutput(io.Discard)
	}

	if *provision {
		provisionDB()
		os.Exit(0)
	}

	if *create {
		if validateRequiredFlags("c", name, path, command) {
			createScript(name, path, command, description)
		}
	} else if *update {
		if validateRequiredFlags("u", name, path, command) {
			updateScript(name, path, command, description)
		}
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

func validateRequiredFlags(flagName string, name *string, path *string, command *string) bool {
	var missingFlags []string

	if *name == "" {
		missingFlags = append(missingFlags, "name")
	}

	if *path == "" {
		missingFlags = append(missingFlags, "path")
	}

	if *command == "" {
		missingFlags = append(missingFlags, "command")
	}

	if len(missingFlags) > 0 {
		fmt.Printf("Missing required flags for -%s:\n", flagName)

		for _, missingFlag := range missingFlags {
			fmt.Printf("-%s\n", missingFlag)
		}
		return false
	}

	return true
}
