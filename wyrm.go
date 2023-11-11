package main

import (
	"fmt"

	"github.com/jnieblas/wyrm/dal"
	"github.com/jnieblas/wyrm/migration"
)

func main() {
	migration.ProvisionDB()
	dal.CreateScript()
	res := dal.GetScripts()
	fmt.Println("res:", res)
}
