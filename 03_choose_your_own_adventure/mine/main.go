package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "input a filename you want to open")
	flag.Parse()
	fmt.Println("flag: ", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))

}
