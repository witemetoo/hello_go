package main

import "os"
import "log"
import "fmt"

func readFile() {
	data := make([]byte, 100)
	file, err := os.Open("hello.tx")
	if err != nil {
		log.Fatal(err)
	}
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %s bytes:%q\n", count, data[:count])
}

func main() {
	readFile()
}
