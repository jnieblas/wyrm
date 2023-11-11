package service

import (
	"log"

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

	log.Println("Script created successfully.")
}

func GetScripts() {
	scripts := dao.GetScripts()
	for _, script := range scripts {
		// print scripts for CLI
	}
}

func GetScript(name *string) {
	dao.GetScript(*name)
	// do something with it
}
