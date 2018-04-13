package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	p := &Chapter{
		Title: "Welcome to Create Your Own Adventure main page",
	}

	t, _ := template.ParseFiles("template.html")
	t.Execute(w, *p)
}

func introHandler(w http.ResponseWriter, r *http.Request, s *Chapter) {
	t, _ := template.ParseFiles("template.html")
	fmt.Println("this function at least fires")
	fmt.Println(s)
	t.Execute(w, s)
}

func main() {

	filename := flag.String("file", "gopher.json", "Importing JSON file")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	d := json.NewDecoder(f)
	var story Story
	if err := d.Decode(&story); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", story)

	//http.HandleFunc("/", rootHandler)
	//http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
	//introHandler(w, r, &text.Intro)
	//})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
