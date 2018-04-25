package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	filename := flag.String("file", "", "input a filename you want to open")
	flag.Parse()
	fmt.Println("flag: ", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	var s Story
	err = dec.Decode(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", s)

	// ==============================

	//fileinfo, err := file.Stat()
	//if err != nil {
	//log.Fatal(err)
	//}

	//filesize := fileinfo.Size()
	//buffer := make([]byte, filesize)

	//bytesread, err := file.Read(buffer)
	//if err != nil {
	//log.Fatal(err)
	//}

	//fmt.Println("bytes read: ", bytesread)
	//fmt.Println("bytestream to string: ", string(buffer))

}
