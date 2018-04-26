package story

import (
	"encoding/json"
	"io"
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

func DecodeStory(f io.Reader) (Story, error) {
	dec := json.NewDecoder(f)

	var s Story
	err := dec.Decode(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
