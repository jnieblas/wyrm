package main

import (
	"fmt"

	"github.com/jnieblas/wyrm/services"
)

func main() {
	services.ProvisionDB()
	services.CreateScript()
	res := services.GetScripts()
	fmt.Println("res:", res)
}
