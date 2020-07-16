package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Append second line
	file, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("First line\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteString("second line"); err != nil {
		log.Fatal(err)
	}

	//Print the contents of the file
	data, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
