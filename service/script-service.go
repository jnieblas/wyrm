package service

import (
	"flag"

	"github.com/jnieblas/wyrm/dao"
	"github.com/jnieblas/wyrm/dto"
)

func CreateScript(name *string, path *string, command *string, description *string) {
	flag.Parse()

	script := dto.Script{
		Name:        *name,
		Path:        *path,
		Command:     *command,
		Description: *description,
	}

	dao.CreateScript(&script)
}
