package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func createScript(script *Script) {
	formatHomeDir(&script.Path)

	createScriptExecutor(script)
	fmt.Println("Script created successfully.")
}

func updateScript(name *string, path *string, command *string, description *string) {
	formatHomeDir(path)

	script := Script{
		Name:        *name,
		Path:        *path,
		Command:     *command,
		Description: *description,
	}

	updateScriptExecutor(&script)
	fmt.Println("Script updated successfully.")
}

func getScripts() {
	scripts := getScriptsExecutor()

	fmt.Println("Name, Path, Command, Description")
	for _, script := range scripts {
		fmt.Printf("%s, %s, %s, %s\n", script.Name, script.Path, script.Command, script.Description)
	}
}

func getScript(name *string) {
	script := getScriptExecutor(*name)
	fmt.Print(script)
}

func executeScript(name *string) {
	script := getScriptExecutor(*name)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", script.Command, script.Path)
	default: //Mac & Linux
		cmd = exec.Command(script.Command)
	}

	cmd.Dir = script.Path
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Command execution error:", err)
		return
	}
	fmt.Println(string(output))
}

func formatHomeDir(path *string) {
	if (*path)[1] == '~' {
		usr, err := user.Current()

		if err != nil {
			fmt.Println("Unable to find your home dir:", err)
			return
		}

		fmt.Println("HomeDir: ", usr)
		*path = strings.Replace(*path, "~", usr.HomeDir, 1)
	}
}
