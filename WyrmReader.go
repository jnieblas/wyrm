package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type WyrmReader struct {
	*bufio.Reader
}

func CreateWyrmReader() WyrmReader {
	return WyrmReader{bufio.NewReader(os.Stdin)}
}

func (wr *WyrmReader) Ask(msg string, delegate *string) {
	fmt.Print(msg)
	res, err := wr.ReadString('\n')

	if err != nil {
		log.Fatal("Invalid value: " + err.Error())
	}

	res = strings.Replace(res, "\n", "", -1)
	*delegate = strings.Replace(res, "\r", "", -1)
}
