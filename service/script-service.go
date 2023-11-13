package service

import (
	"fmt"

	"github.com/jnieblas/wyrm/dao"
	"github.com/jnieblas/wyrm/dto"
)

func CreateScript(name *string, path *string, command *string, description *string) {
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

	fmt.Print(script)
}
