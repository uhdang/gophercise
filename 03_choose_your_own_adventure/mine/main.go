package main

import (
	"flag"
	"fmt"
	"github.com/uhdang/gophercise/03_choose_your_own_adventure/mine/story"
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

	o, err := story.DecodeStory(file)

	fmt.Println(o["intro"])
	fmt.Println(o["intro"].Title)

}
