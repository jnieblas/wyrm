package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jnieblas/wyrm/migration"
	"github.com/jnieblas/wyrm/service"
)

func main() {
	// Create / update flags
	provision := flag.Bool("provision", false, "Provision database.")
	enableLogging := flag.Bool("l", false, "Enable logging.")
	create := flag.Bool("c", false, "Create a new script; requires -name, -path and -command")
	update := flag.Bool("u", false, "Update an existing script; requires -name, -path and -command")
	info := flag.Bool("i", false, "List information about one or multiple scripts")
	name := flag.String("name", "", "Script name, used to reference a script")
	path := flag.String("path", "", "Absolute path to script")
	command := flag.String("command", "", "Command needed to run script")
	description := flag.String("description", "", "Script description")
	flag.Parse()

	if *provision {
		migration.ProvisionDB()
		os.Exit(0)
	}

	if !*enableLogging {
		log.SetOutput(io.Discard)
	}

	if *create {
		if validateRequiredFlags("c", name, path, command) {
			service.CreateScript(name, path, command, description)
		}
	} else if *update {
		if validateRequiredFlags("u", name, path, command) {
			service.UpdateScript(name, path, command, description)
		}
	} else if *info {
		if *name != "" {
			service.GetScript(name)
		} else {
			service.GetScripts()
		}
	} else {
		if *name != "" {
			service.ExecuteScript(name)
		} else {
			args := flag.Args()

			if len(args) > 0 {
				scriptName := args[0]
				service.ExecuteScript(&scriptName)
			} else {
				fmt.Println("Invalid usage of wyrm.")
				flag.Usage()
			}

		}
	}
}

func validateRequiredFlags(flagName string, name *string, path *string, command *string) bool {
	missingFlags := []string{}

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

		for _, flag := range missingFlags {
			fmt.Printf("-%s\n", flag)
		}
		return false
	}

	return true
}
