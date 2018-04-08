package main

import (
	"encoding/json"
	//"html/template"
	"fmt"
	"io/ioutil"
	"os"
)

type CreateYourOwnAdventure struct {
	Intro     StoryObject `json:"intro"`
	Newyork   StoryObject `json:"new-york"`
	Debate    StoryObject `json:"debate"`
	SeanKelly StoryObject `json:"sean-kelly"`
	MarkBates StoryObject `json:"mark-bates"`
	Denver    StoryObject `json:"denver"`
}

type StoryObject struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []OptionType `json:"options"`
}

type OptionType struct {
	text string `json:"text"`
	arc  string `json:"arc"`
}

func main() {
	// REQUIREMENTS
	// Using html/template package, create HTML pages
	// Create an http.Handler to handle web requests instead of a handler function

	// 1. Import gopher.json file as go-readable format (using "encoidng/json" package)
	raw, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	var c CreateYourOwnAdventure
	json.Unmarshal(raw, &c)
	fmt.Println("======== JUST PRINT OUT =========")
	fmt.Printf("%v", c)
	fmt.Println("-------- toJson --------")
	fmt.Println(toJson(c))

}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		//fmt.Println(err, Error())
		os.Exit(1)
	}
	return string(bytes)
}
