package service

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"

	"github.com/jnieblas/wyrm/dao"
	"github.com/jnieblas/wyrm/dto"
)

func CreateScript(name *string, path *string, command *string, description *string) {
	formatHomeDir(path)

	script := dto.Script{
		Name:        *name,
		Path:        *path,
		Command:     *command,
		Description: *description,
	}

	dao.CreateScript(&script)
	fmt.Println("Script created successfully.")
}

func UpdateScript(name *string, path *string, command *string, description *string) {
	formatHomeDir(path)

	script := dto.Script{
		Name:        *name,
		Path:        *path,
		Command:     *command,
		Description: *description,
	}

	dao.UpdateScript(&script)
	fmt.Println("Script updated successfully.")
}

func GetScripts() {
	scripts := dao.GetScripts()

	fmt.Println("Name, Path, Command, Description")
	for _, script := range scripts {
		fmt.Printf("%s, %s, %s, %s\n", script.Name, script.Path, script.Command, script.Description)
	}
}

func GetScript(name *string) {
	script := dao.GetScript(*name)
	fmt.Print(script)
}

func ExecuteScript(name *string) {
	script := dao.GetScript(*name)
	cmd := exec.Command(script.Command)
	cmd.Dir = script.Path
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Command execution error:", err)
		return
	}
	fmt.Println(string(output))
}

func formatHomeDir(path *string) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Unable to find your home dir:", err)
		return
	}
	*path = strings.Replace(*path, "~", usr.HomeDir, 1)
}
