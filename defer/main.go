package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	written, err := copyFile("srcFile.txt", "dstFile.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("No of bytes written are: ", written)
}

func copyFile(srcFile, dstFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
