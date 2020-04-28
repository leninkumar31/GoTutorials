package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Errpr: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
